package utils

func IsDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func IsAlpha(c rune) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
