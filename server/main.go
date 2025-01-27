package main

import (
	"net"
	"fmt"
	"log"
	"github.com/Distributed-Ledger/server/utils"
	"github.com/Distributed-Ledger/server/services"
	"github.com/Distributed-Ledger/server/functions"
)

func main() {
	// initialize variables
	functions.InitialzeBlocksFolder()
	port := utils.GetServerPort(0)	
	//other_ports := utils.GetOtherPorts(port)

	// start server
    listen, err := net.Listen("tcp", "0.0.0.0:" + port)
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}
	log.Println("Listen to 0.0.0.0:" + port)
	for {
		// receive new connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}
		// handle new connection and basically for short connection
		go services.HandleNewConnection(conn) 
	}
}