package goutils

import (
	"fmt"
)

var chUint32 chan uint32

// UniqueString get unique string.
func UniqueString(prefix string) string {
	makeChUint32()
	return fmt.Sprintf("%s%d", prefix, <-chUint32)
}

// UniqueUint32 get unique uint32.
func UniqueUint32() uint32 {
	makeChUint32()
	return <-chUint32
}

func makeChUint32() {
	if chUint32 == nil {
		chUint32 = make(chan uint32)
		go func() {
			var num uint32
			num = 0
			for {
				num++
				chUint32 <- num
			}
		}()
	}
}
