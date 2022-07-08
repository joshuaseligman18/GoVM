package util

import (
	"fmt"
	"strings"
	"time"
)

// Creates a string representation of a number in hexadecimal format
func ConvertToHexUint32(num uint32) string {
	// Convert the number to hex
	converted := fmt.Sprintf("%X", num)
	// Add leading 0s if needed
	if len(converted) < 8 {
		diff := 8 - len(converted)
		pad := strings.Repeat("0", diff)
		converted = pad + converted
	}
	// Add the '0x' prefix
	converted = "0x" + converted
	return converted
}

// Creates a string representation of a number in hexadecimal format
func ConvertToHexUint64(num uint64) string {
	// Convert the number to hex
	converted := fmt.Sprintf("%X", num)
	// Add leading 0s if needed
	if len(converted) < 16 {
		diff := 16 - len(converted)
		pad := strings.Repeat("0", diff)
		converted = pad + converted
	}
	// Add the '0x' prefix
	converted = "0x" + converted
	return converted
}

// Returns the current time in milliseconds
func GetCurrentTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}