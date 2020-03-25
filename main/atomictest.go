package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			//原子操作保证了这些goroutine之间没有数据竞争
			atomic.AddInt32(&n, 1)
			//n = n + 1
			wg.Done()
		}()
	}
	wg.Wait()
	//fmt.Println(atomic.LoadInt32(&n)) // 1000
	fmt.Printf("%v", n)
}
