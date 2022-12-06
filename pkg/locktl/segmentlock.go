package goutils

import (
	"math"
	"sync"
)

type SegmentLock struct {
	shard      []*shard
	shardCount int
}

type shard struct {
	Mu sync.Mutex // 各个分片Map各自的锁
}

//该参数转成二进制，每个位都赋为1
func computeCapacity(param int) int {
	if param <= 16 {
		return 16
	}
	n := param - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return math.MaxInt32
	}
	return n + 1
}

func MakeSegmentLock(sct int) SegmentLock {
	s := SegmentLock{}
	capacity := computeCapacity(sct)
	s.shardCount = capacity
	s.shard = make([]*shard, capacity)
	for i := 0; i < capacity; i++ {
		s.shard[i] = &shard{
			Mu: sync.Mutex{},
		}
	}
	return s
}

func (s SegmentLock) Lock(key string) {
	sh := s.getShard(key)
	sh.Mu.Lock()
}

func (s SegmentLock) Unlock(key string) {
	sh := s.getShard(key)
	sh.Mu.Unlock()
}

func (s SegmentLock) getShard(key string) *shard {
	return s.shard[uint(fnv32(key))%uint(s.shardCount)]
}

// FNV hash
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
