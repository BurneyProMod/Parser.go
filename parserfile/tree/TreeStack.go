package tree

import "fmt"

//TreeStack : A slice of ParseTrees that functions as the current stack of trees.
type TreeStack struct {
	Stk []ParseTree
}

//Push : Pushes an object onto the stack
func (s *TreeStack) Push(itm ParseTree) {
	s.Stk = append([]ParseTree{itm}, s.Stk...)
}

//Empty : Reports whether the TreeStack that ran it has 0 entries.
//If so, then returns true; otherwise, returns false.
func (s TreeStack) Empty() bool {
	if len(s.Stk) == 0 {
		return true
	}
	return false
}

//Pop : Pops the first ParseTree from the stack.
func (s *TreeStack) Pop() ParseTree {
	top := s.Top()
	s.Stk = s.Stk[1:]
	return top
}

//PopNum : Pops everything (up to the given index) from the stack.
func (s *TreeStack) PopNum(n int) {
	for i := 1; i <= n; i++ {
		s.Pop()
	}
}

//Top : Returns the first ParseTree in the stack.
func (s TreeStack) Top() ParseTree {
	return s.Stk[0]
}

func (s TreeStack) String() string {
	returnString := ""
	for _, element := range s.Stk {
		returnString += " " + fmt.Sprint(element)
	}
	return returnString
}
