/*
	Define here all things related to the rendering of the screen

and not to the prompt or the cursor position.
*/
package escsequence

const (
		clearScreenEscSeq = "\033[2J\033[1;1H" // \033[1:1H is move
)
func clearScreen() {
	//TODO: Think about implementing a pipeline for the execution
	// of a chain of moves
	//move.ToTopLeft

}


