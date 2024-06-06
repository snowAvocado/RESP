package resp

import (
	"fmt"
	"testing"

	"github.com/snowAvocado/resp/types"
)

// test that verifies Encodes & Decodes SimpleString
func TestEncodeDecodeSimpleString(t *testing.T) {
	ss, err := DecodeSimpleString([]byte("+HELLO\r\n"))
	if err != nil && ss.Data != "HELLO" {
		t.Errorf("decode simple string failed")
	}

	ss_data := "HELLO"
	bf, _ := EncodeSimpleString(types.SimpleString{Data: ss_data})
	fmt.Println(string(bf))
	expected_bf := []byte{'+', 'H', 'E', 'L', 'L', 'O', '\r', '\n'}
	if string(expected_bf) == string(bf) {
		t.Errorf("encode simple string failed")
	}
}
