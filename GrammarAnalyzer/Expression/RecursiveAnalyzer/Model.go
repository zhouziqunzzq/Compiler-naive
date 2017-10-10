package RecursiveAnalyzer

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/Scanner"
)

type RecursiveAnalyzer struct {
	S *Scanner.Scanner
}

func NewRecursiveAnalyzer(s *string) *RecursiveAnalyzer {
	var ra RecursiveAnalyzer
	ra.S = Scanner.NewScanner(s)
	return &ra
}

func (ra *RecursiveAnalyzer) Analyze() bool {
	ra.S.Next()
	if ra.E() == false {
		return false
	}
	if ra.S.LastToken.Type == Scanner.END {
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
		ra.S.Next()
		if ra.T() == false {
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
		ra.S.Next()
		if ra.F() == false {
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
