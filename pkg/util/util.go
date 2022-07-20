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

// Creates a string representation of a number in hexadecimal format
func ConvertToHexUint8(num uint8) string {
	// Convert the number to hex
	converted := fmt.Sprintf("%X", num)
	// Add leading 0s if needed
	if len(converted) < 2 {
		diff := 2 - len(converted)
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

// Sign extends a uint32 to a uint64
func SignExtend(val uint32, sizeOfVal int) uint64 {
	// Get the sign
	sign := uint64(val >> sizeOfVal - 1)
	longSign := uint64(0)
	// Repeat it 32 times
	for i := 0; i < 64 - sizeOfVal; i++ {
		longSign = longSign << 1 | sign
	}
	// Combine the original value with the long sign
	result := longSign << (64 - sizeOfVal) | uint64(val)
	return result
}