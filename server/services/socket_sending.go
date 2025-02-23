package services

import (
    "net"
	"log"
	"fmt"
	"time"
	"github.com/Distributed-Ledger/server/models"
	"github.com/Distributed-Ledger/server/functions"
)

const (
	conncectInterval = 5*time.Second
)

// connect nodes for long connection in sequence
func ConnectNodes(other_ports []string, chcmd3 chan models.CMD3Message, chcmd5 chan models.CMD5Message, chcmd5blockcorrect chan bool) {
	for _, other_port := range other_ports{
		go ConnectNode(other_port, chcmd3, chcmd5, chcmd5blockcorrect)
	}
}

// connect node for long connection
func ConnectNode(other_port string, chcmd3 chan models.CMD3Message, chcmd5 chan models.CMD5Message, chcmd5blockcorrect chan bool) {
	
	node_address := "127.0.0.1:" + other_port
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
	log.Println("Connected to: " + node_address)

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
			fmt.Println(information)
			serializedInformation, err := information.Serialize()
			if err != nil {
				log.Println("Error serializing information:", err)
			}
			fmt.Println(serializedInformation)
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
		case information := <-chcmd5:
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
			n, err = conn.Read(buf[:])
			if err != nil {
				log.Println("Recv failed:", err)
			}
			if n == true {
				chcmd5blockcorrect 
			}
		default:
			// Sleep to avoid busy-waiting
			time.Sleep(20 * time.Second)
		}
	}
}