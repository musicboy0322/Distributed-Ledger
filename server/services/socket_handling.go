package services

import (
    "net"
	"log"
	"fmt"
	"encoding/json"
	"github.com/Distributed-Ledger/server/utils"
	"github.com/Distributed-Ledger/server/models"
	"github.com/Distributed-Ledger/server/functions"
)

//storage_queue chan string
func HandleNewConnection(conn net.Conn) {
	defer conn.Close()
	// get remote ip address
	remoteAddr := conn.RemoteAddr().String()
	// read initial message
	buf := make([]byte, 1024)  // Adjust the buffer size based on your message size
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}

	// deserialize json format for detecting long connection or short connection
	var temp map[string]interface{}
    err = json.Unmarshal(buf[:n], &temp)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
	// Check if the "Command" field is present and is a string
	command, ok := temp["Command"].(string)
	if !ok {
		fmt.Println("Error: Command field is missing or not a string")
		return
	}
	
	// Check if the "Category" field is present and is a string
	category, ok := temp["Category"].(string)
	if !ok {
		fmt.Println("Error: Category field is missing or not a string")
		return
	}

	
	// detect long connection or short connection
	if category == "LC" {
		fmt.Println("long connection: " + remoteAddr)
		//functions.HandleLongConnection(conn)
	} else if category == "SC" {
		fmt.Println("short connection: " + remoteAddr + ", about: " + command)
		switch command {
		case "CMD3":
			var currentMessage models.CMD3Message
			err = json.Unmarshal([]byte(buf[:n]), &currentMessage)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			HandleCMD3(conn, currentMessage)
		case "CMD5":
			var currentMessage models.CMD5Message
			err = json.Unmarshal([]byte(buf[:n]), &currentMessage)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			HandleCMD5(conn, currentMessage)
		}
	}
}

/*
func HandleLongConnection(conn net.Conn) {
	for {
		// receive message
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Long connection closed:", err)
			break
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Received from long connection: %s\n", message)

		// 回應數據
		_, err = conn.Write([]byte("Server Acknowledged: " + message + "\n"))
		if err != nil {
			fmt.Println("Error writing to long connection:", err)
			break
		}
	}
}
	*/

func HandleCMD3(conn net.Conn, message models.CMD3Message) {
	fromWallet := message.FromWallet
	toWallet := message.ToWallet
	amount := message.Amount
	log.Println("Write in data: " + fromWallet + ", " + toWallet + ", " + amount)
	if functions.CheckFirstBlock() == false {
		functions.InitialzeFirstBlock()
	}
	targetBlock := functions.CheckWriteBlock()
	if functions.CheckBlockMax(targetBlock) == false {
		functions.WriteTransition(fromWallet, toWallet, amount, targetBlock)
	} else {
		newTxtName := functions.GetNewTxtName(targetBlock)
		content := functions.RewriteTxt(targetBlock, newTxtName)
		sha256Content := utils.Sha256Encrytion(content)
		functions.InitialzeBlock(newTxtName, sha256Content)
		functions.WriteTransition(fromWallet, toWallet, amount, "./blocks/" + newTxtName)
	}
	conn.Write([]byte("true"))
}

func HandleCMD5(conn net.Conn, message models.CMD5Message) {
	clientSha256Content := message.Sha256Content
	blocks := functions.ListAllBlock()
	finalBlock := blocks[len(blocks) - 1]
	sha256Content := functions.GetSha256Value(finalBlock)
	log.Println("Client last sha256:" + clientSha256Content + " Server last256:" + sha256Content)
	if clientSha256Content == sha256Content {
		conn.Write([]byte("true"))
	} else {
		conn.Write([]byte("false"))
	}
}