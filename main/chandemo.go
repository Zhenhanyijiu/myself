package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Data interface{}
}

type Err struct {
	Errcode int
	Errmsg  string
}

func WriteAsync(errChan chan Err, err Err) {
	select {
	case errChan <- err:
	default:
	}
}

func ReadAsync(errChan chan Err) *Err {
	select {
	case e := <-errChan:
		return &e
	default:
	}
	return nil
}

func main() {

	producer := 10
	message := make(chan Message)
	err := make(chan Err, 1)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	//生产者
	wg.Add(producer)
	for i := 0; i < producer; i++ {

		go func(id int) {
			defer fmt.Printf("i=%v\n", id)
			defer wg.Done()
			//模拟正常业务使用过程中出错处理
			if false {
				WriteAsync(err, Err{Errcode: 0xff, Errmsg: "计算失败"})
				cancel()
				return
			}
			//发送正常业务数据
			select {
			case <-ctx.Done():
				fmt.Println(">>>>ctx.Done")
				return
			case message <- Message{Data: fmt.Sprintf("id = %d, %s\n", id, time.Now())}:
			}
		}(i)
	}

	//同步生产者，消费者
	var pc sync.WaitGroup

	pc.Add(2)
	defer pc.Wait()
	go func() {
		defer pc.Done()
		//等待所有go程结束
		wg.Wait()
		cancel()
		fmt.Printf("wait ok\n")
	}()

	//消费者
	func() {
		defer pc.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("####ctx.done---\n")
				e := ReadAsync(err)
				if e == nil { //正常退出
					fmt.Println("exit success")
				} else { //异常退出
					fmt.Printf("fail:%#v\n", e)
				}

				return
			case v := <-message:
				fmt.Printf("message = %s\n", v)
			}
		}
	}()
}
