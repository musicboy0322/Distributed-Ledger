package main

import (
	"github.com/Distributed-Ledger/server/utils"
	"github.com/Distributed-Ledger/server/services"
	"github.com/Distributed-Ledger/server/functions"
)

func main() {
	// initialize variables
	functions.InitialzeBlocksFolder()
	port := utils.GetServerPort(0)	
	//other_ports := utils.GetOtherPorts(port)

	services.StartServer(port)
}