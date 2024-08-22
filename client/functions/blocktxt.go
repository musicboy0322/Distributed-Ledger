package functions

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

func InitialzeBlocksFolder() {
	// gather current directory
	currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Fail to gather current directory:", err)
        return
    }
	// set target folder path
	dirPath := currentDir + "/blocks" 
	_, err = os.Stat(dirPath)
	// detect target folder existence
    if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Fail to create folder:", err)
			return
		}
    }
}

func CheckFirstBlock() bool {
	_, err := os.Stat("./blocks/1.txt")
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func InitialzeFirstBlock() {
	// set first block content
	message := "Sha256 of previous block:\nNext block:"
	// create first block
	file, err := os.Create("./blocks/1.txt")
	if err != nil {
		fmt.Println("Fail to create file:", err)
		return
	}
	defer file.Close()
	// use bufio to write data in file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(message)
	if err != nil {
		fmt.Println("Encounter error when wrting data in file:", err)
		return
	}
	// refresh buffer to confirm all data write in file
	writer.Flush()
}

func CheckWriteBlock() string {
	// set directory path and variable to detect needed write block
	dirPath := "./blocks"
	targetTxt := false
	// read all files in folder blocks
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Can not read dictionary")
	}
	// loop all files in folder blocks
	for _, file := range files {
		fileName := dirPath + "/" + file.Name()
		block, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Fail to open file:", err)
		}
		defer block.Close()
		scanner := bufio.NewScanner(block)
		// filter target block to write
		lineNumber := 1
		for scanner.Scan() {
			if lineNumber == 2 {
				text := scanner.Text()
				parts := strings.Split(text, " ")
				if len(parts) != 3 {
					targetTxt = true
				}
			}
			lineNumber++
		}
		if targetTxt {
			return fileName
		}
	}
	return ""
}

func WriteTransition(fromWallet string, toWallet string, money string, blockFile string) {
	// use append setting to open target file
	file, err := os.OpenFile(blockFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Fail to open file:", err)
        return
    }
    defer file.Close()
	// use bufio to write data in file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("\n" + fromWallet + ", " + toWallet + ", " + money)
	if err != nil {
		fmt.Println("Encounter error when writing data in file:", err)
		return
	}
	// refresh buffer to confirm all data write in file
	writer.Flush()
}

func CheckBlockMax(blockFile string) bool {
	// open target file
	block, err := os.Open(blockFile)
	if err != nil {
		fmt.Println("無法開啟檔案:", err)
	}
	defer block.Close()
	scanner := bufio.NewScanner(block)
	// detect target block has enough position to write in data
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
	}
	if lineNumber == 7 {
		return true
	}
	return false
}

func GetNewTxtName(blockFile string) string {
	temp := strings.Split(blockFile, "/")[2]
	number := strings.Split(temp, ".")[0]
	numberInt, _ := strconv.Atoi(number)
	numberInt++
	number = strconv.Itoa(numberInt)
	return number + ".txt"
}

func RewriteTxt(blockFile string, newTxtName string) string {
	block, err := os.Open(blockFile)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer block.Close()
	// variable content to fill in all the content in target block
	var content string
	lineNumber := 1
	scanner := bufio.NewScanner(block)
	for scanner.Scan() {
		if lineNumber == 1 {
			content = scanner.Text()
			lineNumber++
			continue
		}
		if lineNumber == 2 {
			content = content + "\n" + "Next block: " + newTxtName
			lineNumber++
			continue
		}
		content = content + "\n" + scanner.Text()
		lineNumber++
	}
	file, err := os.OpenFile(blockFile, os.O_WRONLY | os.O_TRUNC, 0644)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	writer.Flush()
	return content
}

func InitialzeBlock(newTxtName string, sha256Content string) {
	message := "Sha256 of previous block: " + sha256Content + "\nNext block:"
	file, err := os.Create("./blocks/" + newTxtName)
	if err != nil {
		fmt.Println("Faile to open file:", err)
		return
	}
	defer file.Close()
	// use bufio to write data in file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(message)
	if err != nil {
		fmt.Println("寫入檔案時出錯:", err)
		return
	}
	// refresh buffer to confirm all data write in file
	writer.Flush()
}

func SearchLog(wallet string, blocks []string) string {
	var historyLog string
	for _, block := range blocks {
		blockFile, err := os.Open(block)
		if err != nil {
			fmt.Println("Fail to open file:", err)
		}
		defer blockFile.Close()

		scanner := bufio.NewScanner(blockFile)

		lineNumber := 1
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), wallet) {
				historyLog = historyLog + "\n" + scanner.Text()
			}
			lineNumber++
		}
	}
	return historyLog
}

func ListAllBlock() []string {
	var blocks []string
	var nextBlock string = "./blocks/1.txt"
	var nextBlockName string
	for true {
		blocks = append(blocks, nextBlock)
		block, err := os.Open(nextBlock)
		if err != nil {
			fmt.Println("Fail to open file:", err)
		}
		defer block.Close()
		scanner := bufio.NewScanner(block)
		lineNumber := 1
		for scanner.Scan() {
			if lineNumber == 2 {
				nextBlockName = strings.Split(scanner.Text(), ":")[1]
				nextBlockName = strings.TrimSpace(nextBlockName)
				lineNumber++
				continue
			}
			lineNumber++
		}
		if nextBlockName == "" {
			break
		}
		nextBlock = "./blocks/" + nextBlockName
	}
	return blocks
}

func GetAllBlockContent(blockFileName string) string {
	var content string
	block, err := os.Open(blockFileName)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer block.Close()
	scanner := bufio.NewScanner(block)
	lineNumber := 1
	for scanner.Scan() {
		if lineNumber == 1 {
			content = scanner.Text()
			lineNumber++
			continue
		}
		content = content + "\n" + scanner.Text()
		lineNumber++
	}
	return content
}

func CheckSha256(nextBlockFile string, sha256Content string) bool {
	block, err := os.Open(nextBlockFile)
	if err != nil {
		fmt.Println("Fail to open file:", err)
	}
	defer block.Close()
	scanner := bufio.NewScanner(block)
	for scanner.Scan() {
		if strings.Split(scanner.Text(), " ")[4] == sha256Content {
			return true
		}
		break
	}
	return false
}