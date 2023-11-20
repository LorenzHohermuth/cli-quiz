package setup

import (
	"bytes"
	"bufio"
	"fmt"
	"os"
)

func DisplayMenu() {
	fmt.Println("Welcome to Cli Quiz an attempt at Making Quizlet better")
	fmt.Println("Please paste this in your Console at the Quizlet Learn Set Website")
	fmt.Println("And paste the output into the Terminal here")
	fmt.Println("\n" + GetFileContent("./scrapingScript.js"))
	fmt.Println(ReadIn())

}

func GetFileContent(path string) string {
	data, err := os.ReadFile(path)
	Check(err)
	return string(data)
}

func Check (err error) {
	if err != nil {
		panic(err)
	}
}

func ReadIn() string {
	var buffer bytes.Buffer
    scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return buffer.String()
}