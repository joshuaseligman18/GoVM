package util

import (
	"fmt"
	"strings"
)

func ConvertToHex(num int64, minLength int) string {
	converted := fmt.Sprintf("%x", num)
	if len(converted) < minLength {
		diff := minLength - len(converted)
		pad := strings.Repeat("0", diff)
		converted = pad + converted
	}
	converted = "0x" + strings.ToUpper(converted)
	return converted
}