package server

type Packet struct {
	PacketType string     `json:"type"`
	Data       PacketData `json:"data"`
}

type PacketData struct {
}
