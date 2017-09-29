package Table

type IdentifierTable struct {
	T map[string]int
}

func NewIdentifierTable() *IdentifierTable {
	var it IdentifierTable
	it.T = make(map[string]int)
	return &it
}
