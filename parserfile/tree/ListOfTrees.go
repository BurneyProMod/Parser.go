package tree

import "fmt"

//ListOfTrees : A slice that contains ParseTrees that is used to store a list of the trees.
type ListOfTrees struct {
	Lot []ParseTree
}

//NewListOfTrees : Initializer function for ListOfTrees
func (l *ListOfTrees) NewListOfTrees(tree ParseTree) {
	l.AddLast(tree)
}

//AddLast : Adds the given ParseTree to the end of slice
func (l *ListOfTrees) AddLast(tree ParseTree) {
	l.Lot = append(l.Lot, tree)
}

//PrintTreeWork : Prints all the elements in this ListOfTrees
func (l ListOfTrees) PrintTreeWork(indentLevel int) {
	for _, element := range l.Lot {
		element.PrintTreeWork(indentLevel)
	}
}

func (l ListOfTrees) String() string {
	returnString := ""
	for _, element := range l.Lot {
		returnString += fmt.Sprint(element) + " "
	}
	return returnString
}
