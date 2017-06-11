package goutils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

func Encode(obj interface{}) []byte {
	var bs bytes.Buffer
	enc := gob.NewEncoder(&bs)
	enc.Encode(obj)
	return bs.Bytes()
}

func Decode(bs []byte, obj interface{}) {
	buff := bytes.NewBuffer(bs)
	dec := gob.NewDecoder(buff)
	dec.Decode(obj)
}

func PrefixBytes(prefix string, bs []byte) []byte {
	buf := bytes.NewBuffer([]byte(prefix))
	buf.Write(bs)
	return buf.Bytes()
}

func PrefixUint32(prefix string, num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	return PrefixBytes(prefix, bs)
}
