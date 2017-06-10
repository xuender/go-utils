package goutils

import (
	"bytes"
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
