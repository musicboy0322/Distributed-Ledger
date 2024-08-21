package main

import (
	"os"
	"net"
	"fmt"
	"github.com/Distributed-Ledger/server/functions"
)

func main() {

	currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("获取当前目录失败:", err)
        return
    }
	dirPath := currentDir + "/blocks" 
	_, err = os.Stat(dirPath)
    if os.IsNotExist(err) {
		// 目录不存在，创建目录
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("创建目录失败:", err)
			return
		}
		fmt.Println("目录不存在，已成功创建:", dirPath)
    }

    listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go functions.Process(conn) 
	}
}