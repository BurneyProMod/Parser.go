package parserfile

import (
	"./tree"
	"fmt"
	"strconv"
)

//Parser : A bottom-up parser. The main class of parserfile.
type Parser struct {
	Filename string
}

type actionChoice int

var (
	outputTable    [][]string
	inputArray     []string
	grammar        [][]string
	aTable         [][]string
	gTable         [][]string
	choice         actionChoice
	actionValue    string
	newPush        pstackEntry
	aTableIndex    map[string]int
	gTableIndex    map[string]int
	notGrammatical bool
	iQ             inputQueue
	pStack         parseStack
	tStack         tree.TreeStack
)

const (
	accept actionChoice = iota
	ungrammatical
	shift
	reduce
)

func (p Parser) setupParser() {
	setupInputArray()
	choice = ungrammatical
	setupTables()
	notGrammatical = false
	fmt.Println("File to parse: ", p.Filename)
	outputTable = make([][]string, 20)
	for i := range outputTable {
		outputTable[i] = make([]string, 11)
	}
	setupInputQueue()
	setupTableIndices()
	newPush = pstackEntry{stateSym: "0", grammarSym: ""}
	setupPStack()
}

func setupInputArray() {
	//    [] inputArray = { "id"};                                  // simplest possible grammatical input
	//    [] inputArray = { "id", "+", "id"};                       // simple grammatical input
	inputArray = []string{"id", "+", "id", "*", "id"} //grammatical input from book
	//    [] inputArray = { "id", "+", "id", "+", "id"};            // left assoc
	//    [] inputArray = { "(", "id", ")"};                        // parens1
	//    [] inputArray = { "id", "+", "(", "id", "+", "id", ")"};  // parens2
	//    [] inputArray = { "id", "+", "id", "*"};                  // ungrammatical input
}

func setupTables() {
	aTable = [][]string{ // action table
		{"S5", "", "", "S4", "", ""},     // 0
		{"", "S6", "", "", "", "accept"}, // 1
		{"", "R2", "S7", "", "R2", "R2"}, // 2
		{"", "R4", "R4", "", "R4", "R4"}, // 3
		{"S5", "", "", "S4", "", ""},     // 4
		{"", "R6", "R6", "", "R6", "R6"}, // 5
		{"S5", "", "", "S4", "", ""},     // 6
		{"S5", "", "", "S4", "", ""},     // 7
		{"", "S6", "", "", "S11", ""},    // 8
		{"", "R1", "S7", "", "R1", "R1"}, // 9
		{"", "R3", "R3", "", "R3", "R3"}, // 10
		{"", "R5", "R5", "", "R5", "R5"}, // 11
	}
	gTable = [][]string{ // goto table
		{"1", "2", "3"}, // 0
		{"", "", ""},    // 1
		{"", "", ""},    // 2
		{"", "", ""},    // 3
		{"8", "2", "3"}, // 4
		{"", "", ""},    // 5
		{"", "9", "3"},  // 6
		{"", "", "10"},  // 7
		{"", "", ""},    // 8
		{"", "", ""},    // 9
		{"", "", ""},    // 10
		{"", "", ""},    // 11
	}
	grammar = [][]string{
		{"E", "->", "E", "+", "T"}, // 1
		{"E", "->", "T"},           // 2
		{"T", "->", "T", "*", "F"}, // 3
		{"T", "->", "F"},           // 4
		{"F", "->", "(", "E", ")"}, // 5
		{"F", "->", "id"},          // 6
	}
}

func setupPStack() {
	pStack.stk = make([]pstackEntry, 0)
	tStack.Stk = make([]tree.ParseTree, 0)
}

func setupInputQueue() {
	iQ.newInputQueue(inputArray)
}

func setupTableIndices() {
	aTableIndex = make(map[string]int)
	gTableIndex = make(map[string]int)
	aTableIndex["id"] = 0
	aTableIndex["+"] = 1
	aTableIndex["*"] = 2
	aTableIndex["("] = 3
	aTableIndex[")"] = 4
	aTableIndex["$"] = 5

	gTableIndex["E"] = 0
	gTableIndex["T"] = 1
	gTableIndex["F"] = 2
}

func printParseTree() {
	fmt.Println("Parse tree:")
	topOfTStack := tStack.Top()
	fmt.Println(topOfTStack.String())
	topOfTStack.PrintTree()
	//	fmt.Println("Hello World from " + getClass().getName());
}

//Parse : Parses stuff.
func (p Parser) Parse() {
	p.setupParser()
	fmt.Println("                input          action    action  value   length  temp            goto      goto   stack       ")
	fmt.Println("Stack           tokens         lookup    value   of LHS  of RHS  stack           lookup    value  action      parse tree stack")
	fmt.Println("______________________________________________________________________________________________________________________________")
	for {
		parse1Step()
		if !(choice == shift || choice == reduce) {
			break
		}
	}
	if !(notGrammatical) {
		fmt.Println("")
		printParseTree()
	}
}

func parse1Step() {
	valueofLHS := ""
	tempStack := ""
	gotoValueIndex1st := ""
	valueofLHSIndex := -1
	lengthofRHS := 0
	lengthofRHSstr := ""
	gotoValue := ""
	initQue := iQ.String()
	inputQueFront := iQ.first()
	pStack.push(newPush) // value is created on previous cycle
	initPStack := pStack.String()
	newPushStr := ""
	gotoLookupStr := ""
	// work starts here
	evaluateActionChoice() // sets the value of "choice" and "actionValue"
	actionIndex1st := pStack.top().stateSym
	switch choice {
	case accept:
		break
	case ungrammatical:
		notGrammatical = true
		break
	case shift:
		substring := actionValue[1:len(actionValue)] //get substring from index 1 of actionValue
		newPush = pstackEntry{stateSym: substring, grammarSym: iQ.first()}
		if iQ.first() == "id" {
			id := tree.TermSym{TS: "id"}
			tStack.Push(id)
		}
		iQ.pop()
		break
	case reduce:
		substring := actionValue[1:len(actionValue)] //get substring from index 1 of actionValue
		x, _ := strconv.Atoi(substring)
		valueofLHSIndex = x                               // 2nd index + 1 for action lookup
		valueofLHS = grammar[valueofLHSIndex-1][0]        // lookup LHS from grammar table (0-based)
		lengthofRHS = len(grammar[valueofLHSIndex-1]) - 2 // compute length of RHS
		var popped parseStack
		popped.stk = make([]pstackEntry, lengthofRHS)
		for i := 1; i <= lengthofRHS; i++ {
			p := pStack.pop()
			popped.push(p)
		}
		//                pStack.popNum(lengthofRHS);
		tempStack = pStack.String()
		gotoValueIndex1st = pStack.top().stateSym
		y, _ := strconv.Atoi(gotoValueIndex1st)
		gotoValue = gTable[y][gTableIndex[valueofLHS]]
		newPush = pstackEntry{stateSym: gotoValue, grammarSym: valueofLHS}
		// start building tree
		if lengthofRHS == 1 {
			oldPt := []tree.ParseTree{tStack.Pop()}
			child := tree.ListOfTrees{Lot: oldPt}
			lhsSym := tree.LHSSym{Sym: valueofLHS}
			tr := tree.NonLeafTree{Parent: lhsSym, Children: child}
			tStack.Push(tr)
		} else if lengthofRHS == 3 && popped.top().grammarSym == "(" {
			var aarg2 tree.ParseTree
			aarg2 = tStack.Pop()
			aarg1 := []tree.ParseTree{tree.TermSym{TS: "("}}
			aarg3 := tree.TermSym{TS: ")"}
			llot := tree.ListOfTrees{Lot: aarg1}
			llot.AddLast(aarg2)
			llot.AddLast(aarg3)
			pparent := tree.LHSSym{Sym: valueofLHS}
			ttr := tree.NonLeafTree{Parent: pparent, Children: llot}
			tStack.Push(ttr)
		} else if lengthofRHS == 3 {
			arg3 := tStack.Pop()
			arg1 := []tree.ParseTree{tStack.Pop()}
			popped.pop()
			operator := popped.top().grammarSym
			arg2 := tree.TermSym{TS: operator}
			lot := tree.ListOfTrees{Lot: arg1}
			lot.AddLast(arg2)
			lot.AddLast(arg3)
			parent := tree.LHSSym{Sym: valueofLHS}
			tr := tree.NonLeafTree{Parent: parent, Children: lot}
			tStack.Push(tr)
		}
		break
	}
	if choice == shift || choice == reduce {
		newPushStr = "push " + newPush.String()
	}
	if lengthofRHS > 0 {
		lengthofRHSstr = strconv.Itoa(lengthofRHS)
	}
	if choice == reduce {
		gotoLookupStr = "[" + gotoValueIndex1st + "," + valueofLHS + "]"
	}
	if notGrammatical {
		tStack.Stk = make([]tree.ParseTree, 1)
	}
	fmt.Printf("%-14s  %-14s [%2s,%2s]   %-6s  %-6s  %-7s %-14s  %-6s    %-5s  %-9s  %-20s\n", initPStack, initQue, actionIndex1st, inputQueFront, actionValue, valueofLHS, lengthofRHSstr, tempStack, gotoLookupStr, gotoValue, newPushStr, tStack.String())
}

func evaluateActionChoice() {
	x, _ := strconv.Atoi(pStack.top().stateSym)
	// get action value
	actionValue = aTable[x][aTableIndex[iQ.first()]]
	// parse action choice
	if actionValue == "accept" {
		choice = accept
	} else if actionValue == "" {
		choice = ungrammatical
		actionValue = "ungrammatical"
	} else if actionValue[0] == 'S' {
		choice = shift
	} else if actionValue[0] == 'R' {
		choice = reduce
	}
}
