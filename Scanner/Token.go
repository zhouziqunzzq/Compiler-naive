package Scanner

type TokenType int

const (
	KEYWORD         = 1
	DELIMITER       = 2
	CHARCONSTANT    = 3
	STRINGCONSTANT  = 4
	INTEGERCONSTANT = 5
	FLOATCONSTANT   = 6
	IDENTIFIER      = 7
	ERROR           = -1
	END             = -2
)

var TokenTypeName = map[TokenType]string{
	1:  "KEYWORD",
	2:  "DELIMITER",
	3:  "CHARCONSTANT",
	4:  "STRINGCONSTANT",
	5:  "INTEGERCONSTANT",
	6:  "FLOATCONSTANT",
	7:  "IDENTIFIER",
	-1: "ERROR",
	-2: "END",
}

type Token struct {
	Type TokenType
	Word string
	ID   int
}
