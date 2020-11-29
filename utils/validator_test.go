package utils

import "testing"

func TestIsDigit(t *testing.T) {
	for _, v := range []rune{
		'a', '哈', '0', '9', '1'} {
		t.Logf("%c is digit? %t", v, IsDigit(v))
	}
}

func TestIsAlpha(t *testing.T) {
	for _, v := range []rune{
		'a', 'Z', '哈', '0', '9', '1'} {
		t.Logf("%c is digit? %t", v, IsAlpha(v))
	}
}
