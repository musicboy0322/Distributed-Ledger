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
		temp := strings.Split(recvStr, ",")
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
		conn.Write([]byte("ok"))
	}
}