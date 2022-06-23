package goutils

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// ====================== 分段锁测试 ======================
var cnt1, cnt2, cnt3, cnt4 int
var mu SegmentLock

func init() {
	mu = MakeSegmentLock(3)
	rand.Seed(time.Now().UnixNano())
}

func AddCntN(iter int, key string) {
	mu.GetShard(key).Mu.Lock()
	defer mu.GetShard(key).Mu.Unlock()
	for i := 0; i < iter; i++ {
		switch key {
		case "A":
			cnt1++
		case "B":
			cnt2++
		case "C":
			cnt3++
		default:
			cnt4++
		}
	}
	fmt.Println("sleep start, key:" + key)
	time.Sleep(time.Second)
}

func TestSegmentLock(t *testing.T) {
	//f, _ := os.Create("TestSegmentLock.dat")
	//defer f.Close()
	//_ = trace.Start(f)
	//defer trace.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			switch rand.Int() % 4 {
			case 0:
				AddCntN(10, "A")
			case 1:
				AddCntN(10, "B")
			case 2:
				AddCntN(10, "C")
			default:
				AddCntN(10, "D")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt1)
	fmt.Println(cnt2)
	fmt.Println(cnt3)
	fmt.Println(cnt4)
}
