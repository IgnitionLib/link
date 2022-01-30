package network_driver

import (
	"ignition-link/src/link"
	"ignition-link/src/link/protocol"

	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type NetworkNode struct {
	Address string
	Driver  *NetworkDriver

	con net.Conn
}

func (this *NetworkNode) Connect() {
	con, err := net.Dial("tcp", this.Address)
	if err != nil {
		log.Fatalln(err)
	}

	this.con = con

	defer con.Close()

	con.Write([]byte("QUERY NODE"))

	serverReader := bufio.NewReader(con)

	for {

		// Waiting for the server response
		serverResponse, err := serverReader.ReadString(';')

		switch err {
		case nil:
			d := strings.TrimSpace(serverResponse)

			p := protocol.ParsePacket(d)

			if strings.ToLower(p.PType) == "node-info" {
				// Extract packet info
				if len(p.Fields) < 4 {
					break
				}

				projectName := p.Fields[0].Value
				projectID := p.Fields[1].Value
				platform := p.Fields[2].Value
				version := p.Fields[3].Value

				// Register Node
				this.Driver.NodeManager.RegisterNode(link.Node{
					Driver:  "network",
					Address: this.Address,

					ProjectName: projectName,
					ProjectId:   projectID,

					Version:  version,
					Platform: platform,
				})
			}

		case io.EOF:
			log.Println("server closed the connection")
			return
		default:
			log.Printf("server error: %v\n", err)
			return
		}
	}
}

func (this *NetworkNode) Send(data string) {
	this.con.Write([]byte(data))
}
