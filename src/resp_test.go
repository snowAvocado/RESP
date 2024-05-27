package RESP

import (
	"fmt"
	"testing"
)

// test that verifies Encodes & Decodes SimpleString
func TestEncodeDecodeSimpleString(t *testing.T) {
	ss, _ := DecodeSimpleString([]byte("+HELLO\r\n"))
	if ss != "HELLO" {
		t.Errorf("decode simple string failed")
	}

	bf, _ := EncodeSimpleString("HELLO")
	fmt.Println(string(bf))
	expected_bf := []byte{'+', 'H', 'E', 'L', 'L', 'O', '\r', '\n'}
	if string(expected_bf) == string(bf) {
		t.Errorf("encode simple string failed")
	}
}
