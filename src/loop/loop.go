package loop

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type loop struct {
}

func PrintFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		PrintFileContents(file)
	}
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
