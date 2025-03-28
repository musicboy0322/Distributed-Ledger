package services

import (
    "net"
	"log"
	"time"
	"github.com/Distributed-Ledger/server/models"
	"github.com/Distributed-Ledger/server/functions"
)

const (
	conncectInterval = 5*time.Second
)

// connect nodes for long connection in sequence
func ConnectNodes(other_servers []string, chcmd3 chan models.CMD3Message) {
	for _, other_server := range other_servers{
		go ConnectNode(other_server, chcmd3)
	}
}

// connect node for long connection
func ConnectNode(other_server string, chcmd3 chan models.CMD3Message) {
	
	node_address := other_server
	node_status := false

	// checking node is available to connect or not
	for node_status == false {
		node_status = functions.CheckNodeAvaliable(node_address)
		time.Sleep(conncectInterval)
	}

	// Connect to the server
	conn, err := net.Dial("tcp", node_address)
	if err != nil {
		log.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	log.Println("Long connection to: " + node_address)

	information := models.INITMessage {"LC", "INIT"}
	serializedInformation, err := information.Serialize()
	if err != nil {
		log.Println("Error serializing information:", err)
	}
	_, err = conn.Write(serializedInformation) 
	if err != nil {
		log.Println("err :", err)
	}
	buf := [512]byte{}
	_, err = conn.Read(buf[:])
	if err != nil {
		log.Println("Recv failed:", err)
	}

	for {
		select {
		case information := <-chcmd3:
			serializedInformation, err := information.Serialize()
			if err != nil {
				log.Println("Error serializing information:", err)
			}
			// send information to target port
			_, err = conn.Write(serializedInformation) 
			if err != nil {
				log.Println("err :", err)
			}
			buf := [512]byte{}
			_, err = conn.Read(buf[:])
			if err != nil {
				log.Println("Recv failed:", err)
			}
			time.Sleep(30 * time.Second)
		default:
			// Sleep to avoid busy-waiting
			time.Sleep(20 * time.Second)
		}
	}
}