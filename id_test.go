package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	assert := assert.New(t)
	id := NewID('O')

	assert.Len(id, 18)
	assert.Len(id.String(), 24)
	assert.True(ID{}.IsNew())
	assert.False(id.IsNew())

	b, err := json.Marshal(id)

	assert.Len(id, 18)
	assert.Len(b, 26)
	assert.Nil(err)
}
func TestID_String(t *testing.T) {
	assert := assert.New(t)
	i := new(ID)
	err := i.Parse("I-Cekw67uyMpBGZLRP2HFVbe")

	assert.Nil(err)
	assert.EqualValues(i[0], 'I')
	assert.EqualValues(i[1], '-')
}
func TestID_ParseBytes(t *testing.T) {
	assert := assert.New(t)
	id := NewID('O')
	bs := id[:]
	i := new(ID)
	err := i.ParseBytes(bs)

	assert.Nil(err)
	assert.Len(i.String(), 24)
}
func TestID_UnmarshalJSON(t *testing.T) {
	assert := assert.New(t)
	i := new(ID)
	err := json.Unmarshal([]byte(`"I-Cekw67uyMpBGZLRP2HFVbe"`), i)

	assert.Nil(err)
	assert.Len(i[:], 18)
	assert.EqualValues(i[0], 'I')
	assert.Equal(i.String(), "I-Cekw67uyMpBGZLRP2HFVbe")
}

func ExampleNewID() {
	fmt.Println(ID{}.IsNew())
	id := NewID('A')
	fmt.Println(id.IsNew())

	// Output:
	// true
	// false
}

func BenchmarkNewID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewID('A')
	}
}
