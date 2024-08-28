package learn_golang_std_lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBufIO(t *testing.T) {

	reader := strings.NewReader("Hallo, Joko\ntest test")

	bufioReader := bufio.NewReader(reader)

	for {
		line, _, err := bufioReader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Hello, test test\n")

	writer.WriteString("Hello2, test2 test2\n")

	writer.Flush()
}
