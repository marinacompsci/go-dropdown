package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/marinacompsci/go-dropdown/internal/repository"
	c "github.com/marinacompsci/go-dropdown/internal/cursor"
)

type Screen struct {
	repo *repository.ExampleRepository

	cursor *c.Cursor
	prompt *Prompt
	menu *Menu
	inSelectionMode bool
}


func NewScreen(p *Prompt, m *Menu, r *repository.ExampleRepository) *Screen {
	return &Screen{
		repo: r,
		prompt: p,
		menu: m,
		//TODO: Add keyboard struct to read byte
		inSelectionMode: false,
		cursor: &c.Cursor{
			Col: 1,
			Line: 1,
		},
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
				s.cursor.MoveBelowPrompt()
			} else {
				s.cursor.MoveToPrompt(s.prompt.length())
			}
			return nil
		case ErrKeyUp:
			if s.inSelectionMode {
				newCursorY := s.cursor.Line - 1
				// Y = 1 is where the prompt line lives,
				// let's not go back there, this is what ESC is for.
				if newCursorY <= 1 {
					return nil
				}
				s.cursor.MoveUp()

			}

			return nil
		case ErrKeyDown:
			if s.inSelectionMode {
				paginationMax := 20
				newCursorY := s.cursor.Line + 1
				// The cursor's origin(furthest point north) is 1 which is where the prompt line lives,
				// so in order to actually go down N points from the origin
				// we have to stop at N+1 so we add 1 to paginationMax and to menu.length
				if newCursorY > s.menu.length()+1 || newCursorY > paginationMax+1 { 
					return nil
				}
				s.cursor.MoveDown()
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

		s.cursor.MoveToPrompt(s.prompt.length())

	}

	return nil
}


func reset() {
	fmt.Print(ScreenClear)
	fmt.Print(CursorToOrigin)
}
