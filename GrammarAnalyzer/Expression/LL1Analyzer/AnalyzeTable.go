package LL1Analyzer

import (
	"github.com/zhouziqunzzq/Compiler/Scanner"
)

type AnalyzeTable struct {
	// Use the string at the top of the stack and a TokenType to get a map[string]int
	// If specific string is not required, i.e. all string of this TokenType is ok,
	// the result will be stored in "" of this map
	T map[string]map[Scanner.TokenType]map[string]int
}

func NewAnalyzeTable() *AnalyzeTable {
	var at AnalyzeTable
	// Now just hard-code the Grammar here
	at.T = make(map[string]map[Scanner.TokenType]map[string]int)
	// line E
	at.T["E"] = make(map[Scanner.TokenType]map[string]int)
	at.T["E"][Scanner.INTEGERCONSTANT] = make(map[string]int)
	at.T["E"][Scanner.FLOATCONSTANT] = make(map[string]int)
	at.T["E"][Scanner.IDENTIFIER] = make(map[string]int)
	at.T["E"][Scanner.DELIMITER] = make(map[string]int)
	at.T["E"][Scanner.INTEGERCONSTANT][""] = 1
	at.T["E"][Scanner.FLOATCONSTANT][""] = 1
	at.T["E"][Scanner.IDENTIFIER][""] = 1
	at.T["E"][Scanner.DELIMITER]["("] = 1
	// line E1
	at.T["E1"] = make(map[Scanner.TokenType]map[string]int)
	at.T["E1"][Scanner.DELIMITER] = make(map[string]int)
	at.T["E1"][Scanner.END] = make(map[string]int)
	at.T["E1"][Scanner.DELIMITER]["+"] = 2
	at.T["E1"][Scanner.DELIMITER]["-"] = 2
	at.T["E1"][Scanner.DELIMITER][")"] = 3
	at.T["E1"][Scanner.END][""] = 3
	// line T
	at.T["T"] = make(map[Scanner.TokenType]map[string]int)
	at.T["T"][Scanner.INTEGERCONSTANT] = make(map[string]int)
	at.T["T"][Scanner.FLOATCONSTANT] = make(map[string]int)
	at.T["T"][Scanner.IDENTIFIER] = make(map[string]int)
	at.T["T"][Scanner.DELIMITER] = make(map[string]int)
	at.T["T"][Scanner.INTEGERCONSTANT][""] = 4
	at.T["T"][Scanner.FLOATCONSTANT][""] = 4
	at.T["T"][Scanner.IDENTIFIER][""] = 4
	at.T["T"][Scanner.DELIMITER]["("] = 4
	// line T1
	at.T["T1"] = make(map[Scanner.TokenType]map[string]int)
	at.T["T1"][Scanner.DELIMITER] = make(map[string]int)
	at.T["T1"][Scanner.END] = make(map[string]int)
	at.T["T1"][Scanner.DELIMITER]["+"] = 6
	at.T["T1"][Scanner.DELIMITER]["-"] = 6
	at.T["T1"][Scanner.DELIMITER]["*"] = 5
	at.T["T1"][Scanner.DELIMITER]["/"] = 5
	at.T["T1"][Scanner.DELIMITER][")"] = 6
	at.T["T1"][Scanner.END][""] = 6
	// line F
	at.T["F"] = make(map[Scanner.TokenType]map[string]int)
	at.T["F"][Scanner.INTEGERCONSTANT] = make(map[string]int)
	at.T["F"][Scanner.FLOATCONSTANT] = make(map[string]int)
	at.T["F"][Scanner.IDENTIFIER] = make(map[string]int)
	at.T["F"][Scanner.DELIMITER] = make(map[string]int)
	at.T["F"][Scanner.INTEGERCONSTANT][""] = 7
	at.T["F"][Scanner.FLOATCONSTANT][""] = 7
	at.T["F"][Scanner.IDENTIFIER][""] = 7
	at.T["F"][Scanner.DELIMITER]["("] = 8
	return &at
}
