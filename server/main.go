package main

import (
	"fmt"
	"net"
	"log"
	"github.com/Distributed-Ledger/server/utils"
	"github.com/Distributed-Ledger/server/models"
	"github.com/Distributed-Ledger/server/services"
	"github.com/Distributed-Ledger/server/functions"
)

func main() {
	// initialize variables
	var other_servers []string
	functions.InitialzeBlocksFolder()
	CMD3Channel := make(chan models.CMD3Message, 50)

	port := "8082"
	other_servers = utils.GetOtherServers(port)

	// open server
	listen, err := net.Listen("tcp", "0.0.0.0:" + port)
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}
	log.Println("Node open: 0.0.0.0:" + port)

	// open sending functionality
	services.ConnectNodes(other_servers, CMD3Channel)

	// open receiving functionality
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept failed: ", err)
			continue
		}
		// handle new connection and basically for short connection
		go services.HandleNewConnection(conn, CMD3Channel, len(other_servers)) 
	}
}