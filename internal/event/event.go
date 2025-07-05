package event


type Event int

const (
	UserPressedKeyCtrlC Event 	= iota
	UserPressedKeyDel
	UserPressedKeyUp
	UserPressedKeyDown
	UserPressedKeyEsc
	UserPressedNormalChar
)
