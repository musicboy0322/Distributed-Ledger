package functions

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

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
		fmt.Println("無法開啟檔案:", err)
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
	// 開啟或創建檔案
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("無法創建檔案:", err)
		return
	}
	defer file.Close()

	// 使用 bufio 寫入檔案
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("100")
	if err != nil {
		fmt.Println("寫入檔案時出錯:", err)
		return
	}

	// 刷新緩衝區，確保所有內容都寫入檔案
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
		fmt.Println("無法開啟檔案:", err)
	}
	defer fromFile.Close()

	fromScanner := bufio.NewScanner(fromFile)
	for fromScanner.Scan() {
		text := fromScanner.Text()
		// 將字串轉換為整數
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
		fmt.Println("無法開啟檔案:", err)
	}
	defer toFile.Close()

	toScanner := bufio.NewScanner(toFile)
	for toScanner.Scan() {
		text := toScanner.Text()
		// 將字串轉換為整數
		number, _ := strconv.Atoi(text)
		toFileMoney = number + moneyInt
	}

	toFile, err = os.OpenFile(toFileName, os.O_WRONLY | os.O_TRUNC, 0644)
	toWriter := bufio.NewWriter(toFile)
	_, err = toWriter.WriteString(strconv.Itoa(toFileMoney))
	toWriter.Flush()
	return true
}
