package goutils

import (
	"log"
	"sort"
)

// Lesser 排序使用.
type Lesser interface {
	Less(b interface{}) bool
}

// Results 最优结果集.
type Results struct {
	Len    int
	lesses []Lesser
	size   int
}

// Add 增加结果.
func (r *Results) Add(data Lesser) {
	r.lesses[r.Len] = data
	log.Println(r)
	sort.Slice(r.lesses, func(i, j int) bool {
		if r.lesses[i] == nil {
			return false
		}
		if r.lesses[j] == nil {
			return false
		}
		return r.lesses[i].Less(r.lesses[j])
	})
	r.Len++
	if r.Len > r.size {
		r.Len = r.size
	}
}

// AddResults 增加结果集.
func (r *Results) AddResults(results *Results) {
	for i := 0; i < results.Len; i++ {
		r.Add(results.Get(i))
	}
}

// Get 获取数据.
func (r *Results) Get(i int) Lesser {
	return r.lesses[i]
}

// NewResults 新建结果集.
func NewResults(size int) *Results {
	return &Results{
		Len:    0,
		lesses: make([]Lesser, size+1),
		size:   size,
	}
}
