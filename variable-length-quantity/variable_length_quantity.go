package variablelengthquantity

import "errors"

//DecodeVarint decodes a series of bytes into a series of uint32s via VLQ
func DecodeVarint(bytes []byte) (output []uint32, err error) {
	var i uint32
	if 0x80&bytes[len(bytes)-1] != 0 {
		return output, errors.New("incomplete sequence")
	}
	for _, b := range bytes {
		i <<= 7
		i |= uint32(0x7f & b)
		if 0x80&b == 0 {
			output = append(output, i)
			i = 0
		}
	}

	return output, nil
}

//EncodeVarint encodes a series of uint32s into a series of bytes via VLQ
func EncodeVarint(uint32s []uint32) (output []byte) {
	for _, n := range uint32s {
		bytes := []byte{}
		for ok := true; ok; ok = n > 0 {
			b := 0x7f & byte(n)
			if len(bytes) != 0 {
				b |= 0x80
			}
			bytes = append([]byte{b}, bytes...)
			n >>= 7
		}
		output = append(output, bytes...)
	}
	return
}
