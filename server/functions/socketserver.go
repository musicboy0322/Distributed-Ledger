package functions

import (
    "bufio"
    "fmt"
    "net"
	"strings"
	"github.com/Distributed-Ledger/client/utils"
	"github.com/Distributed-Ledger/client/functions"
)

func Process(conn net.Conn) {
	defer conn.Close() 
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) 
		if err != nil {
			fmt.Println("Read from client failed:", err)
			break
		}
		recvStr := string(buf[:n])

		cmd := strings.Split(recvStr, ":")[0]

		switch cmd {
		case "CMD3":
			temp := strings.Split(strings.Split(recvStr, ":")[1], ",")
			fromWallet := temp[0]
			toWallet := temp[1]
			amount := temp[2]
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
			conn.Write([]byte("Write in blockchain sucessfully"))
		case "CMD5":
			clientSha256Content := strings.Split(recvStr, ":")[1]
			blocks := functions.ListAllBlock()
			finalBlock := blocks[len(blocks) - 1]
			sha256Content := functions.GetSha256Value(finalBlock)
			if clientSha256Content == sha256Content {
				conn.Write([]byte("correct"))
			} else {
				conn.Write([]byte("wrong"))
			}
		}
	}	
}