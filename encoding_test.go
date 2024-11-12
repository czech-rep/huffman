package main

import (
	"reflect"
	"testing"
)

func TestBinaryStrToBytes(t *testing.T) {
	if result := BinaryToBytes("11111111"); !reflect.DeepEqual(result, []byte{255, 8}) {
		t.Errorf("mismatch bytes %v", result)
	}
	if result := BinaryToBytes("1111111110000000"); !reflect.DeepEqual(result, []byte{255, 128, 8}) {
		t.Errorf("mismatch bytes %v", result)
	}
	if result := BinaryToBytes("111"); !reflect.DeepEqual(result, []byte{7, 3}) {
		t.Errorf("mismatch bytes %v", result)
	}
	if result := BinaryToBytes("000"); !reflect.DeepEqual(result, []byte{0, 3}) {
		t.Errorf("mismatch bytes %v", result)
	}
	if result := BinaryToBytes(""); !reflect.DeepEqual(result, []byte{}) {
		t.Errorf("mismatch bytes %v", result)
	}
}

func TestBytesToBits(t *testing.T) {
	if result, _ := BytesToBits([]byte{255, 8}); result != "11111111" {
		t.Errorf("mismatch string 1 %v", result)
	}
	if result, _ := BytesToBits([]byte{255, 2}); result != "11" {
		t.Errorf("mismatch string 2 %v", result)
	}
	if result, _ := BytesToBits([]byte{0, 2}); result != "00" {
		t.Errorf("mismatch string 3 %v", result)
	}
	if result, _ := BytesToBits([]byte{128, 0, 2}); result != "1000000000" {
		t.Errorf("mismatch string 4 %v", result)
	}
	if result, _ := BytesToBits([]byte{32, 0, 2}); result != "0010000000" {
		t.Errorf("mismatch string 5 %v", result)
	}
	if result, _ := BytesToBits([]byte{}); result != "" {
		t.Errorf("mismatch string 6 %v", result)
	}
}
func TestEncodeAndDecode(t *testing.T) {
	inputs := []string{
		"101010101010110101101",
		"00000000000000000000000000000",
		"11111111111111111111111111111",
		"",
	}
	for _, input := range inputs {
		asBytes := BinaryToBytes(input)
		decoded, _ := BytesToBits(asBytes)
		if decoded != input {
			t.Errorf("Fail: expected %v got %v", input, decoded)
		}
	}
}
