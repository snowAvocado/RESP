package resp

func Decode(buf []byte) (Datatype, error) {
	if len(buf) == 0 {
		panic("Decode: input empty buffer")
	}

	switch buf[0] {
	case '+':
		ss, err := DecodeSimpleString(buf)
		return ss, err
	default:
		panic("Decode : invalid RESP string")
	}

}

func Encode(datatype Datatype) ([]byte, error) {
	switch t := datatype.(type) {
	case SimpleString:
		buf, err := EncodeSimpleString(t)
		return buf, err
	default:
		panic("Encode : invalid RESP datatype")
	}

}
