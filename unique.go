package goutils

import (
	"fmt"
)

var chId chan uint32

func UniqueId(prefix string) string {
	if chId == nil {
		makeChId()
	}
	return fmt.Sprintf("%s%d", prefix, <-chId)
}

func UniqueUint32() uint32 {
	if chId == nil {
		makeChId()
	}
	return <-chId
}

func makeChId() {
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
