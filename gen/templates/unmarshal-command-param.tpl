{{if eq .Type "VARIANT"}}
    {{template "unmarshal-variant" .}}
  {{else if eq .Type "STRUCT_BYTE"}}{{$name := ToGoName .Name}}
{{ template "exists" .}}
    {{range .BitField}}
      {{if .IsNotReserved}}
        cmd.{{$name}}.{{ToGoName .FieldName}} = (payload[i]{{with .FieldMask}}&{{.}}{{end}}){{with .Shifter}}>>{{.}}{{end}}
      {{end}}
    {{end}}
    {{range .FieldEnum}}
      cmd.{{$name}}.{{ToGoName .FieldName}} = (payload[i]{{with .FieldMask}}&{{.}}{{end}}){{with .Shifter}}>>{{.}}{{end}}
    {{end}}
    {{range .BitFlag}}
      {{if .IsNotReserved}}
        cmd.{{$name}}.{{ToGoName .FlagName}} = payload[i] & {{.FlagMask}} == {{.FlagMask}}
      {{end}}
    {{end}}
    i += 1

  {{else if eq .Type "ARRAY"}}
{{ template "exists" .}}
    {{if (index .ArrayAttrib 0).IsAscii}}
      cmd.{{ToGoName .Name}} = string(payload[i:i+{{(index .ArrayAttrib 0).Length}}])
    {{else}}
      cmd.{{ToGoName .Name}} = payload[i:i+{{(index .ArrayAttrib 0).Length}}]
    {{end}}
    i += {{(index .ArrayAttrib 0).Length}}
  {{else if eq .Type "BITMASK"}}
{{ template "exists" .}}
    cmd.{{ToGoName .Name}} = payload[i:]
  {{else if eq .Type "DWORD"}}
{{ template "exists" .}}
    cmd.{{ToGoName .Name}} = binary.BigEndian.Uint32(payload[i:i+4])
    i += 4
  {{else if eq .Type "BIT_24"}}
{{ template "exists" .}}
    cmd.{{ToGoName .Name}} = binary.BigEndian.Uint32(payload[i:i+3])
    i += 3
  {{else if eq .Type "WORD"}}
    {{if not .IsOptional}}
    if len(payload) <= i {
      return errors.New("slice index out of bounds")
    }
    {{end}}

    cmd.{{ToGoName .Name}} = binary.BigEndian.Uint16(payload[i:i+2])
    i += 2
  {{else if eq .Type "MARKER"}}
{{ template "exists" .}}
    i += 1 // skipping MARKER
    if len(payload) <= i {
      return nil
    }
  {{else}}
  
{{ template "exists" .}}
    {{if .IsNotReserved}}
      cmd.{{ToGoName .Name}} = payload[i]
      i++
    {{else}}
      i++ // skipping over reserved bit
    {{end}}

  {{end}}

{{ define "exists" }}
    if len(payload) <= i {
    {{if .IsOptional -}}
      return nil // field is optional
    {{- else -}}
      return errors.New("slice index out of bounds")
    {{- end}}
    }
{{end}}