package goutils

import (
	"fmt"
)

var chId chan uint32

func UniqueId(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, <-chId)
}

func UniqueUint32() uint32 {
	return <-chId
}

func init() {
	chId = make(chan uint32)
	go func() {
		var id uint32
		id = 0
		for {
			chId <- id
			id += 1
		}
	}()
}
