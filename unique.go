package goutils

import (
	"fmt"
)

var chId chan uint32

func UniqueId(prefix string) string {
	makeChId()
	return fmt.Sprintf("%s%d", prefix, <-chId)
}

func UniqueUint32() uint32 {
	makeChId()
	return <-chId
}

func makeChId() {
	if chId == nil {
		chId = make(chan uint32)
		go func() {
			var id uint32
			id = 0
			for {
				id += 1
				chId <- id
			}
		}()
	}
}
