package Table

// Keyword Token format: Keyword - (KEYWORD, ID)
type KeywordTable struct {
	T map[string]int
}

func NewKeywordTable() *KeywordTable {
	var t KeywordTable
	t.T = make(map[string]int)
	k := []string{"auto", "double", "int", "struct", "break", "else", "long",
		"switch", "case", "enum", "register", "typedef", "char", "extern", "return",
		"union", "const", "float", "short", "unsigned", "continue", "for", "signed",
		"void", "default", "goto", "sizeof", "volatile", "do", "if", "while", "static"}
	for i, kw := range k {
		t.T[kw] = i + 1
	}
	return &t
}
