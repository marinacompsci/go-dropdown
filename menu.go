/**
Write to, update and hold list of results.
**/
package main

import (
	"strings"
	"github.com/marinacompsci/go-dropdown/internal/exception"

)

type Menu struct {
	//TODO: do not export list
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

func (m *Menu) length() int {
	return len(m.list)
}


func formatWithNewLines(l []string) (error, string) {
	if len(l) == 0 { 
		return exception.ErrEmptyListAsArg, ""
	}
	return nil, strings.Join(l, "\n\r")
}

