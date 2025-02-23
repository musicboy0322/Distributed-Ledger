package main

import (
	"fmt"
	"os"

	"github.com/Distributed-Ledger/client/services"
	"github.com/Distributed-Ledger/client/utils"
)

func main() {
	var (
		command string
		wallet string
		fromWallet string
		toWallet string
		amount string
	)

	services.InitialzeFolder()
	ports := utils.GetEnterPorts()

	for true {
		fmt.Print("Enter a number (1)Check Money (2)Check Log (3)Transition (4)Quit : ")
		fmt.Scanln(&command)
		switch command {
		case "1":
			services.CheckMoney(wallet)
		case "2":
			services.CheckLog(wallet, ports)
		case "3":
			services.Transition(fromWallet, toWallet, amount, ports)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Do not have this command, please check and enter right number")
		}
	}
}