package types

import (
	"errors"
	"fmt"
)

type datatype interface {
	PrintData()
}

type SimpleString struct {
	Data string
}

func (ss SimpleString) PrintData() {
	fmt.Printf("type %T, data %v", ss.Data, ss.Data)

}

type SimpleStringP struct {
}

type Protocol interface {
	Decode([]byte) (datatype, error)
	encode(datatype) ([]byte, error)
}

func (r SimpleStringP) Decode(bytes []byte) (datatype, error) {
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
	ss.Data = string(bytes[1:i])
	return ss, nil
}
