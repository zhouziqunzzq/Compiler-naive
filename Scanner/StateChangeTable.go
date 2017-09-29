package Scanner

const (
	NodeNumber = 20
)

type StateChangeTable struct {
	t []map[rune]int
}

func fillAlphabet(m *map[rune]int, t int) {
	for i := int('a'); i <= int('z'); i++ {
		(*m)[rune(i)] = t
	}
	for i := int('A'); i <= int('Z'); i++ {
		(*m)[rune(i)] = t
	}
}

func fillNumber(m *map[rune]int, t int) {
	for i := int('0'); i <= int('9'); i++ {
		(*m)[rune(i)] = t
	}
}

func fillSingleDelimiter(m *map[rune]int, t int) {
	//d := []rune{'+', '-', '*', '/', '>', '<', '=', ',', '.', '(', ')', '[', ']', '{', '}', ';', '!'}
	d := []rune{'+', '-', '*', '/', ',', '.', '(', ')', '[', ']', '{', '}', ';'}
	for _, r := range d {
		(*m)[r] = t
	}
}

func fillDoubleDelimiter(m *map[rune]int, t int) {
	//d := []rune{'+', '-', '*', '/', '>', '<', '=', ',', '.', '(', ')', '[', ']', '{', '}', ';', '!'}
	d := []rune{'>', '<', '=', '!', '&', '|'}
	for _, r := range d {
		(*m)[r] = t
	}
}

func fillAll(m *map[rune]int, t int) {
	for i := 0; i < 128; i++ {
		(*m)[rune(i)] = t
	}
}

func (sct *StateChangeTable) init() {
	sct.t = make([]map[rune]int, NodeNumber)
	for i := 0; i < NodeNumber; i++ {
		sct.t[i] = make(map[rune]int)
	}
	// Node 1
	sct.t[1]['\n'], sct.t[1]['\r'], sct.t[1][' '], sct.t[1]['\t'] = 1, 1, 1, 1
	sct.t[1]['\''], sct.t[1]['"'] = 5, 8
	fillAlphabet(&sct.t[1], 2)
	sct.t[1]['_'] = 2
	fillNumber(&sct.t[1], 3)
	fillSingleDelimiter(&sct.t[1], 11)
	fillDoubleDelimiter(&sct.t[1], 18)
	// Node 2
	fillAll(&sct.t[2], 12)
	fillAlphabet(&sct.t[2], 2)
	fillNumber(&sct.t[2], 2)
	sct.t[2]['_'] = 2
	// Node 3
	fillAll(&sct.t[3], 14)
	fillNumber(&sct.t[3], 3)
	sct.t[3]['.'] = 4
	// Node 4
	fillAll(&sct.t[4], 13)
	fillNumber(&sct.t[4], 4)
	// Node 5
	fillAll(&sct.t[5], 6)
	sct.t[5]['\''] = 7
	// Node 6
	sct.t[6]['\''] = 7
	// Node 7
	fillAll(&sct.t[7], 15)
	// Node 8
	fillAll(&sct.t[8], 9)
	sct.t[8]['"'] = 10
	// Node 9
	fillAll(&sct.t[9], 9)
	sct.t[9]['"'] = 10
	// Node 10
	fillAll(&sct.t[10], 16)
	// Node 11
	fillAll(&sct.t[11], 17)
	// Node 18
	fillAll(&sct.t[18], 19)
	fillDoubleDelimiter(&sct.t[18], 18)
}

func NewStateChangeTable() *StateChangeTable {
	var t StateChangeTable
	t.init()
	return &t
}
