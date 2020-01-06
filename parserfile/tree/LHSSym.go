package tree

import "fmt"

//LHSSym : The Symbol on the left-hand side.
type LHSSym struct {
	Sym string
}

//PrintTreeWork : prints out the LHSSym given an indent level.
func (l LHSSym) PrintTreeWork(indentLevel int) {
	outString := ""
	for i := 0; i < indentLevel; i++ {
		outString += "    "
	}
	fmt.Println(outString, l)
}

//PrintTree : prints the LHSSym with indent level 0.
func (l LHSSym) PrintTree() {
	l.PrintTreeWork(0)
}

func (l LHSSym) String() string {
	return l.Sym
}
