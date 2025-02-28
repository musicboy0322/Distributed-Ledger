package services

import (
	"fmt"
	"github.com/Distributed-Ledger/client/utils"
	"github.com/Distributed-Ledger/client/models"
	"github.com/Distributed-Ledger/client/functions"
)

func InitialzeFolder() {
	functions.InitialzeBlocksFolder()
	functions.InitialzeWalletsFolder()
}

func CheckMoney(wallet string) {
	// get user message
	fmt.Print("Enter which wallet: ")
	fmt.Scanln(&wallet)
	if functions.CheckWallet(wallet) == false {
		functions.InitialzeWallet(wallet)
	}
	balance := functions.CheckBalance(wallet)
	// display result
	fmt.Println("Balance: " + balance)
}

func CheckLog(wallet string, ports []int) {
	// get user message
	fmt.Print("Enter which wallet: ")
	fmt.Scanln(&wallet)
	// set socket connect message
	information := models.CMD2Message {
		Category: "SC",
		Command: "CMD2",
		Wallet: wallet,
	}
	// random choose a port for connecting
	port := utils.GetRandomPort(ports)
	// socket connect
	result := functions.SocketConnection(port, information)
	fmt.Println("History transitions: " + result)	
}

func Transition(fromWallet string, toWallet string, amount string, ports []int) {
	// procedure of asking information
	fmt.Print("Enter which wallet to use: ")
	fmt.Scanln(&fromWallet)
	if functions.CheckWallet(fromWallet) == false {
		functions.InitialzeWallet(fromWallet)
	}
	fmt.Print("Enter which wallet to give: ")
	fmt.Scanln(&toWallet)
	if functions.CheckWallet(toWallet) == false {
		functions.InitialzeWallet(toWallet)
	}
	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	information := models.CMD3Message {
		Category: "SC",
		Command: "CMD3",
		FromWallet: fromWallet,
		ToWallet: toWallet,
		Amount: amount,
	}

	// random choose a port for connecting
	port := utils.GetRandomPort(ports)
	// process
	if functions.TransitMoney(fromWallet, toWallet, amount) == false {
		fmt.Println("Do not have enough money to complete transition")
	} else {
		result := functions.SocketConnection(port, information)
		if result == "false" {
			fmt.Println("Fail to write in block")
		}
		fmt.Println("Sucessfully write in block")
	}
}

/*
func CheckChain() {
	blockSafety := true
	blocks := functions.ListAllBlock()
	if len(blocks) > 1 {
		for i := 0; i < len(blocks); i++ {
			if i == len(blocks) - 1 {
				break
			}
			content := functions.GetAllBlockContent(blocks[i])
			sha256Content := utils.Sha256Encrytion(content)
			if functions.CheckSha256(blocks[i+1], sha256Content) == false {
				blockSafety = false
				break
			}
		}
	}
	if blockSafety {
		fmt.Println("Chain safe!")
	} else {
		fmt.Println("Dangerous!Chain has been changed!") 
	}
}
*/

/*
func CheckAllChain(ports []int) {
	// process
	blocks := functions.ListAllBlock()
	if len(blocks) == 1 {
		fmt.Println("Only has one block now")
		return
	}
	finalBlock := blocks[len(blocks) - 1]
	sha256Content := functions.GetSha256Value(finalBlock)

	information := models.CMD5Message {
		Category: "SC",
		Command: "CMD5",
		Sha256Content: sha256Content,
	}

	port := utils.GetRandomPort(ports)

	result := functions.SocketConnection(port, information)
	if result == false {
		fmt.Println("Fail to check other block or some block been changed")
	}
	fmt.Println("All block safe")
}
*/