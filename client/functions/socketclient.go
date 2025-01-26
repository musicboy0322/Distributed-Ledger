package functions

import (
    "fmt"
    "net"
	"strconv"
	"github.com/Distributed-Ledger/client/models"
)

func SocketConnection(port int, information models.CommandMessage) bool {
	// turn port data type from int to string
	portString := strconv.Itoa(port)
	// connect to target port
	conn, err := net.Dial("tcp", "0.0.0.0:" + portString)
	if err != nil {
		fmt.Println("err :", err)
	}
	defer conn.Close()
	// serialize information
	serializedInformation, err := information.Serialize()
	if err != nil {
		fmt.Println("Error serializing information:", err)
		return false
	}
	// send information to target port
	_, err = conn.Write(serializedInformation) 
	if err != nil {
		fmt.Println("err :", err)
	}
	buf := [512]byte{}
	_, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("Recv failed:", err)
		return false
	}
	return true
}
