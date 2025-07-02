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
	inSelectionMode bool
	cursorX int
	cursorY int
}


func NewScreen(p *Prompt, m *Menu, r *repository.ExampleRepository) *Screen {
	return &Screen{
		repo: r,
		prompt: p,
		menu: m,
		//TODO: Add keyboard struct to read byte
		inSelectionMode: false,
		//TODO: Add cursor struct with state and methods
		cursorX: 1,
		cursorY: 1,
	}
}


func (s *Screen) PreRender() {
	reset()
	fmt.Print(PromptSymbol)
}


func (s *Screen) ReadPrompt(b byte) error {
	if err := s.prompt.Read(b); err != nil {
		switch err {
		case ErrKeyEsc:
			s.inSelectionMode = !s.inSelectionMode
			if s.inSelectionMode && s.menu.length() > 0 {
				moveCursorToPosition(2, 1)
				s.cursorY = 2
			} else {
				//TODO: put the promptSymbol inside the prompt so that
				// adding 1 moves us to the right side of the prompt ">[]" when the prompt is empty
				moveCursorToPosition(1, s.prompt.length()+2)
				s.cursorY = 1
			}
			return nil
		case ErrKeyUp:
			if s.inSelectionMode {
				newCursorY := s.cursorY - 1
				// Y = 1 is where the prompt line lives,
				// let's not go back there, this is what ESC is for.
				//TODO: change to < 1 when the prompt includes the promptSymbol
				if newCursorY < 2 {
					return nil
				}
				//TODO: update cursor inside moveCursor, make it a method
				//TODO: change order of args to (column, line) to reflect X,Y
				moveCursorToPosition(newCursorY, 1)
				s.cursorY -= 1
			}

			return nil
		case ErrKeyDown:
			if s.inSelectionMode {
				paginationMax := 20
				newCursorY := s.cursorY + 1
				// The cursor's origin(furthest point north) is 1 which is where the prompt line lives,
				// so in order to actually go down N points from the origin
				// we have to stop at N+1 so we add 1 to paginationMax and to menu.length
				if newCursorY > s.menu.length()+1 || newCursorY > paginationMax+1 { 
					return nil
				}
				moveCursorToPosition(newCursorY, 1)
				s.cursorY += 1
			}
			return nil
		default:
			return err
		}
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

func moveCursorToPosition(line int, column int) {
	fmt.Printf("\033[%d;%dH", line, column)
}


func reset() {
	fmt.Print(ScreenClear)
	fmt.Print(CursorToOrigin)
}
