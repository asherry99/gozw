package application

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"github.com/helioslabs/gozw/zwave/command-class"
	"github.com/helioslabs/gozw/zwave/command-class/association"
	"github.com/helioslabs/gozw/zwave/command-class/battery"
	"github.com/helioslabs/gozw/zwave/command-class/manufacturer-specific"
	"github.com/helioslabs/gozw/zwave/command-class/security"
	"github.com/helioslabs/gozw/zwave/command-class/version"
	"github.com/helioslabs/gozw/zwave/command-class/version-v2"
	"github.com/helioslabs/gozw/zwave/protocol"
	"github.com/helioslabs/gozw/zwave/serial-api"
	"github.com/helioslabs/proto"
)

// Node is an in-memory representation of a Z-Wave node
type Node struct {
	NodeID byte

	Capability          byte
	BasicDeviceClass    byte
	GenericDeviceClass  byte
	SpecificDeviceClass byte

	Failing bool

	CommandClasses proto.CommandClassSet

	NetworkKeySent bool

	ManufacturerID uint16
	ProductTypeID  uint16
	ProductID      uint16

	QueryStageSecurity     bool
	QueryStageManufacturer bool
	QueryStageVersions     bool

	queryStageSecurityComplete     chan bool
	queryStageManufacturerComplete chan bool
	queryStageVersionsComplete     chan bool

	application *Layer
}

func NewNode(application *Layer, nodeID byte) (*Node, error) {
	node := &Node{
		NodeID: nodeID,

		CommandClasses: proto.CommandClassSet{},

		QueryStageSecurity:     false,
		QueryStageManufacturer: false,
		QueryStageVersions:     false,

		queryStageSecurityComplete:     make(chan bool),
		queryStageManufacturerComplete: make(chan bool),
		queryStageVersionsComplete:     make(chan bool),

		application: application,
	}

	err := node.loadFromDb()
	if err != nil {
		initErr := node.initialize()
		if initErr != nil {
			return nil, initErr
		}

		node.saveToDb()
	}

	return node, nil
}

func (n *Node) loadFromDb() error {
	var data []byte
	err := n.application.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		data = bucket.Get([]byte{n.NodeID})

		if len(data) == 0 {
			return errors.New("Node not found")
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(data, n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) initialize() error {
	nodeInfo, err := n.application.serialAPI.GetNodeProtocolInfo(n.NodeID)
	if err != nil {
		fmt.Println(err)
	} else {
		n.setFromNodeProtocolInfo(nodeInfo)
	}

	if n.NodeID == 1 {
		// self is never failing
		n.Failing = false
	} else {
		failing, err := n.application.serialAPI.IsFailedNode(n.NodeID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		n.Failing = failing
	}

	return n.saveToDb()
}

func (n *Node) saveToDb() error {
	data, err := msgpack.Marshal(n)
	if err != nil {
		return err
	}

	return n.application.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		return bucket.Put([]byte{n.NodeID}, data)
	})
}

func (n *Node) IsSecure() bool {
	return n.CommandClasses.Supports(commandclass.Security)
}

func (n *Node) IsListening() bool {
	return n.Capability&0x80 == 0x80
}

func (n *Node) GetBasicDeviceClassName() string {
	return protocol.GetBasicDeviceTypeName(n.BasicDeviceClass)
}

func (n *Node) GetGenericDeviceClassName() string {
	return protocol.GetGenericDeviceTypeName(n.GenericDeviceClass)
}

func (n *Node) GetSpecificDeviceClassName() string {
	return protocol.GetSpecificDeviceTypeName(n.GenericDeviceClass, n.SpecificDeviceClass)
}

func (n *Node) SendCommand(command commandclass.Command) error {
	commandClass := commandclass.ID(command.CommandClassID())

	if commandClass == commandclass.Security {
		switch command.(type) {
		case *security.CommandsSupportedGet, *security.CommandsSupportedReport:
			return n.application.SendDataSecure(n.NodeID, command)
		}
	}

	if !n.CommandClasses.Supports(commandClass) {
		return errors.New("Command class not supported")
	}

	if n.CommandClasses.IsSecure(commandClass) {
		return n.application.SendDataSecure(n.NodeID, command)
	}

	return n.application.SendData(n.NodeID, command)
}

func (n *Node) AddAssociation(groupID byte, nodeIDs ...byte) error {
	// sort of an arbitrary limit for now, but I'm not sure what it should be
	if len(nodeIDs) > 20 {
		return errors.New("Too many associated nodes")
	}

	fmt.Println("Associating")

	return n.SendCommand(&association.Set{
		GroupingIdentifier: groupID,
		NodeId:             nodeIDs,
	})
}

func (n *Node) LoadSupportedSecurityCommands() error {
	return n.application.SendDataSecure(n.NodeID, &security.CommandsSupportedGet{})
}

func (n *Node) RequestNodeInformationFrame() error {
	return n.application.serialAPI.RequestNodeInfo(n.NodeID)
}

func (n *Node) LoadCommandClassVersions() error {
	for _, cc := range n.CommandClasses {
		time.Sleep(1 * time.Second)
		cmd := &version.CommandClassGet{RequestedCommandClass: byte(cc.CommandClass)}
		var err error

		if !cc.Secure {
			err = n.application.SendData(n.NodeID, cmd)
		} else {
			err = n.application.SendDataSecure(n.NodeID, cmd)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) LoadManufacturerInfo() error {
	return n.SendCommand(&manufacturerspecific.Get{})
}

func (n *Node) nextQueryStage() {
	if !n.QueryStageSecurity {
		n.LoadSupportedSecurityCommands()
		return
	}

	if !n.QueryStageManufacturer {
		n.LoadManufacturerInfo()
		return
	}

	if !n.QueryStageVersions {
		n.LoadCommandClassVersions()
		return
	}
}

func (n *Node) emitNodeEvent(event commandclass.Command) {
	buf, err := event.MarshalBinary()
	if err != nil {
		fmt.Printf("error encoding: %v\n", err)
		return
	}

	n.application.EventBus.Publish("event", proto.Event{
		Payload: proto.NodeCommandEvent{
			NodeID:         n.NodeID,
			CommandClassID: event.CommandClassID(),
			CommandID:      event.CommandID(),
			CommandData:    buf,
		},
	})
}

func (n *Node) receiveControllerUpdate(update serialapi.ControllerUpdate) {
	n.setFromApplicationControllerUpdate(update)
	n.saveToDb()
}

func (n *Node) setFromAddNodeCallback(nodeInfo *serialapi.AddRemoveNodeCallback) {
	n.NodeID = nodeInfo.Source
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cc := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(commandclass.ID(cc))
	}

	n.saveToDb()
}

func (n *Node) setFromApplicationControllerUpdate(nodeInfo serialapi.ControllerUpdate) {
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cc := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(commandclass.ID(cc))
	}

	n.saveToDb()
}

func (n *Node) setFromNodeProtocolInfo(nodeInfo *serialapi.NodeProtocolInfo) {
	n.Capability = nodeInfo.Capability
	n.BasicDeviceClass = nodeInfo.BasicDeviceClass
	n.GenericDeviceClass = nodeInfo.GenericDeviceClass
	n.SpecificDeviceClass = nodeInfo.SpecificDeviceClass

	n.saveToDb()
}

func (n *Node) receiveSecurityCommandsSupportedReport(cc security.CommandsSupportedReport) {
	for _, cc := range cc.CommandClassSupport {
		n.CommandClasses.SetSecure(commandclass.ID(cc), true)
	}

	// TODO: do we really need to know about controlled command classes?
	// for _, cc := range cc.CommandClassControl {
	// 	n.SecureControlledCommandClasses[commandclass.ID(cc)] = true
	// }

	select {
	case n.queryStageSecurityComplete <- true:
	default:
	}

	n.QueryStageSecurity = true
	n.saveToDb()
	n.nextQueryStage()
}

func (n *Node) receiveManufacturerInfo(mfgInfo manufacturerspecific.Report) {
	n.ManufacturerID = mfgInfo.ManufacturerId
	n.ProductTypeID = mfgInfo.ProductTypeId
	n.ProductID = mfgInfo.ProductId

	select {
	case n.queryStageManufacturerComplete <- true:
	default:
	}

	n.QueryStageManufacturer = true
	n.saveToDb()
	n.nextQueryStage()
}

func (n *Node) receiveCommandClassVersion(id commandclass.ID, version uint8) {
	n.CommandClasses.SetVersion(id, version)

	if n.CommandClasses.AllVersionsReceived() {
		select {
		case n.queryStageVersionsComplete <- true:
		default:
		}

		n.QueryStageVersions = true
		defer n.nextQueryStage()
	}

	n.saveToDb()
}

func (n *Node) receiveApplicationCommand(cmd serialapi.ApplicationCommand) {
	cc := commandclass.ID(cmd.CommandData[0])
	ver := n.CommandClasses.GetVersion(cc)
	if ver == 0 {
		ver = 1

		if !(cc == commandclass.Version || cc == commandclass.Security) {
			fmt.Printf("error: no version loaded for %s\n", cc)
		}
	}

	command, err := commandclass.Parse(ver, cmd.CommandData)
	if err != nil {
		fmt.Println("error parsing command class", err)
		return
	}

	switch command.(type) {

	case *battery.Report:
		if cmd.CommandData[2] == 0xFF {
			fmt.Printf("Node %d: low battery alert\n", n.NodeID)
		} else {
			fmt.Printf("Node %d: battery level is %d\n", n.NodeID, command.(*battery.Report))
		}
		n.emitNodeEvent(command)

	case *security.CommandsSupportedReport:
		fmt.Println("security commands supported report")
		n.receiveSecurityCommandsSupportedReport(*command.(*security.CommandsSupportedReport))
		fmt.Println(n.GetSupportedSecureCommandClassStrings())

	case *manufacturerspecific.Report:
		spew.Dump(command.(*manufacturerspecific.Report))
		n.receiveManufacturerInfo(*command.(*manufacturerspecific.Report))
		n.emitNodeEvent(command)

	case *version.CommandClassReport:
		spew.Dump(command.(*version.CommandClassReport))
		report := command.(*version.CommandClassReport)
		n.receiveCommandClassVersion(commandclass.ID(report.RequestedCommandClass), report.CommandClassVersion)
		n.saveToDb()

	case *versionv2.CommandClassReport:
		spew.Dump(command.(*versionv2.CommandClassReport))
		report := command.(*versionv2.CommandClassReport)
		n.receiveCommandClassVersion(commandclass.ID(report.RequestedCommandClass), report.CommandClassVersion)
		n.saveToDb()

		// case alarm.Report:
		// 	spew.Dump(command.(alarm.Report))
		//
		// case usercode.Report:
		// 	spew.Dump(command.(usercode.Report))
		//
		// case doorlock.OperationReport:
		// 	spew.Dump(command.(doorlock.OperationReport))
		//
		// case thermostatmode.Report:
		// 	spew.Dump(command.(thermostatmode.Report))
		//
		// case thermostatoperatingstate.Report:
		// 	spew.Dump(command.(thermostatoperatingstate.Report))
		//
		// case thermostatsetpoint.Report:
		// 	spew.Dump(command.(thermostatsetpoint.Report))

	default:
		spew.Dump(command)
		n.emitNodeEvent(command)
	}
}

func (n *Node) String() string {
	str := fmt.Sprintf("Node %d: \n", n.NodeID)
	str += fmt.Sprintf("  Failing? %t\n", n.Failing)
	str += fmt.Sprintf("  Is listening? %t\n", n.IsListening())
	str += fmt.Sprintf("  Is secure? %t\n", n.IsSecure())
	str += fmt.Sprintf("  Basic device class: %s\n", n.GetBasicDeviceClassName())
	str += fmt.Sprintf("  Generic device class: %s\n", n.GetGenericDeviceClassName())
	str += fmt.Sprintf("  Specific device class: %s\n", n.GetSpecificDeviceClassName())
	str += fmt.Sprintf("  Manufacturer ID: %#x\n", n.ManufacturerID)
	str += fmt.Sprintf("  Product Type ID: %#x\n", n.ProductTypeID)
	str += fmt.Sprintf("  Product ID: %#x\n", n.ProductID)
	str += fmt.Sprintf("  Supported command classes:\n")

	for _, cc := range n.CommandClasses {
		if cc.Secure {
			str += fmt.Sprintf("    - %s (v%d) (secure)\n", cc.CommandClass.String(), cc.Version)
		} else {
			str += fmt.Sprintf("    - %s (v%d)\n", cc.CommandClass.String(), cc.Version)
		}
	}

	// if len(n.SecureControlledCommandClasses) > 0 {
	// 	secureCommands := commandClassSetToStrings(n.SecureControlledCommandClasses)
	// 	str += fmt.Sprintf("  Controlled command classes (secure):\n")
	// 	for _, cc := range secureCommands {
	// 		str += fmt.Sprintf("    - %s\n", cc)
	// 	}
	// }

	return str
}

func (n *Node) GetSupportedCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(false))
	if len(strings) == 0 {
		return []string{
			"None (probably not loaded; need to request a NIF)",
		}
	}

	return strings
}

func (n *Node) GetSupportedSecureCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(true))
	return strings
}

func commandClassSetToStrings(commandClasses []commandclass.ID) []string {
	if len(commandClasses) == 0 {
		return []string{}
	}

	ccStrings := []string{}

	for _, cc := range commandClasses {
		ccStrings = append(ccStrings, cc.String())
	}

	return ccStrings
}
