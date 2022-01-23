package network_driver

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type NetworkNode struct {
	Address string
}

func (this *NetworkNode) Connect() {
	con, err := net.Dial("tcp", this.Address)
	if err != nil {
		log.Fatalln(err)
	}

	defer con.Close()

	con.Write([]byte("QUERY NODE"))

	serverReader := bufio.NewReader(con)

	for {

		// Waiting for the server response
		serverResponse, err := serverReader.ReadString(';')

		switch err {
		case nil:
			log.Println(strings.TrimSpace(serverResponse))
		case io.EOF:
			log.Println("server closed the connection")
			return
		default:
			log.Printf("server error: %v\n", err)
			return
		}
	}
}
