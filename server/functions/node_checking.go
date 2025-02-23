package functions

import (
	"net"
	"fmt"
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
		fmt.Printf("Server %s is not reachable: %v\n", node_address, err)
		return false
	}
	defer conn.Close()
	return true
}

func CheckNodeAlive(node_address string) {

}