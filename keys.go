package main

type Key int 

// Non-printable keys
const (
	KEY_CTRL_C Key		= 3
	KEY_ESC Key			= 27
	KEY_DEL Key 		= 127
)

// Navigation keys
const (
	KEY_DOWN Key 		= 106 // 'J' and not ARROW KEY DOWN
	KEY_UP Key 			= 107 // 'K' and not ARROW KEY UP
)
