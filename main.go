package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)


var (
	ErrUserInterrupted = errors.New("User pressed CTRL-C")
	ErrEmptyListAsArg = errors.New("Empty list passed as argument")
	ErrEmptyListAsResult = errors.New("Input not found in DB")
)


func main() {
	Data := []string{"maçã", "banana", "cereja", "laranja", "uva", "limão", "manga", "abacaxi", "pêra", "pêssego",
	"morango", "framboesa", "mirtilo", "amora", "melancia", "melão", "kiwi", "papaia", "coco", "figo",
	"tâmara", "romã", "caqui", "goiaba", "maracujá", "acerola", "cajá", "pitanga", "jabuticaba", "caju",
	"cupuaçu", "açaí", "bacuri", "buriti", "cagaita", "cambucá", "cambuci", "camu-camu", "cajuí", "carambola",
	"ciriguela", "fruta-pão", "graviola", "guaraná", "ingá", "jaca", "jambo", "jenipapo", "juá", "licuri",
	"mangaba", "murici", "pequi", "pitomba", "seriguela", "umbu", "uvaia", "abacate", "ameixa", "damasco",
	"nêspera", "tangerina", "toranja", "lima", "cidra", "bergamota", "mexerica", "ponkan", "kinkan", "yuzu",
	"physalis", "cranberry", "goji", "amêndoa", "avelã", "castanha", "noz", "pistache", "macadâmia", "pecã",
	"jujuba", "lichia", "longan", "rambutan", "durião", "mangostão", "fruta-dragão", "atemoia", "pinha", "fruta-conde",
	"sapoti", "cajamanga", "abiu", "bacupari", "biribá", "araçá", "grumixama", "cambuí", "gabiroba", "marmelada"}

	prompt := NewPrompt()
	menu := NewMenu()
	screen := NewScreen(prompt, menu, Data)


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
			continue
		}

		if err := screen.ReadPrompt(b); errors.Is(err, ErrUserInterrupted) {
			return
		}
	}
}
