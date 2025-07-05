/*
*
Coordinates user input with data fetching, UI updates and
cursor movement.
Works as orchestrator, knows who to call thus does not need
to know everything - "I know a guy".
*
*/
package main

import (
	"fmt"
	"slices"
	"strings"

	c "github.com/marinacompsci/go-dropdown/internal/cursor"
	ev "github.com/marinacompsci/go-dropdown/internal/event"
	t "github.com/marinacompsci/go-dropdown/internal/terminal"
	"github.com/marinacompsci/go-dropdown/internal/exception"
	"github.com/marinacompsci/go-dropdown/internal/repository"
)

type Screen struct {
	repo *repository.ExampleRepository

	//TODO: maybe refactor into an UI package
	inSelectionMode bool
	cursor *c.Cursor
	menu *Menu
	prompt *Prompt
	
	terminal *t.Terminal
}


func NewScreen(r *repository.ExampleRepository) *Screen {
	return &Screen{
		repo: r,

		inSelectionMode: false,
		cursor: c.NewCursor(),
		menu: NewMenu(),
		prompt: NewPrompt(),

		terminal: t.NewTerminal(),
	}
}


func (s *Screen) PreRender() {
	fmt.Print(ScreenClear)
	fmt.Print(CursorToOrigin)
	fmt.Print(PromptSymbol)
}


func (s *Screen) ReadPrompt(b byte) error {
	event := s.terminal.Read(b)

	switch event {
	case ev.UserPressedKeyCtrlC:
		return exception.ErrUserInterrupted
	case ev.UserPressedKeyEsc:
		s.inSelectionMode = !s.inSelectionMode
		if s.inSelectionMode && s.menu.length() > 0 {
			s.cursor.MoveBelowPrompt()
		} else {
			s.cursor.MoveToPrompt(s.prompt.length())
		}
	case ev.UserPressedKeyUp:
		if s.inSelectionMode {
			newCursorY := s.cursor.Line - 1
			// Y = 1 is where the prompt line lives,
			// let's not go back there, this is what ESC is for.
			if newCursorY > 1 {
				s.cursor.MoveUp()
			}
		}
	case ev.UserPressedKeyDown:
		if s.inSelectionMode {
			paginationMax := 20
			newCursorY := s.cursor.Line + 1
			// The cursor's origin(furthest point north) is 1 which is where the prompt line lives,
			// so in order to actually go down N points from the origin
			// we have to stop at N+1 so we add 1 to paginationMax and to menu.length
			if newCursorY <= s.menu.length()+1 && newCursorY <= paginationMax+1 { 
				s.cursor.MoveDown()
			}
		}
	case ev.UserPressedKeyDel:
		s.prompt.trimInput()
		s.reset()

		if err := s.filterListByInput(s.prompt.Stringified()); err != nil {
			return err
		}
	case ev.UserPressedNormalChar:
		s.prompt.appendByte(b)
		s.reset()

		if err := s.filterListByInput(s.prompt.Stringified()); err != nil {
			return err
		}
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
		return exception.ErrInputNotInDB
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
		return exception.ErrInputNotInDB
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


func (s *Screen) reset() {
	fmt.Print(ScreenClear)
	fmt.Print(CursorToOrigin)
	s.prompt.WriteFormatted()
}
