package tree

import "fmt"

//ParseTree : an interface for all tree structures in the parser.
type ParseTree interface {
	fmt.Stringer
	PrintTree()
	PrintTreeWork(int)
}
