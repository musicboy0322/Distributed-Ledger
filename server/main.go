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
	functions.InitialzeBlocksFolder()
	port := utils.GetServerPort(0)	
	CMD3Channel := make(chan models.CMD3Message, 50)
	CMD5Channel := make(chan models.CMD5Message, 50)
	CMD5BlockCorrectChannel := make(chan bool)
	other_ports := utils.GetOtherPorts(port)

	// open server
	listen, err := net.Listen("tcp", "127.0.0.1:" + port)
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}
	log.Println("Listen to 127.0.0.1:" + port)

	// sending functionality
	services.ConnectNodes(other_ports, CMD3Channel, CMD5Channel, CMD5BlockCorrectChannel)

	// receiving functionality
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept failed: ", err)
			continue
		}
		// handle new connection and basically for short connection
		go services.HandleNewConnection(conn, CMD3Channel, CMD5Channel, CMD5BlockCorrectChannel) 
	}
}