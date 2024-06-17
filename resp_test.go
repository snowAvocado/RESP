package resp

import (
	"testing"
)

// test Encode & Decode SimpleString
func TestEncodeDecodeSimpleString(t *testing.T) {
	datatype, err := Decode([]byte("+HELLO\r\n"))
	ss := datatype.(SimpleString)
	if err != nil && ss.Data != "HELLO" {
		t.Errorf("decode simple string failed")
	}

	ss_data := "HELLO"
	buf, _ := Encode(SimpleString{Data: ss_data})
	expected_buf := []byte{'+', 'H', 'E', 'L', 'L', 'O', '\r', '\n'}
	if string(expected_buf) != string(buf) {
		t.Errorf("encode simple string failed %v != %v ", string(expected_buf), string(buf))
	}
}
