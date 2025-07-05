package terminal

import (
	ev "github.com/marinacompsci/go-dropdown/internal/event"
)

type Terminal struct {
}

func NewTerminal() *Terminal {
	return &Terminal{
	}
}

func (t *Terminal) Read(b byte) ev.Event {
	key := Key(b)

	switch(key) {
	case KEY_CTRL_C:
		return ev.UserPressedKeyCtrlC
	case KEY_DEL:
		return ev.UserPressedKeyDel
	case KEY_DOWN:
		return ev.UserPressedKeyDown
	case KEY_UP:
		return ev.UserPressedKeyUp
	case KEY_ESC:
		return ev.UserPressedKeyEsc
	default:
		return ev.UserPressedNormalChar
	}
}
