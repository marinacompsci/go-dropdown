package main

import (
	"fmt"
	"slices"
	"strings"
)


type Screen struct {
	data []string
	prompt *Prompt
	menu *Menu
}


func NewScreen(p *Prompt, m *Menu, data []string) *Screen {
	return &Screen{
		data: data,
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
	inputNotFound := true

	if len(s.data) == 0 {
		return ErrEmptyListAsResult
	}

	var tmpList []string
	for _, item := range s.data {
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
