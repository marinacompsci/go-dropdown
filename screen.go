package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/marinacompsci/go-dropdown/internal/repository"
)


type Screen struct {
	repo *repository.ExampleRepository
	prompt *Prompt
	menu *Menu
}


func NewScreen(p *Prompt, m *Menu, r *repository.ExampleRepository) *Screen {
	return &Screen{
		repo: r,
		prompt: p,
		menu: m,
	}
}


func (s *Screen) PreRender() {
	reset()
	fmt.Print(PromptSymbol)
}


func (s *Screen) ReadPrompt(b byte) error {
	if err := s.prompt.Read(b); err != nil {
		return err
	}

	reset()

	s.prompt.WriteFormatted()


	if err := s.filterListByInput(s.prompt.Stringified()); err != nil {
		return err
	}

	return nil
}

func (s *Screen) filterListByInput(input string) error {
	data, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	inputNotFound := true

	if len(data) == 0 {
		return ErrEmptyListAsResult
	}

	var tmpList []string
	for _, item := range data {
		if strings.Contains(item, input) {
			inputNotFound = false
			item = strings.Replace(item, input, fmt.Sprintf("%s%s%s%s", ColorForegroundBlack, ColorBackgroundYellow, input, ColorResetAll), -1)
			tmpList = append(tmpList, item)
		}
	}

	if inputNotFound {
		return ErrEmptyListAsResult
	}

	slices.Sort(tmpList)


	if !s.prompt.IsEmpty() {
		s.menu.Update(tmpList)
		err, l := s.menu.FormattedList()
		if err != nil {
			return err
		}
		fmt.Print(l)

		moveCursorToPosition(1, s.prompt.length()+2)
	}

	return nil
}

/*********************** HELPERS *******************/
func moveCursorToPosition(line int, column int) {
	fmt.Printf("\033[%d;%dH", line, column)
}


func reset() {
	fmt.Print(ScreenClear)
	fmt.Print(CursorToOrigin)
}
