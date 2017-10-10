package Table

// CharacterConstant Token format: CharConstant - (CHAR, ID)
type CharacterConstantTable struct {
	T map[rune]int
}

// StringConstant Token format: StrConstant - (STRING, ID)
type StringConstantTable struct {
	T map[string]int
}

// IntegerConstant Token format: Integer - (INT, ID)
type IntegerConstantTable struct {
	T map[int]int
}

// FloatConstant Token format: Float - (FLOAT, ID)
type FloatConstantTable struct {
	T map[float64]int
}

func NewConstantTable() (*CharacterConstantTable, *StringConstantTable, *IntegerConstantTable, *FloatConstantTable) {
	var (
		cc CharacterConstantTable
		sc StringConstantTable
		ic IntegerConstantTable
		fc FloatConstantTable
	)
	cc.T = make(map[rune]int)
	sc.T = make(map[string]int)
	ic.T = make(map[int]int)
	fc.T = make(map[float64]int)
	return &cc, &sc, &ic, &fc
	//return new(CharacterConstantTable), new(StringConstantTable), new(IntegerConstantTable), new(FloatConstantTable)
}
