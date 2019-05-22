package main

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func main() {
	client, err := gozw.NewDefaultClient("/tmp/data.db", "/dev/tty.usbmodem1431", 115200, networkKey)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	spew.Dump(client.Controller)

	for _, node := range client.Nodes() {
		fmt.Println(node.String())
	}

	fmt.Println("removing node, put device in unpairing mode")
	if _, err := client.RemoveNode(); err != nil {
		log.Fatalf("failed to remove node: %v", err)
	}

	fmt.Println("adding node, put device in pairing mode")
	node, err := client.AddNode()
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	fmt.Println(node.String())
}
