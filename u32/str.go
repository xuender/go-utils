package u32

import (
	"encoding/base64"
	"encoding/binary"
	"strings"
)

// Uint32ToStr uint32 to string.
func Uint32ToStr(num uint32) string {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, num)
	end := 4
	for i := 3; i > 0; i-- {
		if bytes[i] == 0 {
			end = i
		} else {
			break
		}
	}
	return strings.Trim(base64.StdEncoding.EncodeToString(bytes[0:end]), "=")
}

// StrToUint32 string to uint32.
func StrToUint32(str string) uint32 {
	for i := 0; i < len(str)%4; i++ {
		str = str + "="
	}
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return 0
	}
	bs := make([]byte, 4)
	for i := 0; i < len(bytes); i++ {
		bs[i] = bytes[i]
	}
	return binary.LittleEndian.Uint32(bs)
}
