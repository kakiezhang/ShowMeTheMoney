package utils

import (
	"bufio"
	"os"
)

func GetStandardInput() []byte {
	var rs string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rs += scanner.Text()
	}
	return []byte(rs)
}
