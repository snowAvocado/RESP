package types

import (
	"fmt"
)

type RespDatatype interface {
	PrintData()
}

type SimpleString struct {
	Data string
}

func (ss SimpleString) PrintData() {
	fmt.Printf("RESP data type %T, data %v", ss.Data, ss.Data)

}
