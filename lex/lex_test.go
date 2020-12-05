package lex

import (
	"calculater_interpreter/token"
	"testing"
)

func TestLexiTokenize(t *testing.T) {
	var line string

	line = "age>=45;"
	line = "age >45;"
	line = "int age= 45;"
	line = "int1 = 45;"
	line = "2+3+4+5;"
	line = "2+3*4;"
	line = "2/3-4;"

	sl := NewSimpleLexi()
	var str token.TokenReader
	str = sl.Tokenize(line)
	t.Logf("line: %s, token: %+v",
		line, sl.token)

	for token := str.Read(); token != nil; {
		t.Logf("%+v", token)
		token = str.Read()
	}
}
