package Scanner

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/Scanner/Table"
)

type Scanner struct {
	sct       *StateChangeTable
	State     int
	lastState int
	CurIndex  int
	content   []rune
	Buffer    []rune
	LastToken *Token
	kt        *Table.KeywordTable
	dt        *Table.DelimiterTable
	charct    *Table.CharacterConstantTable
	strct     *Table.StringConstantTable
	intct     *Table.IntegerConstantTable
	floatct   *Table.FloatConstantTable
	it        *Table.IdentifierTable
	hft       *HandleFuncTable
}

func NewScanner(s *string) *Scanner {
	var sc Scanner
	sc.content = []rune(*s)
	sc.Buffer = make([]rune, 0, 256)
	sc.State = 1
	sc.lastState = 1
	sc.LastToken = new(Token)

	sc.CurIndex = 0
	sc.sct = NewStateChangeTable()
	sc.kt = Table.NewKeywordTable()
	sc.it = Table.NewIdentifierTable()
	sc.dt = Table.NewDelimiterTable()
	sc.charct, sc.strct, sc.intct, sc.floatct = Table.NewConstantTable()
	sc.hft = NewHandleFuncTable()
	return &sc
}

func (sc *Scanner) Rewind() {
	sc.CurIndex--
	sc.Buffer = sc.Buffer[:len(sc.Buffer)-1]
}

func (sc *Scanner) Reset() {
	sc.Buffer = sc.Buffer[:0]
	sc.State = 1
}

/*func (sc *Scanner) GetContent() []rune {
	return sc.content
}*/

func (sc *Scanner) Next() {
	for sc.CurIndex != len(sc.content) && !(sc.State == 1 && sc.lastState != 1) {
		//fmt.Printf("CurState: %v, LastState: %v\n", sc.State, sc.lastState)
		// Update lastState
		sc.lastState = sc.State
		// Get a char and move forward
		c := sc.content[sc.CurIndex]
		sc.Buffer = append(sc.Buffer, c)
		sc.CurIndex++
		//fmt.Printf("CurIndex: %v, CurChar: %v\n", sc.CurIndex, string(c))
		// Jump to next State or error
		if nextState, ok := sc.sct.t[sc.State][c]; ok {
			sc.State = nextState
		} else {
			//fmt.Println("NO NEXTSTATE FOUND!")
			sc.LastToken.Type = ERROR
			return
		}
		// Call handle func if needed
		if f, ok := sc.hft.t[sc.State]; ok {
			f(sc)
		}
	}
	sc.lastState = 1
	if sc.CurIndex == len(sc.content) {
		sc.LastToken.Type = END
	}
}

func (sc *Scanner) PrintTables() {
	fmt.Println("##### CharConstant Table #####")
	fmt.Println("ID\tCHAR")
	for k, v := range sc.charct.T {
		fmt.Printf("%v\t%v\n", v, string(k))
	}
	fmt.Println("##### StringConstant Table #####")
	fmt.Println("ID\tSTRING")
	for k, v := range sc.strct.T {
		fmt.Printf("%v\t%v\n", v, k)
	}
	fmt.Println("##### IntegerConstant Table #####")
	fmt.Println("ID\tINT")
	for k, v := range sc.intct.T {
		fmt.Printf("%v\t%v\n", v, k)
	}
	fmt.Println("##### FloatConstant Table #####")
	fmt.Println("ID\tFLOAT")
	for k, v := range sc.floatct.T {
		fmt.Printf("%v\t%v\n", v, k)
	}
	fmt.Println("##### Identifier Table #####")
	fmt.Println("ID\tIDENTIFIER")
	for k, v := range sc.it.T {
		fmt.Printf("%v\t%v\n", v, k)
	}
}
