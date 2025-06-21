package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	const (
		clearScreenEscSeq = "\033[2J\033[1;1H"
	)

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}


	defer term.Restore(fd, oldState)

	fmt.Print(clearScreenEscSeq)
	fmt.Print(">")
	var input []string
	for {
		reader := bufio.NewReader(os.Stdin)
		b, err := reader.ReadByte();
		if err != nil {
			fmt.Printf("ERROR: %v", err)
		}

		if b == 3 { // CTRL-C
			return
		} else if b == 127 { // DELETE KEY
			input = input[:len(input)-1]
		} else {
			input = append(input, string(b))
		}

		fmt.Print(clearScreenEscSeq)
		fmt.Printf("\r>%s", strings.Join(input, ""))
		fmt.Printf("\n\rapple\n\rcherry")
		fmt.Printf("\033[1;%dH", len(input)+2) // Move cursor back to the end of the prompt
	}
}
