package main

import (
	"strings"

)

type Menu struct {
	list []string
}

func NewMenu() *Menu {
	var list []string
	return &Menu{
		list: list,
	}
}

func (m *Menu) Update(newList []string) {
	m.list = newList
}

func (m *Menu) FormattedList() (error, string) {
	err, list := formatWithNewLines(m.list)
	if err != nil {
		return err, ""
	}
	s := "\n\r" + list
	return nil, s
}


/*********************** HELPERS ***********************/
func formatWithNewLines(l []string) (error, string) {
	if len(l) == 0 { 
		return ErrEmptyListAsArg, ""
	}
	return nil, strings.Join(l, "\n\r")
}

