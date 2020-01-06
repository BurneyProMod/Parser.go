package tree

import "fmt"

//TermSym : RepresenTS a symbol in the parser. Implements ParseTree.
type TermSym struct {
	TS string
}

func (me TermSym) String() string {
	return me.TS
}

//PrintTree : PrinTS TermSym with indentation level 0
func (me TermSym) PrintTree() {
	me.PrintTreeWork(0)
}

//PrintTreeWork : prinTS out the TermSym given an indent level.
func (me TermSym) PrintTreeWork(indentLevel int) {
	ouTString := ""
	for i := 0; i < indentLevel; i++ {
		ouTString += "    "
	}
	fmt.Println(ouTString, me)
}
