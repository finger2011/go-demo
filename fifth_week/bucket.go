package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

var bucketDataLength = 100
var errAddValue = errors.New("add bucket value error")
var errAddIndex = errors.New("add bucket index error")
var errAddAtomic = errors.New("add atomic error")

type bucket struct {
	startAt int64
	data    *bucketData
}

type bucketData struct {
	// 数据分散存储在数组中，防止对同一个数组的争抢
	length   int
	index    int32
	dataList []int32
	m        sync.Mutex
}

func (b *bucket) getValue() int32 {
	var count int32
	var tmpData = b.data.dataList
	for i := 0; i < b.data.length; i++ {
		count += tmpData[i]
	}
	return count
}

func (b *bucket) addValue(val int32) error {
	var oldIndex = atomic.LoadInt32(&b.data.index)
	var old = atomic.LoadInt32(&b.data.dataList[oldIndex])
	if !atomic.CompareAndSwapInt32(&b.data.dataList[oldIndex], old, old+val) {
		return errAddValue
	}
	if !atomic.CompareAndSwapInt32(&b.data.index, oldIndex, (oldIndex+1)%int32(b.data.length)) {
		return errAddIndex
	}
	return nil
}

func createBucket(start int64) *bucket {
	return &bucket{
		startAt: start,
		data: &bucketData{
			length:   bucketDataLength,
			dataList: make([]int32, bucketDataLength),
			index:    0,
		},
	}
}
