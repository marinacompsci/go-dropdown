package exception

import (
	"errors"
)



var (
	ErrUserInterrupted = errors.New("User pressed CTRL-C")
)


var (
	ErrEmptyListAsArg = errors.New("Empty list passed as argument")
	//TODO: rename this error to be more mnemonic 
	ErrInputNotInDB = errors.New("Input not found in DB")
)
