package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"myself/proto/streamtest"
	"net"
)

type server struct {
	sa string
}

func (s *server) ExampleDeal(in *streamtest.InMsg, steam streamtest.GTest_ExampleDealServer) error {
	for i := 0; i < 10; i++ {
		err := steam.Send(&streamtest.OutMsg{Great: fmt.Sprintf("hello %v %v", in.Name, i)})
		if err != nil {
			fmt.Printf("server stream error(%v)\n", err)
			return err
		}
	}
	return nil
}

func (s *server) ExampleDealStreamTemp(streamServer streamtest.GTest_ExampleDealStreamServer) error {
	outChan := make(chan streamtest.OutMsg, 10)
	go func() {
		num := 0
		defer close(outChan)
		for {
			in, err := streamServer.Recv()
			num++
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("error(%v)\n", err)
				return
			}
			out := streamtest.OutMsg{
				Great: fmt.Sprintf("greating hello %v %v", in.Name, num),
			}
			outChan <- out
		}
	}()
	for {
		out, ok := <-outChan
		if !ok {
			break
		}
		err := streamServer.Send(&out)
		if err != nil {
			fmt.Printf("error(%v)\n", err)
			return err
		}
	}
	return nil
}
func (s *server) ExampleDealStream(streamServer streamtest.GTest_ExampleDealStreamServer) error {
	//outChan := make(chan streamtest.OutMsg, 10)
	defer fmt.Printf("===========steam end\n")
	num := 0
	//defer close(outChan)
	for {
		in, err := streamServer.Recv()
		num++
		if err == io.EOF {
			fmt.Printf("===eof(%v)\n", err)
			break
		}
		if err != nil {
			fmt.Printf("error(%v)\n", err)
			return err
		}
		out := streamtest.OutMsg{
			Great: fmt.Sprintf("greating hello %v, %v", in.Name, num),
		}
		if in.Name == "eof" {
			out = streamtest.OutMsg{
				Great: fmt.Sprintf("last packet %v, %v", in.Name, num),
			}
		}
		//if num == 3 {
		//	time.Sleep(time.Second * 2000)
		//}
		err = streamServer.Send(&out)
		if err != nil {
			fmt.Printf("===send: error(%v)\n", err)
			return err
		}
		//outChan <- out
	}

	//for {
	//	out, ok := <-outChan
	//	if !ok {
	//		break
	//	}
	//	err := streamServer.Send(&out)
	//	if err != nil {
	//		fmt.Printf("error(%v)\n", err)
	//		return err
	//	}
	//}

	return nil
}
func main() {
	ser := server{sa: "192.168.5.25:7777"}
	lis, err := net.Listen("tcp", ser.sa)
	if err != nil {
		fmt.Printf("tcp listen error(%v)\n", err)
		return
	}
	grpS := grpc.NewServer()
	streamtest.RegisterGTestServer(grpS, &ser)
	reflection.Register(grpS)
	fmt.Printf("start grpctest(%v) ...\n", ser.sa)
	if err := grpS.Serve(lis); err != nil {
		fmt.Printf("serve error(%v)\n", err)
		return
	}
}
