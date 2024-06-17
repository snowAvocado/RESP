package resp

import (
	"errors"
	"fmt"
)

type SimpleString struct {
	Data string
}

func (ss SimpleString) PrintData() {
	fmt.Printf("RESP data type %T, value %v", ss.Data, ss.Data)

}

func EncodeSimpleString(ss SimpleString) ([]byte, error) {
	var wb = make([]byte, 1)
	wb[0] = '+'
	var byteSlice = []byte(ss.Data)
	for i := 0; i < len(ss.Data); {
		if byteSlice[i] == '\r' || byteSlice[i] == '\n' {
			return wb, errors.New("simple string should not contain \n or \r")
		}
		wb = append(wb, byteSlice[i])
		i++
	}
	wb = append(wb, '\r', '\n')
	return wb, nil
}

func DecodeSimpleString(rb []byte) (SimpleString, error) {
	var ss SimpleString

	byte_len := len(rb)
	cr_idx := byte_len - 2
	lf_idx := byte_len - 1

	if byte_len <= 3 || rb[cr_idx] != '\r' || rb[lf_idx] != '\n' {
		return ss, errors.New("invalid Resp Simple string")
	}

	i := 1
	for i < cr_idx {
		if rb[i] == '\r' || rb[i] == '\n' {
			return ss, errors.New("simple string buffer shall contain \n or \r only at the end")
		}
		i++
	}

	ss.Data = string(rb[1:cr_idx])
	return ss, nil

}
