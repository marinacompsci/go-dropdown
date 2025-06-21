package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/term"
)

func main() {
	const (
		//TODO: Separate 'move' from '?' (UI stuff)
		clearScreenEscSeq = "\033[2J\033[1;1H" // \033[1:1H is move
	)

	var input []string
	var searchList []string

	searchList = append(searchList, "maçã", "banana", "cereja", "laranja", "uva", "limão", "manga", "abacaxi", "pêra", "pêssego",
    "morango", "framboesa", "mirtilo", "amora", "melancia", "melão", "kiwi", "papaia", "coco", "figo",
    "tâmara", "romã", "caqui", "goiaba", "maracujá", "acerola", "cajá", "pitanga", "jabuticaba", "caju",
    "cupuaçu", "açaí", "bacuri", "buriti", "cagaita", "cambucá", "cambuci", "camu-camu", "cajuí", "carambola",
    "ciriguela", "fruta-pão", "graviola", "guaraná", "ingá", "jaca", "jambo", "jenipapo", "juá", "licuri",
    "mangaba", "murici", "pequi", "pitomba", "seriguela", "umbu", "uvaia", "abacate", "ameixa", "damasco",
    "nêspera", "tangerina", "toranja", "lima", "cidra", "bergamota", "mexerica", "ponkan", "kinkan", "yuzu",
    "physalis", "cranberry", "goji", "amêndoa", "avelã", "castanha", "noz", "pistache", "macadâmia", "pecã",
    "jujuba", "lichia", "longan", "rambutan", "durião", "mangostão", "fruta-dragão", "atemoia", "pinha", "fruta-conde",
    "sapoti", "cajamanga", "abiu", "bacupari", "biribá", "araçá", "grumixama", "cambuí", "gabiroba", "marmelada")

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}


	defer term.Restore(fd, oldState)

	fmt.Print(clearScreenEscSeq)
	fmt.Print(">")
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

		inputStr := strings.Join(input, "")

		fmt.Print(clearScreenEscSeq)
		fmt.Printf("\r>%s", strings.Join(input, ""))

		var resultList []string
		searchTokenInList(inputStr, searchList, &resultList)
		if len(input) > 0 {
			//TODO: save "\n\r" in constant
			fmt.Print("\n\r" + stringifyList(resultList, "\n\r"))
		}

		fmt.Printf("\033[1;%dH", len(input)+2) // Move cursor back to the end of the prompt
	}
}


//TODO: put those functions in a separate package (helper?)
func stringifyList(l []string, sep string) string {
	if len(l) == 0 { return "" }

	return strings.Join(l, sep)
}

func searchTokenInList(token string, list []string, resultList *[]string) error {
	if len(list) == 0 {
		return errors.New("List is empty")
	}
	for _, item := range list {
		if strings.Contains(item, token) {
			*resultList = append(*resultList, item)
		}
	}
	slices.Sort(*resultList)

	return nil
}
