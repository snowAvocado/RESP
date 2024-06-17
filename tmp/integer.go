package resp

// import (
// 	"errors"
// 	"strconv"
// )

// func EncodeInteger(n int) ([]byte, error) {
// 	ss := strconv.Itoa(n)
// 	var byteSlice = []byte(ss)
// 	var wb = make([]byte, len(byteSlice))
// 	wb[0] = ':'
// 	return EncodeString(ss, wb)
// }

// func DecodeInteger(rb []byte) (int, error) {
// 	if rb[0] != ':' {
// 		return -1, errors.New("invalid Resp Integer")
// 	}

// 	ss, err := DecodeString(rb)
// 	if err != nil {
// 		return -1, err
// 	}
// 	integer, err := strconv.Atoi(ss)
// 	if err != nil {
// 		return -1, err
// 	}

// 	return integer, nil
// }
