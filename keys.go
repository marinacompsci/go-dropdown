package main

type Key int 

// Special keys
const (
	KEY_CTRL_C Key		= 3
	KEY_ESC Key			= 27
	KEY_DEL Key 		= 127
)

const (
	//TODO: first I must enter into selection mode with ESC
	KEY_DOWN Key 		= 106 // 'J' and not ARROW KEY DOWN
	KEY_UP Key 			= 107 // 'K' and not ARROW KEY UP
)
