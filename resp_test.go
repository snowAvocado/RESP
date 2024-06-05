package resp

import (
	"fmt"
	"testing"

	"github.com/snowAvocado/resp/types"
)

// // test that verifies Encodes & Decodes SimpleString
// func TestEncodeDecodeSimpleString(t *testing.T) {
// 	ss, _ := DecodeSimpleString([]byte("+HELLO\r\n"))
// 	if ss != "HELLO" {
// 		t.Errorf("decode simple string failed")
// 	}

// 	bf, _ := EncodeSimpleString("HELLO")
// 	fmt.Println(string(bf))
// 	expected_bf := []byte{'+', 'H', 'E', 'L', 'L', 'O', '\r', '\n'}
// 	if string(expected_bf) == string(bf) {
// 		t.Errorf("encode simple string failed")
// 	}
// }

// test that verifies Encodes & Decodes SimpleString
func TestEncodeDecodeSimpleString(t *testing.T) {
	var ssp types.SimpleStringP

	sstype, _ := ssp.Decode([]byte("+HELLO\r\n"))
	ss, ok := sstype.(types.SimpleString)
	if ok && ss.Data != "HELLO" {
		t.Errorf("decode simple string failed")
	}

	bf, _ := EncodeSimpleString("HELLO")
	fmt.Println(string(bf))
	expected_bf := []byte{'+', 'H', 'E', 'L', 'L', 'O', '\r', '\n'}
	if string(expected_bf) == string(bf) {
		t.Errorf("encode simple string failed")
	}
}
