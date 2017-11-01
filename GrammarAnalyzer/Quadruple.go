package GrammarAnalyzer

import (
	"github.com/zhouziqunzzq/Compiler/Scanner"
)

type Quadruple struct {
	Op     string
	Opr1   Scanner.Token
	Opr2   Scanner.Token
	Target string
}

func NewQuadruple(top string, topr1 Scanner.Token, topr2 Scanner.Token, ttarget string) *Quadruple {
	q := Quadruple{top, topr1, topr2, ttarget}
	return &q
}
