package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"sync"
	"time"
)

func sendMessage() {
	url := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	err = producer.Publish("test1", []byte("hello world##"))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//productor
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			sendMessage()
		}
		//time.Sleep(time.Second * 10)
	}()
	//consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		url := "127.0.0.1:4150"
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test1", "struggle", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()
	wg.Wait()
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	url := "192.168.5.95:4150"
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test", "struggle", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()

	waiter.Wait()
}

/////////////////////////////////////////////
var producer *nsq.Producer

// 主函数
func mainSend() {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	InitProducer(strIP1)

	running := true

	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}

		for err := Publish("test", command); err != nil; err = Publish("test", command) {
			//切换IP重连
			strIP1, strIP2 = strIP2, strIP1
			InitProducer(strIP1)
		}
	}
}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil", err)
}

///////////////////////////
//Nsq接收测试
// 消费者
type ConsumerT struct{}

// 主函数
func mainReceive() {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}

//处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 添加消费者接口

	//建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接
	// if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	//  panic(err)
	// }
}
