package functions

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func InitialzeWalletsFolder() {
	currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Fail to get current directory:", err)
        return
    }
	dirPath := currentDir + "/wallets" 
	_, err = os.Stat(dirPath)
    if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Fail to create file:", err)
			return
		}
    }
}

func CheckWallet(wallet string) bool {
	wallet = "./wallets/" + wallet + ".txt"
	_, err := os.Stat(wallet)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func CheckBalance(wallet string) string {
	fileName := "./wallets/" + wallet + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func InitialzeWallet(wallet string) {
	fileName := "./wallets/" + wallet + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fail to open file:", err)
		return
	}
	defer file.Close()
	// use bufio to write data in file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("100")
	if err != nil {
		fmt.Println("Encounter error when writing in data", err)
		return
	}
	// refresh buffer to confirm all data write in file
	writer.Flush()
}

func TransitMoney(fromWallet string, toWallet string, money string) bool {
	var fromFileMoney int
	var toFileMoney int
	var moneyInt int
	moneyInt, _ = strconv.Atoi(money)
	fromFileName := "./wallets/" + fromWallet + ".txt"
	toFileName := "./wallets/" + toWallet + ".txt"
	// from file
	fromFile, err := os.Open(fromFileName)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer fromFile.Close()

	fromScanner := bufio.NewScanner(fromFile)
	for fromScanner.Scan() {
		text := fromScanner.Text()
		// turn string into int
		number, _ := strconv.Atoi(text)
		fromFileMoney = number - moneyInt
		if fromFileMoney < 0 {
			return false
		}
	}
	fromFile, err = os.OpenFile(fromFileName, os.O_WRONLY | os.O_TRUNC, 0644)
	fromWriter := bufio.NewWriter(fromFile)
	_, err = fromWriter.WriteString(strconv.Itoa(fromFileMoney))
	fromWriter.Flush()
	// to file
	toFile, err := os.Open(toFileName)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer toFile.Close()
	toScanner := bufio.NewScanner(toFile)
	for toScanner.Scan() {
		text := toScanner.Text()
		// turn string into int
		number, _ := strconv.Atoi(text)
		toFileMoney = number + moneyInt
	}
	toFile, err = os.OpenFile(toFileName, os.O_WRONLY | os.O_TRUNC, 0644)
	toWriter := bufio.NewWriter(toFile)
	_, err = toWriter.WriteString(strconv.Itoa(toFileMoney))
	toWriter.Flush()
	return true
}
