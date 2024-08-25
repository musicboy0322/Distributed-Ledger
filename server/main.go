package main

import (
	"os"
	"net"
	"fmt"
	"log"
	"github.com/Distributed-Ledger/server/functions"
	"github.com/Distributed-Ledger/server/utils"
)

func main() {

	port := utils.GetServerPort()	

	currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Fail to get current directory:", err)
        return
    }
	dirPath := currentDir + "/blocks" 
	_, err = os.Stat(dirPath)
    if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Fail to create file:", err)
			return
		}
    }
    listen, err := net.Listen("tcp", "127.0.0.1:" + port)
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}
	log.Println("Listen to 127.0.0.1:" + port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}
		go functions.Process(conn) 
	}
}