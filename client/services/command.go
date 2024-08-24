package services

import (
	"fmt"

	"github.com/Distributed-Ledger/client/utils"
	"github.com/Distributed-Ledger/client/functions"
)

func InitialzeFolder() {
	functions.InitialzeBlocksFolder()
	functions.InitialzeWalletsFolder()
}

func CheckMoney(wallet string) {
	fmt.Print("Enter which wallet: ")
	fmt.Scanln(&wallet)
	if functions.CheckWallet(wallet) == false {
		functions.InitialzeWallet(wallet)
	}
	fmt.Println("Balance: " + functions.CheckBalance(wallet))
}

func CheckLog(wallet string) {
	fmt.Print("Enter which wallet: ")
	fmt.Scanln(&wallet)
	blocks := functions.ListAllBlock()
	fmt.Println("History transitions: " + functions.SearchLog(wallet, blocks))
}

func Transition(fromWallet string, toWallet string, amount string) {
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
	information := "CMD3:" + fromWallet + "," + toWallet + "," + amount
	if functions.TransitMoney(fromWallet, toWallet, amount) == false {
		fmt.Println("Do not have enough money to complete transition")
	} else {
		if functions.CheckFirstBlock() == false {
			functions.InitialzeFirstBlock()
		}
		targetBlock := functions.CheckWriteBlock()
		if functions.CheckBlockMax(targetBlock) == false {
			functions.WriteTransition(fromWallet, toWallet, amount, targetBlock)
			functions.SocketConnection(information)
		} else {
			newTxtName := functions.GetNewTxtName(targetBlock)
			content := functions.RewriteTxt(targetBlock, newTxtName)
			sha256Content := utils.Sha256Encrytion(content)
			functions.InitialzeBlock(newTxtName, sha256Content)
			functions.WriteTransition(fromWallet, toWallet, amount, "./blocks/" + newTxtName)
			functions.SocketConnection(information)
		}
	}
}

func CheckChain() {
	blockSafety := true
	blocks := functions.ListAllBlock()
	if len(blocks) > 1 {
		for i := 0; i < len(blocks); i++ {
			if i == 6 {
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

func CheckAllChain() {
	blocks := functions.ListAllBlock()
	finalBlock := blocks[len(blocks) - 1]
	sha256Content := functions.GetSha256Value(finalBlock)
	information := "CMD5:" + sha256Content
	fmt.Println(functions.SocketConnection(information))
}