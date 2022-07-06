package util

import (
	"fmt"
	"strings"
)

// Creates a string representation of a number in hexadecimal format
func ConvertToHex(num uint32, minLength int) string {
	// Convert the number to hex
	converted := fmt.Sprintf("%X", num)
	// Add leading 0s if needed
	if len(converted) < minLength {
		diff := minLength - len(converted)
		pad := strings.Repeat("0", diff)
		converted = pad + converted
	}
	// Add the '0x' prefix and make it all uppercase
	converted = "0x" + converted
	return converted
}