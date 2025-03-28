package functions

import (
    "fmt"
    "net"
	"strings"
	"github.com/Distributed-Ledger/client/models"
)

func SocketConnection(server string, information models.CommandMessage) string {
	// connect to target port
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("err :", err)
	}
	defer conn.Close()
	// serialize information
	serializedInformation, err := information.Serialize()
	if err != nil {
		fmt.Println("Error serializing information:", err)
		return "false"
	}
	// send information to target port
	_, err = conn.Write(serializedInformation) 
	if err != nil {
		fmt.Println("err :", err)
	}
	// Read all data
	var response strings.Builder
	buf := make([]byte, 1024) // Increase buffer size
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		response.Write(buf[:n]) // Append to response
	}
	return response.String()
}
