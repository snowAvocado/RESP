package resp

import (
	"errors"

	"github.com/snowAvocado/resp/types"
)

func EncodeSimpleString(ss types.SimpleString) ([]byte, error) {
	var byteSlice = []byte(ss.Data)
	var wb = make([]byte, len(byteSlice))
	wb[0] = '+'
	return EncodeString(ss.Data, wb)
}

func EncodeString(ss string, wb []byte) ([]byte, error) {
	var byteSlice = []byte(ss)
	i := 0
	for i < len(ss) {
		if byteSlice[i] == '\r' && byteSlice[i+1] == '\n' {
			break
		}
		if byteSlice[i] == '\r' && byteSlice[i+1] != '\n' {
			return wb, errors.New("invalid Resp Simplestring")
		}
		if byteSlice[i] == '\n' {
			return wb, errors.New("invalid Resp Simplestring")
		}
		wb = append(wb, byteSlice[i])
		i++
	}
	wb = append(wb, '\r', '\n')
	return wb, nil
}

func DecodeString(rb []byte) (string, error) {
	i := 1
	for i < len(rb) {
		if rb[i] == '\r' && rb[i+1] == '\n' {
			break
		}
		if rb[i] == '\r' && rb[i+1] != '\n' {
			return "", errors.New("invalid Resp Simplestring")
		}
		if rb[i] == '\n' {
			return "", errors.New("invalid Resp Simplestring")
		}
		i++
	}

	return string(rb[1:i]), nil
}

func decodeCRLFString(rb []byte) (string, int) {
	i := 0
	for i < len(rb) {
		if rb[i] == '\r' && rb[i+1] == '\n' {
			break
		}
		if rb[i] == '\r' && rb[i+1] != '\n' {
			return "", 0
		}
		if rb[i] == '\n' {
			return "", 0
		}
		i++
	}
	return string(rb[0:i]), i + 2
}

func DecodeSimpleString(rb []byte) (types.SimpleString, error) {
	var ss types.SimpleString
	if rb[0] != '+' {
		return ss, errors.New("invalid Resp Simplestring")
	}
	ss_data, err := DecodeString(rb)
	if err != nil {
		ss.Data = ss_data
	}
	return ss, err
}
