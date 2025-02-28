package services

import (
    "net"
	"log"
	"fmt"
	"time"
	"encoding/json"
	"github.com/Distributed-Ledger/server/utils"
	"github.com/Distributed-Ledger/server/models"
	"github.com/Distributed-Ledger/server/functions"
)

func HandleNewConnection(conn net.Conn, chcmd3 chan models.CMD3Message, node_amount int) {
	defer conn.Close()
	// get remote ip address
	remoteAddr := conn.RemoteAddr().String()
	// read initial message
	buf := make([]byte, 1024)  // Adjust the buffer size based on your message size
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("reading message:", err)
	}
	// deserialize json format for detecting long connection or short connection
	var temp map[string]interface{}
    _ = json.Unmarshal(buf[:n], &temp)
	// Check if the "Command" field is present and is a string
	command, _ := temp["Command"].(string)

	// Check if the "Category" field is present and is a string
	category, _ := temp["Category"].(string)

	// detect long connection or short connection
	if category == "LC" {
		log.Println("Long connection from: " + remoteAddr)
		HandleLongConnection(conn)
	} else if category == "SC" {
		log.Println("Short connection: " + remoteAddr + ", about: " + command)
		switch command {
		case "CMD2":
			var currentMessage models.CMD2Message
			_ = json.Unmarshal([]byte(buf[:n]), &currentMessage)
			HandleCMD2(conn, currentMessage)
		case "CMD3":
			var currentMessage models.CMD3Message
			_ = json.Unmarshal([]byte(buf[:n]), &currentMessage)
			HandleCMD3(conn, currentMessage)
			currentMessage.Category = "LC"
			for i := 0; i < node_amount; i++ {
				chcmd3 <- currentMessage
			}
	}
	}
}

func HandleLongConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		n, _ := conn.Read(buf)

		// deserialize json format for detecting long connection or short connection
		var temp map[string]interface{}
		_ = json.Unmarshal(buf[:n], &temp)
		// Check if the "Command" field is present and is a string
		command, _ := temp["Command"].(string)
		switch command {
		case "CMD3":
			var currentMessage models.CMD3Message
			_ = json.Unmarshal([]byte(buf[:n]), &currentMessage)
			HandleCMD3(conn, currentMessage)
		}
		// 回應數據
		_, err := conn.Write([]byte("Server Acknowledged: "))
		if err != nil {
			fmt.Println("Error writing to long connection:", err)
			break
		}
	}
}

func HandleCMD2(conn net.Conn, message models.CMD2Message) {
	wallet := message.Wallet
	blocks := functions.ListAllBlock()
	transitionLog := functions.SearchLog(wallet, blocks)
	conn.Write([]byte(transitionLog))
}

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

/*
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
*/