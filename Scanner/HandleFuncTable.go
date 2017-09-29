package Scanner

import (
	"strconv"
)

type HandleFuncTable struct {
	t map[int]func(s *Scanner)
}

func NewHandleFuncTable() *HandleFuncTable {
	var hft HandleFuncTable
	hft.t = make(map[int]func(s *Scanner))
	hft.t[1] = HandleEmpty
	hft.t[12] = HandleKeywordIdentifier
	hft.t[13] = HandleFloatConstant
	hft.t[14] = HandleIntegerConstant
	hft.t[15] = HandleCharConstant
	hft.t[16] = HandleStringConstant
	hft.t[17] = HandleSingleDelimiter
	hft.t[19] = HandleDoubleDelimiter
	return &hft
}

func HandleKeywordIdentifier(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	str := string(s.buffer)

	if i, ok := s.kt.T[str]; ok {
		// Keyword
		s.LastToken.Type = KEYWORD
		s.LastToken.ID = i
		s.LastToken.Word = str
	} else {
		// Identifier
		var newID int
		if i, ok := s.it.T[str]; ok {
			newID = i
		} else {
			newID = len(s.it.T) + 1
			s.it.T[str] = newID
		}
		s.LastToken.Type = IDENTIFIER
		s.LastToken.ID = newID
		s.LastToken.Word = str
	}
}

func HandleFloatConstant(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	str := string(s.buffer)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		s.LastToken.Type = ERROR
		return
	}
	var newID int
	if i, ok := s.floatct.T[f]; ok {
		newID = i
	} else {
		newID = len(s.floatct.T) + 1
		s.floatct.T[f] = newID
	}
	s.LastToken.Type = FLOATCONSTANT
	s.LastToken.ID = newID
	s.LastToken.Word = str
}

func HandleIntegerConstant(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	str := string(s.buffer)
	n, err := strconv.Atoi(str)
	if err != nil {
		s.LastToken.Type = ERROR
		return
	}
	var newID int
	if i, ok := s.intct.T[n]; ok {
		newID = i
	} else {
		newID = len(s.intct.T) + 1
		s.intct.T[n] = newID
	}
	s.LastToken.Type = INTEGERCONSTANT
	s.LastToken.ID = newID
	s.LastToken.Word = str
}

func HandleCharConstant(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a char
	var ch rune
	if len(s.buffer) > 2 {
		ch = s.buffer[1]
	}

	var newID int
	if i, ok := s.charct.T[ch]; ok {
		newID = i
	} else {
		newID = len(s.charct.T) + 1
		s.charct.T[ch] = newID
	}
	s.LastToken.Type = CHARCONSTANT
	s.LastToken.ID = newID
	s.LastToken.Word = string(s.buffer)
}

func HandleStringConstant(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	var str string
	if len(s.buffer) > 2 {
		str = string(s.buffer[1 : len(s.buffer)-1])
	}

	var newID int
	if i, ok := s.strct.T[str]; ok {
		newID = i
	} else {
		newID = len(s.strct.T) + 1
		s.strct.T[str] = newID
	}
	s.LastToken.Type = STRINGCONSTANT
	s.LastToken.ID = newID
	s.LastToken.Word = string(s.buffer)
}

func HandleSingleDelimiter(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	str := string(s.buffer)

	if i, ok := s.dt.T[str]; ok {
		// Delimiter
		s.LastToken.Type = DELIMITER
		s.LastToken.ID = i
		s.LastToken.Word = str
	} else {
		s.LastToken.Type = ERROR
	}
}

func HandleDoubleDelimiter(s *Scanner) {
	// Rewind the Scanner first
	s.Rewind()
	// Reset the Scanner in the end
	defer s.Reset()
	// Convert buffer to a string
	str := string(s.buffer)

	if i, ok := s.dt.T[str]; ok {
		// Delimiter
		s.LastToken.Type = DELIMITER
		s.LastToken.ID = i
		s.LastToken.Word = str
	} else {
		s.LastToken.Type = ERROR
	}
}

func HandleEmpty(s *Scanner) {
	// Reset the Scanner
	s.Reset()
}
