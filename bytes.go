package goutils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

// Decode is bytes to obj.
func Decode(bs []byte, obj interface{}) error {
	buff := bytes.NewBuffer(bs)
	dec := gob.NewDecoder(buff)
	return dec.Decode(obj)
}

// Encode is obj to bytes.
func Encode(obj interface{}) ([]byte, error) {
	var bs bytes.Buffer
	enc := gob.NewEncoder(&bs)
	err := enc.Encode(obj)
	return bs.Bytes(), err
}

// PrefixBytes is prefix add bytes to bytes.
func PrefixBytes(prefix string, bs []byte) []byte {
	buf := bytes.NewBuffer([]byte(prefix))
	buf.Write(bs)
	return buf.Bytes()
}

// PrefixUint32 is prefix add uint32 to bytes.
func PrefixUint32(prefix string, num uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, num)
	return PrefixBytes(prefix, bs)
}
