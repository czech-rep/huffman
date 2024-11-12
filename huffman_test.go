package main

import "testing"

// a 11, b 10, c 01, d 00
var SimpleTree *Node = &Node{
	"", 2,
	&Node{"", 1, &Node{"a", 1, nil, nil}, &Node{"b", 1, nil, nil}},
	&Node{"", 1, &Node{"c", 1, nil, nil}, &Node{"d", 1, nil, nil}},
}

// Decode
func TestDecode(t *testing.T) {
	message := "110000011010"
	decoded, _ := Decode(SimpleTree, message)

	expected := "addcbb"
	if decoded != expected {
		t.Errorf("decode failed giving %v instead of %v", decoded, expected)
	}
}
func TestDecodeWrongCharacter(t *testing.T) {
	message := "2"
	_, err := Decode(SimpleTree, message)

	expected := "Unhandled character: 2"
	if err == nil || err.Error() != expected {
		t.Errorf("mismatch error %v", err)
	}

}
func TestDecodeInvalidCode(t *testing.T) {
	message := "111" // 11 is a and 1 is not in tree
	_, err := Decode(SimpleTree, message)

	expected := "Finished decoding inside tree"
	if err == nil || err.Error() != expected {
		t.Errorf("mismatch error %v", err)
	}
}

// Encode
func TestEncodeSimple(t *testing.T) {
	input := "aabbcdca"
	encoded, _ := Encode(SimpleTree, input)

	decoded, _ := Decode(SimpleTree, encoded)

	if decoded != input {
		t.Errorf("invalid decoded: %v", decoded)
	}
}

func TestEncodeInvalidCharacter(t *testing.T) {
	input := "e"
	_, err := Encode(SimpleTree, input)

	if err.Error() != "Unhandled character: e" {
		t.Errorf("mismatch error: %v", err.Error())
	}
}
