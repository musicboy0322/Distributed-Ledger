package functions

import (
	"net"
	"log"
	"time"
)

const (
	avaliableTimeout = 15*time.Second
	heartbeatInterval = 15*time.Second
	heartbeatTimeout = 5*time.Second
)

func CheckNodeAvaliable(node_address string) bool {
	conn, err := net.DialTimeout("tcp", node_address, avaliableTimeout)
	if err != nil {
		log.Printf("Long connection not reachable: %s", node_address)
		return false
	}
	defer conn.Close()
	conn.Write([]byte("Checking node avaliable or not"))
	return true
}

func CheckNodeAlive(node_address string) {

}