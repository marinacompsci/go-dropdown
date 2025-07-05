/**
Write, update and keep user's input.
**/
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

func (p *Prompt) appendByte(newInput byte) {
	p.input = append(p.input, string(newInput))
}


func (p *Prompt) WriteFormatted() {
	fmt.Printf("\r%s%s", PromptSymbol, p.Stringified())
}

/*
Triggered when the user presses the DELETE key.
Deletes the last character on the prompt but obviously
not the PromptSymbol.
(Do not use p.length as its intended for outside use
and counts the PromptSymbol as one of the characters.)
*/
func (p *Prompt) trimInput() {
	if l := len(p.input); l > 0 {
		p.input = p.input[:l-1]
	}
}

/*
Return the user's input character count
(Do not use p.length as its intended for outside use
and counts the PromptSymbol as one of the characters.)
*/
func (p *Prompt) IsEmpty() bool {
	return len(p.input) == 0
}

//TODO: put prompt, menu, and screen in different packages to make methods like this private
/*
Return the prompt's length always with the PromptSymbol included.
Outsiders should not have to handle/workaround the fact that
when written the prompt includes the PromptSymbol which increases
its length visually for the terminal by one character although internally
the PromptSymbol is not saved as part of the user's input.
*/
func (p *Prompt) length() int {
	return len(p.input) + 1
}

func (p *Prompt) Stringified() string {
	return strings.Join(p.input, "")
}
