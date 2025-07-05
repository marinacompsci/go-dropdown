/**
Move cursor vertically and horizontally across the screen.
**/
package cursor

import (
	"fmt"
)

type Cursor struct {
	Col int
	Line int
}


func NewCursor() *Cursor {
	return &Cursor{
		Col: 1,
		Line: 1,
	}
}


func (c *Cursor) MoveDown() {
	c.Line += 1
	c.offset(c.Line, c.Col)
}

func (c *Cursor) MoveUp() {
	if c.Line - 1 <= 1 {
		return
	}
	c.Line -= 1
	c.offset(c.Line, c.Col)
}

func (c *Cursor) MoveToPrompt(promptLen int) {
	c.offset(1, promptLen + 1)
}

func (c *Cursor) MoveBelowPrompt() {
	c.offset(2, 1)
}

func (c *Cursor) offset(line int, column int) {
	c.Line = line
	c.Col = column
	fmt.Printf("\033[%d;%dH", c.Line, c.Col)
}

