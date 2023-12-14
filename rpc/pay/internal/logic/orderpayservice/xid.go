package orderpayservicelogic

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 32 位用于存储时间戳(秒)信息，最大可以存储的值为 4294967296，即 2106-02-07 14:28:16.
//
// 10 位用于存储数据节点信息，最大可以存储 1023，即可以有 1024 个节点.
//
// 21 位用于存储每一秒产生的序列号信息，最大可以存储 2097151，即每秒最大可以生产 2097152 个序列号.

const (
	kSequenceBits uint8 = 21 // 序列号占用的位数
	kDataNodeBits uint8 = 10 // 数据节点占用的位数

	kMaxSequence int64 = -1 ^ (-1 << kSequenceBits) // 序列号最大值，存储范围为 0-2097151
	kMaxDataNode int64 = -1 ^ (-1 << kDataNodeBits) // 数据节点最大值，存储范围为 0-1023

	kTimeShift     = kDataNodeBits + kSequenceBits // 时间戳向左的偏移量
	kDataNodeShift = kSequenceBits                 // 数据节点向左的偏移量

	kDataNodeMask = kMaxDataNode << kSequenceBits
)

const (
	MaxDataNode = kMaxDataNode
)

var (
	ErrDataNodeNotAllowed = errors.New(fmt.Sprintf("xid: data node can't be greater than %d or less than 0", kMaxDataNode))
)

type Option func(*XID) error

// WithDataNode 设置数据节点标识
func WithDataNode(node int64) Option {
	return func(x *XID) error {
		if node < 0 || node > kMaxDataNode {
			return ErrDataNodeNotAllowed
		}
		x.node = node
		return nil
	}
}

// WithTimeOffset 设置时间偏移量
func WithTimeOffset(t time.Time) Option {
	return func(x *XID) error {
		if t.IsZero() {
			return nil
		}
		x.timeOffset = t.Unix()
		return nil
	}
}

type XID struct {
	mu         sync.Mutex
	second     int64 // 上一次生成 id 的时间戳（秒）
	node       int64 // 数据节点
	sequence   int64 // 当前秒已经生成的 id 序列号 (从0开始累加)
	timeOffset int64 // 时间偏移量
}

func New(opts ...Option) (*XID, error) {
	var x = &XID{}
	x.second = 0
	x.node = 0
	x.sequence = 0
	x.timeOffset = 0

	var err error
	for _, opt := range opts {
		if err = opt(x); err != nil {
			return nil, err
		}
	}

	return x, nil
}

func (x *XID) DataNode() int64 {
	return x.node
}

func (x *XID) TimeOffset() int64 {
	return x.timeOffset
}

func (x *XID) Next() int64 {
	x.mu.Lock()

	var second = time.Now().Unix()
	if second < x.second {
		x.mu.Unlock()
		return -1
	}

	if x.second == second {
		x.sequence = (x.sequence + 1) & kMaxSequence
		if x.sequence == 0 {
			second = x.getNextSecond()
		}
	} else {
		x.sequence = 0
	}
	x.second = second
	var sequence = x.sequence
	x.mu.Unlock()

	var id = (second-x.timeOffset)<<kTimeShift | (x.node << kDataNodeShift) | (sequence)
	return id
}

func (x *XID) getNextSecond() int64 {
	var second = time.Now().Unix()
	for second < x.second {
		second = time.Now().Unix()
	}
	return second
}

// Time 获取 id 的时间，单位是 second
func Time(s int64) int64 {
	return s >> kTimeShift
}

// DataNode 获取 id 的数据节点标识
func DataNode(s int64) int64 {
	return s & kDataNodeMask >> kDataNodeShift
}

// Sequence 获取 id 的序列号
func Sequence(s int64) int64 {
	return s & kMaxSequence
}

var shared *XID
var once sync.Once

func Next() int64 {
	once.Do(func() {
		shared, _ = New()
	})
	return shared.Next()
}

func Init(opts ...Option) (err error) {
	once.Do(func() {
		shared, err = New(opts...)
	})

	if err != nil {
		once = sync.Once{}
	}

	return err
}

func Default() *XID {
	return shared
}
