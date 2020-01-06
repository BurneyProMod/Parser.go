package parserfile

import "fmt"

type pstackEntry struct {
	stateSym, grammarSym string
}

func (se pstackEntry) String() string {
	return fmt.Sprintf("%s%s", se.grammarSym, se.stateSym)
}
