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
func main() {
	clnt := client{addr: "192.168.6.95:7777"}
	conn, err := grpc.DialContext(context.Background(), clnt.addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}
	defer conn.Close()
	gtcli := streamtest.NewGTestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	end := make(chan struct{})
	stream, err := gtcli.ExampleDealStream(ctx)
	if err != nil {
		fmt.Printf("client to server(%v)\n", err)
	}
	go func() {
		defer close(end)
		for num := 0; num < 1000; num++ {
			in := streamtest.InMsg{Name: fmt.Sprintf("hudi_%v", num+1)}
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
