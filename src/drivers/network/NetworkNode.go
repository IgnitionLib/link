package network_driver

import (
	"ignition-link/src/link/protocol"

	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type NetworkNode struct {
	Address string

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

			protocol.ParsePacket(d)
			log.Println(d)
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
