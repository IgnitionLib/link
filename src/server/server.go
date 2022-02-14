package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = "localhost:47113"
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func gateway(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error - upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("error - read:", err)
			break
		}

		log.Printf("recv: %s", message)

		// Parse packet
		var packet Packet
		err = json.Unmarshal(message, &packet)

		if err != nil {
			c.WriteMessage(mt, []byte(`{}`))
		}

		fmt.Println(packet)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("error - write:", err)
			break
		}
	}
}

func StartServer() {
	log.SetFlags(0)
	http.HandleFunc("/", gateway)
	log.Fatal(http.ListenAndServe(addr, nil))
}
