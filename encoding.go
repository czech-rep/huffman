package main

import (
	"fmt"
)

// module will read bytes from file. Each byte is 8 bits
// we need to get ones and zeros string for the decoder
// for example 1 is "00000001" and 255 is "11111111"
// but the last byte is problematic, because we don't know the length of endocing
// encoder might have returned "11" as last byte, but "11" -> 3 -> "00000011" we would get wrong code
// 3 -> "11000000" is also wrong. So we have to save lenght of our code
// Solution append: last byte with length of significant bits in the previous byte

type InvalidCode string

func (msg InvalidCode) Error() string {
	return string(msg)
}

func BinaryToBytes(binaryStr string) []byte {
	result := []byte{}
	if len(binaryStr) == 0 {
		return result
	}
	var currentBits string
	for i := 0; i < len(binaryStr); i += 8 {
		end := i + 8
		if end > len(binaryStr) {
			end = len(binaryStr)
		}
		currentBits = binaryStr[i:end]
		var newByte byte
		for _, bit := range currentBits {
			newByte <<= 1
			if bit == '1' {
				newByte |= 1
			}
		}
		result = append(result, newByte)
	}
	// append last byte with information about singinicant bits in the previous
	return append(result, byte(len(currentBits)))
}
func BytesToBits(input []byte) (string, error) {
	// from a slice of bytes, returns "001101101011101110"
	result := ""
	if len(input) == 0 {
		return result, nil
	}
	if len(input) == 1 {
		return "", InvalidCode("Code cannot be of length one")
	}
	length := len(input)
	for i := 0; i < length-2; i++ {
		result += fmt.Sprintf("%08b", input[i])
	}
	lastByte := input[length-2]
	lastBitsLength := int(input[length-1])
	if lastBitsLength < 0 || lastBitsLength > 8 {
		msg := fmt.Sprintf("last byte must carry length of significant bits, invalid %v", lastBitsLength)
		return "", InvalidCode(msg)
	}

	result += fmt.Sprintf("%08b", lastByte)[8-lastBitsLength : 8]
	return result, nil
}
