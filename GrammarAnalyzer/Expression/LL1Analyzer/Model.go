package LL1Analyzer

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/Scanner"
)

const (
	GRAMMARCOUNT = 9
	VN           = 1
	VT           = 2
)

type LL1Analyzer struct {
	S                 *Scanner.Scanner
	Grammar           [][]string
	NotationTypeTable map[string]int
	at                *AnalyzeTable
	AS                *AnalyzeStack
}

func NewLL1Analyzer(s *string) *LL1Analyzer {
	var la LL1Analyzer
	la.S = Scanner.NewScanner(s)
	// Now just hard-code the NotationTypeTable here
	la.NotationTypeTable = make(map[string]int)
	la.NotationTypeTable["E"] = VN
	la.NotationTypeTable["E1"] = VN
	la.NotationTypeTable["T"] = VN
	la.NotationTypeTable["T1"] = VN
	la.NotationTypeTable["F"] = VN
	la.NotationTypeTable["I"] = VT
	la.NotationTypeTable["w0"] = VT
	la.NotationTypeTable["w1"] = VT
	la.NotationTypeTable["("] = VT
	la.NotationTypeTable[")"] = VT
	// Now just hard-code the Grammar here
	la.Grammar = make([][]string, GRAMMARCOUNT)
	la.Grammar[1] = []string{"T", "E1"}
	la.Grammar[2] = []string{"w0", "T", "E1"}
	la.Grammar[3] = []string{""}
	la.Grammar[4] = []string{"F", "T1"}
	la.Grammar[5] = []string{"w1", "F", "T1"}
	la.Grammar[6] = []string{""}
	la.Grammar[7] = []string{"I"}
	la.Grammar[8] = []string{"(", "E", ")"}

	la.at = NewAnalyzeTable()
	la.AS = NewAnalyzeStack()

	return &la
}

func isSameVT(s string, t *Scanner.Token) bool {
	switch s {
	case "w0":
		return t.Word == "+" || t.Word == "-"
	case "w1":
		return t.Word == "*" || t.Word == "/"
	case "I":
		return t.Type == Scanner.INTEGERCONSTANT || t.Type == Scanner.FLOATCONSTANT || t.Type == Scanner.IDENTIFIER
	case "(":
		return t.Word == "("
	case ")":
		return t.Word == ")"
	}
	return false
}

func pushReverse(as *AnalyzeStack, s []string) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != "" {
			as.Push(s[i])
		}
	}
}

func getAnalyzeTable(la *LL1Analyzer, m *map[string]int) (v int, ok bool) {
	if la.S.LastToken.Type == Scanner.INTEGERCONSTANT || la.S.LastToken.Type == Scanner.FLOATCONSTANT || la.S.LastToken.Type == Scanner.IDENTIFIER || la.S.LastToken.Type == Scanner.END {
		if v1, ok1 := (*m)[""]; ok1 {
			v = v1
			ok = ok1
		} else {
			ok = false
		}
	} else {
		if v1, ok1 := (*m)[la.S.LastToken.Word]; ok1 {
			v = v1
			ok = ok1
		} else {
			ok = false
		}
	}
	return v, ok
}

func (la *LL1Analyzer) Analyze() bool {
	la.AS.Push("E")
	la.AS.Print()
	for {
		la.S.Next()
		//fmt.Println(la.S.LastToken.Word)
		flag := false
		for !la.AS.IsEmpty() {
			x, ok := la.AS.Pop()
			if ok == false {
				return false
			}
			//fmt.Printf("%v POP.\n", x)
			if t, ok := la.NotationTypeTable[x]; ok {
				if t == VT {
					if isSameVT(x, la.S.LastToken) {
						flag = true
						break
					} else {
						return false
					}
				} else if t == VN {
					if m, ok := la.at.T[x][la.S.LastToken.Type]; ok {
						if v, ok1 := getAnalyzeTable(la, &m); ok1 {
							flag = true
							pushReverse(la.AS, la.Grammar[v])
							la.AS.Print()
						} else {
							return false
						}
					}
				} else {
					return false
				}
			} else {
				return false
			}
		}
		if flag {
			continue
		} else {
			if la.S.LastToken.Type == Scanner.END {
				return true
			} else {
				return false
			}
		}
	}
}
