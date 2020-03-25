package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"myself/proto/streamtest"
	"time"
)

type client struct {
	addr string
}

//服务端流式
//func main() {
//	clnt := client{addr: "192.168.6.95:7777"}
//	conn, err := grpctest.DialContext(context.Background(), clnt.addr, grpctest.WithInsecure())
//	if err != nil {
//		fmt.Printf("error(%v)\n", err)
//	}
//	defer conn.Close()
//	gtcli := streamtest.NewGTestClient(conn)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	in := streamtest.InMsg{Name: "hudi"}
//	stream, err := gtcli.ExampleDeal(ctx, &in)
//	if err != nil {
//		fmt.Printf("client to server(%v)\n", err)
//	}
//	for {
//		res, err := stream.Recv()
//		if err == io.EOF {
//			fmt.Printf("recv eof(%v)\n", err)
//			break
//		}
//		if err != nil {
//			fmt.Printf("recv error(%v)\n", err)
//			return
//		}
//		fmt.Printf("res:(%v)\n", res.Great)
//	}
//}

//双流式
func mainTemp() {
	clnt := client{addr: "192.168.5.25:7777"}
	conn, err := grpc.DialContext(context.Background(), clnt.addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}
	defer conn.Close()
	gtcli := streamtest.NewGTestClient(conn)
	//todo:context has problem
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	end := make(chan struct{})
	stream, err := gtcli.ExampleDealStream(context.Background())
	if err != nil {
		fmt.Printf("client to server(%v)\n", err)
	}
	length := 10
	go func() {
		defer close(end)
		for num := 0; num < length; num++ {
			in := streamtest.InMsg{Name: fmt.Sprintf("hudi_%v", num+1)}
			if num == length-1 {
				in = streamtest.InMsg{Name: fmt.Sprintf("%v", "eof")}
			}
			err := stream.Send(&in)
			if err != nil {
				fmt.Printf("error(%v)\n", err)
				return
			}
		}
		stream.CloseSend()
	}()
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("recv eof(%v)\n", err)
			break
		}
		if err != nil {
			fmt.Printf("recv error(%v)\n", err)
			return
		}
		fmt.Printf("res:(%v)\n", res.Great)
	}
	<-end
}

func checkTimeout(stream streamtest.GTest_ExampleDealStreamClient, ctx context.Context, cancelFunc context.CancelFunc, m <-chan struct{}) {
	tm := time.NewTimer(time.Second * 5)
	defer fmt.Printf("===?????exit go routine...\n")
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("===ctx.done in check\n")
			return
		case <-tm.C:
			//stream.CloseSend()
			cancelFunc()
			fmt.Printf("===timeout\n")
		case <-m:
			tm.Reset(time.Second * 5)
		}
	}
}

func check(cancel context.CancelFunc, ctx context.Context, heartbeat chan struct{}) {
	tk := time.NewTimer(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("chancel\n")
			return
		case <-heartbeat:
			tk.Reset(time.Second)
		case <-tk.C:

			cancel()
		}
	}
}
func main() {
	clnt := client{addr: "192.168.5.25:7777"}
	conn, err := grpc.Dial(clnt.addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}
	defer conn.Close()
	gtcli := streamtest.NewGTestClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	//defer cancel()
	//end := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	stream, err := gtcli.ExampleDealStream(context.Background())
	if err != nil {
		fmt.Printf("client to server(%v)\n", err)
		return
	}
	m := make(chan struct{}, 1)

	go checkTimeout(stream, ctx, cancel, m)
	//go func() {
	//defer close(end)
	length := 10
	for num := 0; num < length; num++ {
		in := streamtest.InMsg{Name: fmt.Sprintf("hudi_%v", num+1)}
		if num == length-1 {
			in = streamtest.InMsg{Name: fmt.Sprintf("%v", "eof")}
		}
		err := stream.Send(&in)
		if err != nil {
			fmt.Printf("===send error(%v)\n", err)
			break
		}

		res, err := stream.Recv()
		//time.Sleep(time.Second * 1)
		//stream.CloseSend()
		//return
		fmt.Printf("===========after recv\n")
		if err == io.EOF {
			fmt.Printf("recv eof(%v)\n", err)
			break
		}
		if err != nil {
			fmt.Printf("recv error(%v)\n", err)
			break
		}
		select {
		case <-ctx.Done():
			fmt.Printf(">>>>>>\n")
			break
		case m <- struct{}{}:
		}
		//m <- struct{}{}
		fmt.Printf("res:(%v)\n", res.Great)
	}
	stream.CloseSend()
	cancel()
	time.Sleep(time.Second * 1000)
	//cancel()
	//}()
	//for {
	//	res, err := stream.Recv()
	//	if err == io.EOF {
	//		fmt.Printf("recv eof(%v)\n", err)
	//		break
	//	}
	//	if err != nil {
	//		fmt.Printf("recv error(%v)\n", err)
	//		return
	//	}
	//	fmt.Printf("res:(%v)\n", res.Great)
	//}
	//<-end
}

func mainDemo() {

	ctx, cancel := context.WithCancel(context.Background())

	h := make(chan struct{}, 1)
	go check(cancel, ctx, h)

	go func() {
		for {
			//write
			//read

			select {
			case <-ctx.Done():
				return
			case h <- struct{}{}:
			}
		}
	}()

	time.Sleep(time.Second * 100000)
	cancel() //正常关闭需要调用cancel
}
