package goutils

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"hash"
	"hash/fnv"
)

// 文件唯一标识
type FileId struct {
	hash hash.Hash
	size int64
}

func NewFileId(file string) (*FileId, error) {
	id := new(FileId)
	id.hash = fnv.New128()
	id.size = 0
	if file == "" {
		return id, nil
	}
	err := ReadBuf(file, func(bs []byte) { id.Write(bs) })
	return id, err
}

func (f *FileId) Write(data []byte) (int, error) {
	f.size += int64(len(data))
	return f.hash.Write(data)
}

func (f *FileId) Id() []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(f.size))
	return bytes.Join([][]byte{
		f.hash.Sum(nil),
		removeVacant(bs),
	}, []byte(""))
}
func (f *FileId) String() string {
	return base64.StdEncoding.EncodeToString(f.Id())
}
func removeVacant(bytes []byte) []byte {
	l := len(bytes)
	for idx, _ := range bytes {
		if bytes[l-idx-1] != 0 {
			return bytes[:l-idx]
		}
	}
	return bytes
}
