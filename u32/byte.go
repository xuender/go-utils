package u32

import "encoding/binary"

// Encode is uint32 to bytes.
func Encode(num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	return bs
}

// Decode is bytes to uint32.
func Decode(bytes []byte) uint32 {
	if len(bytes) != 4 {
		return 0
	}
	return binary.LittleEndian.Uint32(bytes)
}
