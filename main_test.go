package main

import (
	"testing"
)

var MessagePath = "/tmp/message"
var EncodedPath = "/tmp/encoded"
var TreePath = "/tmp/tree.huffman"

func TestEncodeDecode(t *testing.T) {
	message := `hello in tests1590,./<>?;':"[]{}-=_+`
	WriteToFile(MessagePath, message)

	EncodeCmd(MessagePath, EncodedPath, TreePath)
	decoded, err := DecodeCmd(EncodedPath, TreePath)

	if err != nil {
		t.Errorf("decode failed giving: %v", err)
		return
	}
	if decoded != message {
		t.Errorf("decode failed giving: %v instead of %v", decoded, message)
		return
	}
}
