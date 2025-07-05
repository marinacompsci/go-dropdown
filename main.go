package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/marinacompsci/go-dropdown/internal/database"
	"github.com/marinacompsci/go-dropdown/internal/repository"
	"golang.org/x/term"
)

//TODO: create tests

//TODO: separate errors from events(ErrKeyEsc) and put them in a package
var (
	ErrUserInterrupted = errors.New("User pressed CTRL-C")
	ErrEmptyListAsArg = errors.New("Empty list passed as argument")
	ErrEmptyListAsResult = errors.New("Input not found in DB")
	ErrKeyEsc = errors.New("User pressed ESC key")
	ErrKeyDown = errors.New("User pressed J key")
	ErrKeyUp = errors.New("User pressed K key")
)



func main() {
	db := database.ConnectToDB()
	prompt := NewPrompt()
	menu := NewMenu()
	repo := repository.NewExampleRepository(db)
	screen := NewScreen(prompt, menu, repo)


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

		if err := screen.ReadPrompt(b); errors.Is(err, ErrUserInterrupted) {
			return
		}
	}
}
