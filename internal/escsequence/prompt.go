/* Define here all escapes related to the prompt line.
Everything that can't the done atomically i.e.
needs to account for some state like the amount of 
characters in the prompt. 
*/
package escsequence

import (
)

type Prompt struct {
	input *string
}

(p *Prompt) func moveToEOL() {
}
