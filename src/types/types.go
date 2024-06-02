package types

import (
	"errors"
	"fmt"
)

type datatype interface {
	PrintData()
}

type SimpleString struct {
	data string
}

func (ss SimpleString) PrintData() {
	fmt.Printf("type %T, data %v", ss.data, ss.data)

}

type SimpleStringP struct {
}

type Protocol interface {
	Decode([]byte) (datatype, error)
	encode(datatype) ([]byte, error)
}

func (r SimpleStringP) decode(bytes []byte) (datatype, error) {
	var ss SimpleString
	i := 1
	for i < len(bytes) {
		if bytes[i] == '\r' && bytes[i+1] == '\n' {
			break
		}
		if bytes[i] == '\r' && bytes[i+1] != '\n' {
			return ss, errors.New("invalid Resp Simplestring")
		}
		if bytes[i] == '\n' {
			return ss, errors.New("invalid Resp Simplestring")
		}
		i++
	}
	ss.data = string(bytes[1:i])
	return ss, nil
}
