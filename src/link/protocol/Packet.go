package protocol

import (
	"strconv"
	"strings"
)

type PacketField struct {
	DataType string
	Value    string
}

type Packet struct {
	PType  string
	Fields []PacketField
}

func ParsePacket(data string) *Packet {
	p := Packet{}

	/* === Raw String Handling === */

	// Remove semicolon from the end of the raw data
	if string(rune(data[len(data)-1])) == ";" {
		data = data[:len(data)-1]
	}

	// Split the data into parts, seperated by a "|"
	parts := strings.Split(data, " | ")

	// Check if the packet has at least one part
	if len(parts) < 1 {
		return nil
	}

	/* === Analysing the raw data === */

	// Get the Packet Type
	p.PType = strings.ToUpper(parts[0])
	parts = parts[1:]

	// Get all other (optional) fields
	for i := 0; i < len(parts); i++ {
		prt := parts[i]

		DataType := ""
		Value := prt

		// Determine the Datatype
		if canBeInt(prt) {
			DataType = "int"
		} else if canBeBool(prt) {
			DataType = "bool"
		} else {
			DataType = "string"
		}

		p.Fields = append(p.Fields, PacketField{DataType, Value})
	}

	return &p
}

func BuildPacket(packetType string, fields []string) string {
	var parts []string

	parts = append(parts, packetType)

	for i := 0; i < len(fields); i++ {
		parts = append(parts, fields[i])
	}

	return strings.Join(parts, " | ") + ";"
}

/*
 * Returns a boolean indicating whether the passed
 * string can be converted to an integer
 */
func canBeInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

/*
 * Returns a boolean indicating whether the string
 * can be converted to a boolean
 */
func canBeBool(str string) bool {
	return strings.ToLower(str) == "false" || strings.ToLower(str) == "true"
}
