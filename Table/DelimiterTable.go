package Table

type DelimiterTable struct {
	T map[string]int
}

func NewDelimiterTable() *DelimiterTable {
	var t DelimiterTable
	t.T = make(map[string]int)
	k := []string{"+", "-", "*", "/", ">", "<", "=", ">=", "<=", "==", "!", "!=",
		"<<", ">>", ",", ".", "(", ")", "[", "]", "{", "}", ";"}
	for i, kw := range k {
		t.T[kw] = i + 1
	}
	return &t
}
