package resp

// import "github.com/snowAvocado/resp/types"

// // Verbatim strings
// // This type is similar to the bulk string, with the addition of providing a hint about the data's encoding.

// // A verbatim string's RESP encoding is as follows:

// // =<length>\r\n<encoding>:<data>\r\n
// // An equal sign (=) as the first byte.
// // One or more decimal digits (0..9) as the string's total length,
// // in bytes, as an unsigned, base-10 value.
// // The CRLF terminator.
// // Exactly three (3) bytes represent the data's encoding.
// // The colon (:) character separates the encoding and data.
// // The data.
// // A final CRLF.

// func DecodeVerbatimString(buf []byte) (types.RespDatatype, error) {

// }

// func EncodeVerbatimString(datatype types.RespDatatype) ([]byte, error) {

// }
