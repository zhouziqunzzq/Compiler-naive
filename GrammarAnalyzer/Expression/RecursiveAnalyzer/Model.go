package RecursiveAnalyzer

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/GrammarAnalyzer"
	"github.com/zhouziqunzzq/Compiler/Scanner"
	"strconv"
)

const (
	QTMAXSIZE = 100
)

type RecursiveAnalyzer struct {
	S          *Scanner.Scanner
	sem        *SemStack
	QT         []GrammarAnalyzer.Quadruple
	qCounter   int //Counter for the QT
	tmpCounter int //Counter for the temp variable
}

func NewRecursiveAnalyzer(s *string) *RecursiveAnalyzer {
	var ra RecursiveAnalyzer
	ra.S = Scanner.NewScanner(s)
	ra.sem = NewSemStack()
	ra.QT = make([]GrammarAnalyzer.Quadruple, QTMAXSIZE)
	ra.qCounter = 0
	ra.tmpCounter = 0
	return &ra
}

func (ra *RecursiveAnalyzer) Analyze() bool {
	ra.S.Next()
	if ra.E() == false {
		return false
	}
	if ra.S.LastToken.Type == Scanner.END {
		fmt.Println("Generated Quadruple:")
		for i := 0; i < ra.qCounter; i++ {
			fmt.Printf("(%v, %v, %v, %v)\n", ra.QT[i].Op, ra.QT[i].Opr2.Word, ra.QT[i].Opr1.Word, ra.QT[i].Target)
		}
		return true
	} else {
		return false
	}
}

func (ra *RecursiveAnalyzer) E() bool {
	fmt.Println("E")
	if ra.T() == false {
		return false
	}
	for ra.S.LastToken.Type == Scanner.DELIMITER &&
		(ra.S.LastToken.Word == "+" || ra.S.LastToken.Word == "-") {
		fmt.Println(ra.S.LastToken.Word)
		op := ra.S.LastToken.Word
		ra.S.Next()
		if ra.T() == false {
			return false
		}
		if ra.GEQ(op) == false {
			return false
		}
	}
	return true
}

func (ra *RecursiveAnalyzer) T() bool {
	fmt.Println("T")
	if ra.F() == false {
		return false
	}
	for ra.S.LastToken.Type == Scanner.DELIMITER &&
		(ra.S.LastToken.Word == "*" || ra.S.LastToken.Word == "/") {
		fmt.Println(ra.S.LastToken.Word)
		op := ra.S.LastToken.Word
		ra.S.Next()
		if ra.F() == false {
			return false
		}
		if ra.GEQ(op) == false {
			return false
		}
	}
	return true
}

func (ra *RecursiveAnalyzer) F() bool {
	fmt.Println("F")
	if ra.S.LastToken.Type == Scanner.IDENTIFIER || ra.S.LastToken.Type == Scanner.INTEGERCONSTANT ||
		ra.S.LastToken.Type == Scanner.FLOATCONSTANT {
		fmt.Println(ra.S.LastToken.Word)
		ra.sem.Push(*ra.S.LastToken)
		ra.S.Next()
		return true
	} else if ra.S.LastToken.Type == Scanner.DELIMITER && ra.S.LastToken.Word == "(" {
		fmt.Println(ra.S.LastToken.Word)
		ra.S.Next()
		if ra.E() == false {
			return false
		}
		if ra.S.LastToken.Type == Scanner.DELIMITER && ra.S.LastToken.Word == ")" {
			fmt.Println(ra.S.LastToken.Word)
			ra.S.Next()
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (ra *RecursiveAnalyzer) GEQ(op string) bool {
	t1, ok1 := ra.sem.Pop()
	if !ok1 {
		return false
	}
	t2, ok2 := ra.sem.Pop()
	if !ok2 {
		return false
	}
	q := GrammarAnalyzer.NewQuadruple(op, t1, t2, "t"+strconv.Itoa(ra.tmpCounter))
	if ra.qCounter == len(ra.QT) {
		ra.QT = append(ra.QT, *q)
	} else {
		ra.QT[ra.qCounter] = *q
	}
	ra.qCounter++
	tt := Scanner.Token{Scanner.IDENTIFIER, "t" + strconv.Itoa(ra.tmpCounter), 0}
	ra.sem.Push(tt)
	ra.tmpCounter++
	return true
}
