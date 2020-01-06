package tree

//NonLeafTree : Represents a part of the tree with no leaves. Implements ParseTree.
type NonLeafTree struct {
	Parent   LHSSym
	Children ListOfTrees
}

//PrintTreeWork : Prints the Parent and Children of this NonLeafTree, given an indent level
func (n NonLeafTree) PrintTreeWork(indentLevel int) {
	n.Parent.PrintTreeWork(indentLevel)
	n.Children.PrintTreeWork(indentLevel + 1)
}

func (n NonLeafTree) String() string {
	return "[" + n.Parent.String() + " " + n.Children.String() + "]"
}

//PrintTree : Prints this NonLeafTree with indent level 0
func (n NonLeafTree) PrintTree() {
	n.PrintTreeWork(0)
}
