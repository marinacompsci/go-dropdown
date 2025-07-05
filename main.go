package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/marinacompsci/go-dropdown/internal/database"
	"github.com/marinacompsci/go-dropdown/internal/repository"
	"github.com/marinacompsci/go-dropdown/internal/exception"
	"golang.org/x/term"
)


func main() {
	db := database.ConnectToDB()
	repo := repository.NewExampleRepository(db)
	screen := NewScreen(repo)


	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}


	defer term.Restore(fd, oldState)

	screen.PreRender()

	for {
		reader := bufio.NewReader(os.Stdin)
		b, err := reader.ReadByte();
		if err != nil {
			fmt.Printf("ERROR: %v", err)
			return
		}

		if err := screen.ReadPrompt(b); errors.Is(err, exception.ErrUserInterrupted) {
			return
		}
	}
}
