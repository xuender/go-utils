package goutils

import (
	"sort"
	"sync"
)

// Result 结果.
type Result struct {
	Data  interface{}
	Point interface{}
}

// Results 最优结果集.
type Results struct {
	Data []Result
	Len  int
	Size int

	less  func(i, j interface{}) bool
	equal func(i, j interface{}) bool
	mutex sync.Mutex
}

// Add 增加结果.
func (r *Results) Add(data, point interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i := 0; i < r.Len; i++ {
		if r.equal(data, r.Data[i].Data) {
			r.Data[i].Point = point
			return
		}
	}
	r.Data[r.Len] = Result{
		Data:  data,
		Point: point,
	}
	r.Len++
	sort.Slice(r.Data[:r.Len], func(i, j int) bool { return r.less(r.Data[i].Point, r.Data[j].Point) })
	if r.Len > r.Size {
		r.Len = r.Size
	}
}

// AddResults 增加结果集.
func (r *Results) AddResults(results *Results) {
	for i := 0; i < results.Len; i++ {
		r.Add(results.Get(i))
	}
}

// Get 获取数据.
func (r *Results) Get(i int) (interface{}, interface{}) {
	return r.Data[i].Data, r.Data[i].Point
}

// GetData 数据.
func (r *Results) GetData(i int) interface{} {
	return r.Data[i].Data
}

// GetPoint 得分.
func (r *Results) GetPoint(i int) interface{} {
	return r.Data[i].Point
}

// NewResults 新建结果集.
func NewResults(size int, less, equal func(i, j interface{}) bool) *Results {
	return &Results{
		Len:   0,
		Data:  make([]Result, size+1),
		Size:  size,
		less:  less,
		equal: equal,
	}
}
