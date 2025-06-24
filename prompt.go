package main

import (
	"fmt"
	"strings"
)


type Prompt struct {
	input []string
}

func NewPrompt() *Prompt {
	var input []string

	return &Prompt{
		input: input,
	}
}

const PromptSymbol = ">"

func (p *Prompt) Read(b byte) error {
	if b == KeyCtrlC {
		return ErrUserInterrupted
	} else if b == KeyDelete {
		p.trimByte()
	} else {
		p.appendByte(b)
	}
	return nil
}

func (p *Prompt) appendByte(newInput byte) {
	p.input = append(p.input, string(newInput))
}


func (p *Prompt) WriteFormatted() {
	fmt.Printf("\r%s%s", PromptSymbol, p.Stringified())
}


func (p *Prompt) trimByte() {
	if l := p.length(); l > 0 {
		p.input = (p.input)[:l-1]
	}
}

func (p *Prompt) IsEmpty() bool {
	return p.length() == 0
}

func (p *Prompt) length() int {
	return len(p.input)
}

func (p *Prompt) Stringified() string {
	return strings.Join(p.input, "")
}
