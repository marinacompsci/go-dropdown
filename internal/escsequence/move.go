/*
Define here all cursor indempotent moves that can be done
without greater consequence in regards to the prompt's state
and also does not change the UI's state.
*/
package escsequence

import (
	"fmt"
)

type Move struct {
	
}

(m *Move) func ToTopLeft() {
	fmt.Printf("\033[1;1H")
}
