package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var errBucketNum = errors.New("error bucket number")
var errDisableRate = errors.New("disabled rate")

var bucketDuration = int64(100)
var rateDuration = int64(1000)
var retry = 3

type rate struct {
	// 窗口大小，单位 毫秒， 默认1000
	timeDuration int64
	// 每个桶的时间窗口,单位 毫秒， 默认100
	bucketDuration int64
	// 桶最大数量
	bucketNum int
	// 当前桶数量
	bucketSize int32
	// 桶数组
	buckets []*bucket
	// 桶头指针
	headBucket int32
	// 桶尾指针
	tailBucket int32
	// 是否可用
	enabled bool
	m       sync.Mutex
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

//PrintValues 打印值
//TODO 获取快照后再输出->对对象的深拷贝
func (b *rate) PrintValues(start int64) {
	// var start = b.buckets[b.headBucket].startAt
	fmt.Println("start:", time.Unix(start/1e9, start%1e9).Format("2020-01-01 00:00:00.999999999"))
	var end = time.Now().UnixNano()
	fmt.Println("end:", time.Unix(end/1e9, end%1e9).Format("2020-01-01 00:00:00.999999999"))
	var count int32

	// var data, _ = json.Marshal(b)
	// var data1, _ = json.Marshal(b.buckets)

	// var copyObj = new(rate)
	// copyObj.buckets = make([]*bucket,copyObj.bucketNum)
	// copy(copyObj.buckets, b.buckets)
	// json.Unmarshal(data, copyObj)
	// json.Unmarshal(data1, copyObj.buckets)
	// err := deepCopy(copyObj, b)
	// if err != nil {
	// 	fmt.Printf("deep copy error:%v\n", err)
	// 	return
	// }
	for i := b.headBucket; i <= b.tailBucket; i++ {
		var bucketValue = b.buckets[i].getValue()
		fmt.Printf("bucket %d value:%d\n", i, bucketValue)
		count += bucketValue
	}
	fmt.Printf("total value:%d\n", count)
}

//AddValue 插入数据
//TODO 批量插入数据?
func (b *rate) AddValue(val int) error {
	// fmt.Println("add value start")
	//对外提供，需要判断是否可用
	if !b.enabled {
		return errDisableRate
	}
	var backet = b.getCurrentBucket()
	var err error
	for i := 0; i < retry; i++ {
		err = backet.addValue(int32(val))
		if err != errAddValue {
			return nil
		}
		time.Sleep(time.Microsecond * time.Duration((i+1)*(100+rand.Intn(100))))
	}
	return err
}

func (b *rate) getCurrentBucket() *bucket {
	b.m.Lock()
	defer b.m.Unlock()
	var currentTime = int64(time.Now().UnixNano() / 1e6)
	var bucket = b.tail()
	if bucket != nil && currentTime < atomic.LoadInt64(&bucket.startAt)+b.bucketDuration {
		return bucket
	}

	bucket = b.tail()
	if bucket == nil {
		return b.addBucket(currentTime)
	}
	for i := 0; i < b.bucketNum; i++ {
		bucket = b.tail()
		if currentTime < atomic.LoadInt64(&bucket.startAt)+b.bucketDuration {
			return bucket
		}
		if currentTime-(atomic.LoadInt64(&bucket.startAt)+b.bucketDuration) > b.timeDuration {
			// 超过最后一个桶的时间间隔大于整个rate的窗口，直接重建
			b.reset()
			return b.getCurrentBucket()
		}
		b.addBucket(atomic.LoadInt64(&bucket.startAt) + b.bucketDuration)
	}
	return b.tail()
}

func (b *rate) reset() error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	b.bucketSize = 0
	b.buckets = []*bucket{}
	b.headBucket = 0
	b.tailBucket = 0
	return nil
}

func (b *rate) tail() *bucket {
	if atomic.LoadInt32(&b.bucketSize) == 0 {
		return nil
	}
	return b.buckets[atomic.LoadInt32(&b.tailBucket)]
}

func (b *rate) Start() error {
	b.enabled = true
	b.addBucket(int64(time.Now().UnixNano() / 1e6))
	return nil
}

func (b *rate) addBucket(start int64) *bucket {
	var bucket = createBucket(start)
	var size = atomic.LoadInt32(&b.bucketSize)
	var tail = atomic.LoadInt32(&b.tailBucket)
	if size == 0 {
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&b.buckets[tail])), unsafe.Pointer(bucket))
		atomic.StoreInt32(&b.bucketSize, size+1)
	} else if size < int32(b.bucketNum) {
		atomic.StoreInt32(&b.tailBucket, (tail+1)%int32(b.bucketNum))
		tail = atomic.LoadInt32(&b.tailBucket)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&b.buckets[tail])), unsafe.Pointer(bucket))
		atomic.StoreInt32(&b.bucketSize, size+1)
	} else {
		atomic.StoreInt32(&b.tailBucket, (b.tailBucket+1)%int32(b.bucketNum))
		atomic.StoreInt32(&b.headBucket, (b.headBucket+1)%int32(b.bucketNum))
		// b.tailBucket, b.headBucket = (b.tailBucket+1)%int32(b.bucketNum), (b.headBucket+1)%int32(b.bucketNum)
	}
	return bucket
}

func newRate(bucketNum int) (*rate, error) {
	// 桶数量无法整除，导致桶数量无法覆盖整个窗口
	if rateDuration%int64(bucketNum) > 0 {
		return nil, errBucketNum
	}
	return &rate{
		timeDuration:   rateDuration,
		bucketNum:      bucketNum,
		bucketDuration: rateDuration / int64(bucketNum),
		buckets:        make([]*bucket, bucketNum),
	}, nil
}
