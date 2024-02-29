package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err2 := reader.ReadLine()
		if err2 == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	// createFile("file_sample.go", "package main\nimport \"fmt\"\nfunc main() {\nfmt.Println(\"Hello\")\n}")

	file, err := readFile("file_sample.go")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(file)
	}

	// addToFile("file_sample.go", "\n func SayHello() {\n fmt.Println(\"Hello\")\n}")
}
