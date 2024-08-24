package functions

import (
    "fmt"
    "net"
)

func SocketConnection(information string) string {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err :", err)
		return ""
	}
	defer conn.Close()
	_, err = conn.Write([]byte(information)) 
	if err != nil {
		return ""
	}
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Recv failed:", err)
		return ""
	}
	return string(buf[:n])
}
