package main

import (
	"fmt"
	"github.com/antlabs/timer"
	"log"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

func main1() {
	tm := timer.NewTimer()
	tm.AfterFunc(1*time.Second, func() {
		log.Printf("after\n")
	})

	t10 := tm.AfterFunc(5*time.Second, func() {
		log.Printf("after\n")
	})
	t10.Stop()
	fmt.Printf("=========stop\n")

	tm.ScheduleFunc(1*time.Second, func() {
		log.Printf("schedule\n")
	})

	fmt.Printf("=========\n")
	//tm.Stop()
	go tm.Run()
	fmt.Printf(">>>>>>\n")
	time.Sleep(time.Second * 1000)
	//list.New()
}
func main2() {
	//l := list.New()
}

var (
	e  int64 = 0
	n  int64 = 0
	wg sync.WaitGroup
)

//func main2() {
//	tt := map[string]*int64{
//		"test1": new(int64),
//		"test2": new(int64),
//	}
//	//开启10个goroutine
//	t1, _ := tt["test1"]
//	t2, _ := tt["test2"]
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			atomic.AddInt64(&e, 1)
//			atomic.AddInt64(t1, 1)
//			n += 1
//			*t2 += 1
//		}()
//	}
//
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			atomic.AddInt64(&e, 1)
//			atomic.AddInt64(t1, 1)
//			n += 1
//			*t2 += 1
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println("e : ", e, " n : ", n, " t1 :", *t1, " t2 :", *t2, "tt[test1]", *(tt["test1"]))
//}

func main3() {
	tt := map[string]*int64{
		"test1": new(int64),
	}
	//开启10个goroutine
	t1, _ := tt["test1"]
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1, _ := tt["test1"]
			atomic.AddInt64(t1, 1)
		}()
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t1, _ := tt["test1"]
			atomic.AddInt64(t1, 1)
		}()
	}

	wg.Wait()
	fmt.Println(" t1 :", *t1)
	//uintptr()
}

type mp struct {
	mu sync.Mutex
	m  map[int]string
}

func main4() {
	//n := 0
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}
	m := sync.Map{} //
	//rwmu := sync.RWMutex{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, "ss") //= "ss"
			//fmt.Printf("%v ===map(%v),\n", i, m.)
		}(i)
	}
	wg.Wait()
	//for i := 0; i < 5; i++ {
	//	res, _ := m.Load(i)
	//	fmt.Printf(">>>>%v\n", res)
	//}
	//act, load := m.LoadOrStore(6, "daa")
	//fmt.Printf("actual:%v,load:%v\n", act, load)
	fmt.Printf("===map(%v),\n", m)

	//ch := make(chan int, 2)
	//go func() {
	//	ch <- 7
	//	fmt.Printf("++++++++++++++\n")
	//}()
	//time.Sleep(2 * time.Second)
	//n := <-ch
	//time.Sleep(1 * time.Second)
	//fmt.Printf("###%v\n", n)
}
func main5() {
	//n := 0
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}
	m := map[int]string{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mtx.Lock()
			m[i] = "33"
			fmt.Printf("%v ===map(%v),\n", i, m)
			mtx.Unlock()
		}(i)
	}
	wg.Wait()
	//for i := 0; i < 5; i++ {
	//	res, _ := m.Load(i)
	//	fmt.Printf(">>>>%v\n", res)
	//}
	//act, load := m.LoadOrStore(6, "daa")
	//fmt.Printf("actual:%v,load:%v\n", act, load)
	fmt.Printf("===map(%v),\n", m)

	//ch := make(chan int, 2)
	//go func() {
	//	ch <- 7
	//	fmt.Printf("++++++++++++++\n")
	//}()
	//time.Sleep(2 * time.Second)
	//n := <-ch
	//time.Sleep(1 * time.Second)
	//fmt.Printf("###%v\n", n)

}
func main() {

	str := "12:00 a.m."
	r, _ := regexp.Compile(`^\d{1,2}:\d{1,2}(\s)(pm|PM|a.m.|A.M.|am|AM|p.m.|P.M.)`)
	idx := r.FindStringIndex(str)
	fmt.Println("test", idx)

}
