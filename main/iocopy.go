package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Read11 struct {
	buf   bytes.Buffer
	chBuf chan []byte
	slice []byte
}

func (rd *Read11) Read(p []byte) (int, error) {
	//tmp := make([]byte, 3)
	//n, err := io.ReadFull(&rd.buf, tmp)
	//time.Sleep(time.Millisecond * 500)
	//if n > 0 {
	//	copy(p, tmp[:n])
	//	return n, nil
	//}
	//return n, err
	buf, ok := <-rd.chBuf
	if !ok {
		return 0, io.EOF
	}
	copy(p, buf)
	return len(buf), nil
}

type Write11 struct {
	buf   bytes.Buffer
	chBuf chan []byte
	slice []byte
}

func (w *Write11) Write(p []byte) (int, error) {

	res := fmt.Sprintf("<%v>", strings.ToUpper(string(p)))
	fmt.Printf("===%v\n", res)
	_, err := w.buf.WriteString(res)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
func main() {
	//New(os.Stdout).Msg("hello\n")

	//New(NewFilter(os.Stdout, true)).Msg("drop hello\n")
	//New(NewFilter(os.Stdout, false)).Msg("accept hello\n")
	//var errcode int64 = 10
	//atomic.CompareAndSwapInt64(&errcode, 0, 1)
	//atomic.CompareAndSwapInt64(&errcode, 10, 64)
	//fmt.Printf("===%v\n", errcode)
	//var addr int32 = 16
	//newI := atomic.AddInt32(&addr)
	//fmt.Printf("newI:%d,%d\n", newI, addr)
	var rd Read11
	var w Write11
	_, err := rd.buf.WriteString("abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		fmt.Printf("WriteString error(%v)\n", err)
		return
	}
	chBuf := make(chan []byte, 100)
	rd.chBuf = chBuf
	w.chBuf = chBuf
	//
	go func() {
		num := 10
		defer close(chBuf)
		for i := 0; i < num; i++ {
			buf := []byte(fmt.Sprintf("hello %d", i+1))
			chBuf <- buf
			time.Sleep(time.Millisecond * 500)
		}
	}()
	n, err := io.Copy(&w, &rd)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	fmt.Printf("n:%v\n", n)
	fmt.Printf("w.buf(%v)\n", w.buf.String())

	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	//fi := os.FileInfo{}
	fi, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}
	for i, v := range fi {
		fmt.Printf("%v\n", v.Size())
		fmt.Printf("===index:%v,%v,%v\n", i, v.Name(), v.IsDir())
	}
}
