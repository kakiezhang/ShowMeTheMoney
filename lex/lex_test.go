package lex

import "testing"

func TestLexiRecog(t *testing.T) {
	var line string
	var token Token

	line = "age>=45;"
	line = "age >45;"
	line = "int age= 45;"
	line = "int1 = 45;"
	line = "2+3+4+5;"
	token = LexiRecog(line)
	t.Logf("line: %s, token: %+v", line, token)

	for _, v := range tokens {
		t.Logf("%+v", v)
	}
}
