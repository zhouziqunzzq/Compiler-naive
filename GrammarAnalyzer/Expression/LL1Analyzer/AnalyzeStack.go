package LL1Analyzer

import "fmt"

const (
	DEFAULTSIZE = 100
)

type AnalyzeStack struct {
	Stack []string
	Top   int
}

func NewAnalyzeStack() *AnalyzeStack {
	var s AnalyzeStack
	s.Stack = make([]string, DEFAULTSIZE)
	s.Top = 0
	return &s
}

func (s *AnalyzeStack) Push(e string) {
	if len(s.Stack) == s.Top {
		s.Stack = append(s.Stack, e)
	} else {
		s.Stack[s.Top] = e
	}
	s.Top++
}

func (s *AnalyzeStack) Pop() (string, bool) {
	if s.Top == 0 {
		return "", false
	} else {
		s.Top--
		return s.Stack[s.Top], true
	}
}

func (s *AnalyzeStack) IsEmpty() bool {
	return s.Top == 0
}

func (s *AnalyzeStack) Print() {
	if !s.IsEmpty() {
		for i := 0; i < s.Top; i++ {
			fmt.Printf("%v ", s.Stack[i])
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Stack is Empty")
	}

}
