package goutils

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
)

// ID is unique primary key
type ID [18]byte

// NewID is new ID
func NewID(prefix byte) ID {
	var ret ID
	id := uuid.New()
	ret[0] = prefix
	ret[1] = '-'
	for i, v := range id {
		ret[2+i] = v
	}
	return ret
}

var newPrefix = []byte{0, 0}

// IsNew is New ID
func (id ID) IsNew() bool {
	return bytes.HasPrefix(id[:], newPrefix)
}

// String is ID to string
func (id ID) String() string {
	if id.IsNew() {
		return ""
	}
	var u uuid.UUID
	for i := 2; i < 18; i++ {
		u[i-2] = id[i]
	}
	uuidStr := shortuuid.DefaultEncoder.Encode(u)
	ret := make([]byte, len(uuidStr)+2)
	ret[0] = id[0]
	ret[1] = id[1]
	for i, v := range []byte(uuidStr) {
		ret[i+2] = v
	}
	return string(ret)
}

// Parse is string to ID
func (id *ID) Parse(str string) (err error) {
	uuid, err := shortuuid.DefaultEncoder.Decode(str[2:])
	if err != nil {
		return
	}
	id[0] = str[0]
	id[1] = str[1]
	for i, v := range uuid {
		id[2+i] = v
	}
	return
}

// UnmarshalJSON is json Unmarshal
func (id *ID) UnmarshalJSON(data []byte) (err error) {
	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return
	}
	return id.Parse(str)
}

// MarshalJSON is json Marshal
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}
