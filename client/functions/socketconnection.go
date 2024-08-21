package functions

import (
    "fmt"
    "net"
)

func SocketConnection(transitionInformation string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close()
	for {
		_, err = conn.Write([]byte(transitionInformation)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
