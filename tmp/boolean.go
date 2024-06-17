package resp

// import "errors"

// func EncodeBoolean(b bool) ([]byte, error) {
// 	var wb = make([]byte, 3)
// 	wb[0] = '#'
// 	var ss string
// 	if b {
// 		ss = "t"
// 	} else {
// 		ss = "f"
// 	}

// 	return EncodeString(ss, wb)
// }

// func DecodeBoolean(rb []byte) (bool, error) {
// 	if rb[0] != '#' {
// 		return false, errors.New("invalid Resp bool")
// 	}
// 	ss, err := DecodeString(rb)
// 	if err != nil {
// 		return false, err
// 	}
// 	if ss == "t" {
// 		return true, nil
// 	} else if ss == "f" {
// 		return false, nil
// 	} else {
// 		return false, errors.New("invalid Resp bool")
// 	}
// }
