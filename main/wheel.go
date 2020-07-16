package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	//"github.txt.com/zzh20/timewheel"
	"github.com/ouqiang/timewheel"

	"time"
)

// 定义心跳包，设置心跳超时时间，处理函数
//var wheelHeartbeat = timewheel.New(time.Second*1, 30, func(data interface{}) {
//	c := data.(net.Conn)
//	log.Printf("timeout close conn:%v", c)
//	c.Close()
//})

func Put(data interface{}) {
	res := data.(DataT)
	out, _ := json.Marshal(&res)
	fmt.Printf("===%v\n", string(out))
}

type DataT struct {
	Index int
}

func main111() {

	// 初始化时间轮
	// 第一个参数为tick刻度, 即时间轮多久转动一次
	// 第二个参数为时间轮槽slot数量
	// 第三个参数为回调函数
	//tw := timewheel.New(1*time.Second, 3600, func(data interface{}) {
	//	// do something
	//
	//})

	tw := timewheel.New(1*time.Second, 3600, Put)

	// 启动时间轮
	tw.Start()

	// 添加定时器
	// 第一个参数为延迟时间
	// 第二个参数为定时器唯一标识, 删除定时器需传递此参数
	// 第三个参数为用户自定义数据, 此参数将会传递给回调函数, 类型为interface{}
	conn := "hh"
	tw.AddTimer(1*time.Second, conn, DataT{Index: 105626})
	for i := 0; i < 5; i++ {
		nn := strconv.Itoa(i)
		tw.AddTimer(2*time.Second, conn+nn, DataT{Index: 100 + i})
		time.Sleep(time.Millisecond * 500)
	}
	//for i := 0; i < 5; i++ {
	//	nn := strconv.Itoa(i)
	//	tw.RemoveTimer(conn + nn)
	//}
	// 删除定时器, 参数为添加定时器传递的唯一标识
	//tw.RemoveTimer(conn)

	// 停止时间轮
	//tw.Stop()
	time.Sleep(time.Second * 170)
	///select {}
}

func main() {
	var once sync.Once
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
