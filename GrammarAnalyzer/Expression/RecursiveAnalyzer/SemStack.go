package RecursiveAnalyzer

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/Scanner"
)

const (
	DEFAULTSIZE = 100
)

type SemStack struct {
	Stack []Scanner.Token
	Top   int
}

func NewSemStack() *SemStack {
	var s SemStack
	s.Stack = make([]Scanner.Token, DEFAULTSIZE)
	s.Top = 0
	return &s
}

func (s *SemStack) Push(e Scanner.Token) {
	if len(s.Stack) == s.Top {
		s.Stack = append(s.Stack, e)
	} else {
		s.Stack[s.Top] = e
	}
	s.Top++
}

func (s *SemStack) Pop() (Scanner.Token, bool) {
	if s.Top == 0 {
		return Scanner.Token{}, false
	} else {
		s.Top--
		return s.Stack[s.Top], true
	}
}

func (s *SemStack) IsEmpty() bool {
	return s.Top == 0
}

func (s *SemStack) Print() {
	if !s.IsEmpty() {
		for i := 0; i < s.Top; i++ {
			fmt.Printf("%v ", s.Stack[i])
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Stack is Empty")
	}
}
