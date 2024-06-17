package resp

// import (
// 	"errors"
// 	"strconv"
// )

// func DecodeBulkString(rb []byte) (string, error) {
// 	if rb[0] != '$' {
// 		return "", errors.New("invalid Resp Bulkstring")
// 	}
// 	bufOffset := 1
// 	sizeStr, i := decodeCRLFString(rb[bufOffset:])
// 	n, err := strconv.Atoi(sizeStr)
// 	if err != nil {
// 		return "", err
// 	}
// 	bufOffset += i
// 	msgStr, _ := decodeCRLFString(rb[bufOffset:])
// 	if len(msgStr) != n {
// 		return "", errors.New("Decode Resp Bulkstring failed")
// 	}

// 	return msgStr, nil
// }

// func DecodeSimpleError(rb []byte) error {
// 	if rb[0] != '-' {
// 		return nil
// 	}
// 	ss, err := DecodeString(rb)
// 	if err != nil {
// 		return nil
// 	}
// 	return errors.New(ss)
// }

// func DecodeRespArray(rb []byte) ([]interface{}, error) {
// 	var arr []interface{}
// 	if rb[0] != '*' {
// 		return arr, errors.New("invalid Resp bool")
// 	}
// 	ss, err := DecodeString(rb)
// 	if err != nil {
// 		return arr, err
// 	}
// 	arrLen, err := strconv.Atoi(ss)
// 	if err != nil {
// 		return arr, err
// 	}
// 	buffOffset := 3 + len(ss)
// 	for i := 0; i < arrLen; i++ {
// 		var value interface{}
// 		if value, err = DecodeBoolean(rb[buffOffset:]); err == nil {
// 			arr = append(arr, value)
// 			buffOffset += 4
// 		} else if err = DecodeSimpleError(rb[buffOffset:]); err != nil {
// 			arr = append(arr, err)
// 			buffOffset += 3 + len(err.Error())
// 		} else if value, err = DecodeBulkString(rb[buffOffset:]); err == nil {
// 			arr = append(arr, value)
// 			s, _ := value.(string)
// 			strlen := len([]byte(s))
// 			lenStr := strconv.Itoa(strlen)
// 			buffOffset += 1 + len(lenStr) + 4 + strlen
// 		} else if value, err = DecodeInteger(rb[buffOffset:]); err == nil {
// 			arr = append(arr, value)
// 			s, _ := value.(int)
// 			lenStr := strconv.Itoa(s)
// 			buffOffset += 3 + len(lenStr)
// 		} else {
// 			break
// 		}
// 	}
// 	return arr, nil
// }

// func EncodeBulkString(ss string) ([]byte, error) {
// 	var byteSlice = []byte(ss)
// 	lenStr := strconv.Itoa(len(ss))
// 	if len(ss) == 0 {
// 		lenStr = strconv.Itoa(-1)
// 	}
// 	var wb = make([]byte, len(byteSlice)+3)
// 	wb[0] = '$'
// 	msg, err := EncodeString(lenStr, wb)
// 	if err != nil {
// 		return byteSlice, err
// 	}
// 	msg, err = EncodeString(ss, msg)
// 	return msg, err
// }

// func EncodeSimpleError(ss string) ([]byte, error) {
// 	var byteSlice = []byte(ss)
// 	var wb = make([]byte, len(byteSlice))
// 	wb[0] = '-'
// 	return EncodeString(ss, wb)
// }

// func EncodeRespArray(array []interface{}) ([]byte, error) {
// 	var wb = make([]byte, 1)
// 	wb = append(wb, '*')
// 	var arrayLen = len(array)
// 	lenstr := strconv.Itoa(arrayLen)
// 	EncodeString(lenstr, wb)
// 	for _, elem := range array {
// 		switch t := elem.(type) {
// 		case bool:
// 			bf, _ := EncodeBoolean(t)
// 			wb = append(wb, bf...)
// 		case int:
// 			bf, _ := EncodeInteger(t)
// 			wb = append(wb, bf...)
// 		case string:
// 			bf, _ := EncodeBulkString(t)
// 			wb = append(wb, bf...)
// 		case error:
// 			bf, _ := EncodeSimpleError(t.Error())
// 			wb = append(wb, bf...)
// 		default:

// 		}
// 	}
// 	return wb, nil
// }

// // func main() {
// // 	fmt.Println(DecodeSimpleString([]byte("+HELLO\r\n")))
// // 	b, _ := EncodeSimpleString("HELLO")
// // 	fmt.Println(string(b))

// // 	fmt.Println(DecodeSimpleError([]byte("-ERROR TEST\r\n")))
// // 	eb, _ := EncodeSimpleError("ERROR ETST")
// // 	fmt.Println(string(eb))

// // 	fmt.Println(DecodeInteger([]byte(":-100\r\n")))
// // 	ei, _ := EncodeInteger(-100222)
// // 	fmt.Println(string(ei))
// // 	fmt.Println(DecodeBulkString([]byte("$4\r\nECHO\r\n")))
// // 	ebs, _ := EncodeBulkString("hello")
// // 	fmt.Println(string(ebs))
// // 	fmt.Println(DecodeBulkString([]byte("$0\r\n\r\n")))
// // 	fmt.Println(EncodeBulkString(""))
// // 	ebs, _ = EncodeBulkString("")
// // 	fmt.Println(string(ebs))

// // 	fmt.Println(DecodeBoolean([]byte("#z\r\n")))
// // 	ebool, _ := EncodeBoolean(true)
// // 	fmt.Println(string(ebool))
// // 	err := errors.New("ERROR- NOT YET DONE")
// // 	eArray, _ := EncodeRespArray([]interface{}{"hello", true, "server", err})
// // 	fmt.Println(string(eArray))
// // 	fmt.Println(DecodeRespArray([]byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n")))
// // }
// /*
// func Decode([]byte) (datatype, error) {}
// func encode(datatype) ([]byte, error) {}
// */
