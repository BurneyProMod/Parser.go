package parserfile

type inputQueue struct {
	inQ []string
}

//newInputQueue : sets up the inputQueue's slice of strings, given an input array. Adds a nice $ to the
//end of it too!
func (iQ *inputQueue) newInputQueue(inputArray []string) {
	for _, element := range inputArray {
		iQ.inQ = append(iQ.inQ, element)
	}
	iQ.inQ = append(iQ.inQ, "$")
}

//Pop : Pops the first string from the stack.
func (iQ *inputQueue) pop() string {
	var x string
	x, iQ.inQ = iQ.inQ[0], iQ.inQ[1:]
	return x
}

//first : Returns the first string in the stack.
func (iQ inputQueue) first() string {
	return iQ.inQ[0]
}

func (iQ inputQueue) String() string {
	returnString := ""
	for _, s := range iQ.inQ {
		returnString += s
	}
	return returnString
}
