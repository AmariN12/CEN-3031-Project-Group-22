package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"log"
)

type Page struct {
	lines []string
}

func ReadFile(fileName string) Page {
	
	newPage := Page{}
	
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return newPage
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		newPage.lines = append(newPage.lines, scanner.Text())
	}

	return newPage
}

// this will be to read the config file to find where our notes are stored and any other configs we need to save
func ReadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyVal := strings.Split(scanner.Text(), "=")
		fmt.Println(keyVal[0])
		fmt.Println(keyVal[1])
	}
}

func (p *Page) SaveNewPage(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	for _,s := range p.lines {
		_, writeErr := file.WriteString(s)
		if writeErr != nil {
			fmt.Println(writeErr)
		}
		file.WriteString("\n")
	}
}

func main() {

	// need to open a file and save it in the program in some way

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter File to Open:")
	fileName,_ := reader.ReadString('\n')

	fileName = strings.TrimSuffix(fileName, "\n")
	fileName = strings.TrimSuffix(fileName, "\r")

	newPage := ReadFile(fileName)

	for i, element := range newPage.lines {
		fmt.Println(i, element)
	}

	fmt.Println("Enter File to Save to:")
	fileName2,_ := reader.ReadString('\n')

	fileName2 = strings.TrimSuffix(fileName2, "\n")
	fileName2 = strings.TrimSuffix(fileName2, "\r")

	newPage.SaveNewPage(fileName2)
	

	// should probably use a struct and edit it using functions
}