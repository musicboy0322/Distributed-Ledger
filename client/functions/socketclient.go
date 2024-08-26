package functions

import (
    "fmt"
    "net"
	"sync"
)

func SocketConnection(port string, information string, wg *sync.WaitGroup, results chan <- string) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "0.0.0.0:" + port)
	if err != nil {
		fmt.Println("err :", err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(information)) 
	if err != nil {
	}
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Recv failed:", err)
	}
	results <- string(buf[:n])
}
