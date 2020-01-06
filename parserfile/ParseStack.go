package parserfile

type parseStack struct {
	stk []pstackEntry
}

//Push : Pushes an object onto the front of the stack.
func (s *parseStack) push(itm pstackEntry) {
	s.stk = append([]pstackEntry{itm}, s.stk...)
}

//Empty : Reports whether the parseStack that ran it has 0 entries.
//If so, then returns true; otherwise, returns false.
func (s parseStack) empty() bool {
	if len(s.stk) == 0 {
		return true
	}
	return false
}

//Pop : Pops the first pstackEntry from the stack.
func (s *parseStack) pop() pstackEntry {
	top := s.top()
	s.stk = s.stk[1:]
	return top
}

func (s *parseStack) popNum(n int) {
	for i := 1; i <= n; i++ {
		s.pop()
	}
}

//Top : Returns the first pstackEntry in the stack.
func (s parseStack) top() pstackEntry {
	return s.stk[0]
}

func (s parseStack) String() string {
	var reversed parseStack
	for _, element := range s.stk {
		reversed.push(element)
	}
	var returnString string
	for _, element := range reversed.stk {
		returnString += element.String()
	}
	return returnString
}
