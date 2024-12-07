package functions

import (
    "fmt"
    "net"
	"strconv"
)

func SocketConnection(ports []int, information string) bool {
	for _, port := range ports {
		portString := strconv.Itoa(port)
		conn, err := net.Dial("tcp", "0.0.0.0:" + portString)
		if err != nil {
			fmt.Println("err :", err)
			continue
		}
		defer conn.Close()
		_, err = conn.Write([]byte(information)) 
		if err != nil {
		}
		buf := [512]byte{}
		_, err = conn.Read(buf[:])
		if err != nil {
			fmt.Println("Recv failed:", err)
			return false
		}
		return true
	}
	fmt.Println("None entry node avaliable")
	return false
}
