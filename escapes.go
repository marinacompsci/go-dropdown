package main

/* SCREEN */
const (
	ScreenClear = "\033[2J"
)

/* CURSOR */
const (
	CursorToOrigin = "\033[1;1H" // X:0, Y:0
)

/* COLOR */
const (
	ColorForegroundBlack = "\033[30m"
	ColorBackgroundYellow = "\033[43m"
	ColorResetAll = "\033[0m"
)
