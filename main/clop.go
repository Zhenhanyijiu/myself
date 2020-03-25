package main

import (
	"bytes"
	"fmt"
	"github.com/guonaihong/clop"
	"io"
	"net/http"
	"strings"
	"time"
)

//type exam struct {
//	//Header []string `clop:"H=files"`
//	Fname string `clop:"-u; --filename" usage:"file name" valid:"required"`
//}
type args struct {
	Audio    string `clop:"-a; --audio" usage:"voice file name" default:""`
	ListName string `clop:"-l;--list" usage:"test file of voice list" default:""`
	//Rate  string `clop:"-rt;--rate" usage:"sample rate" default:"16000"` //("rate", "16000", "(opt)sample rate")
	//	Header string `clop:"-H;--header" usage:"header parameter," default:
	//"closeVad:true;rate:16000;appkey:xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx;"`
	Header []string `clop:"-H;--header" usage:"header parameter"`
	UrlWs  string   `clop:"-u;--wsurl" usage:"websocket url" default:"ws://192.168.5.25:8101/ws/asr/pcm"`
	Debug  string   `clop:"-d;--debug" usage:"test file of voice list" default:""`
}

func main112() {
	ar := args{}
	err := clop.Bind(&ar)
	if err != nil {
		fmt.Printf("clop, error(%v)\n", err)
		return
	}
	fmt.Printf("exam:%v,%v,%v\n", ar, ar.Header[0], ar.Header[1])
	heads := http.Header{}
	heads.Set("rate", "16000")
	heads.Set("rate", "33")
	heads.Set("rate", "33fff")
	//out, _ := json.Marshal(&heads)
	hh := heads.Get("rate")
	fmt.Printf("out,%v,%v\n", hh, heads["Rate"][0])
}

type zerolog struct {
	w io.Writer
}

func New(w io.Writer) *zerolog {
	return &zerolog{w: w}
}

func (z *zerolog) Msg(s string) {
	z.w.Write([]byte(s))
}

type filter struct {
	w    io.Writer
	drop bool
}

func NewFilter(w io.Writer, drop bool) *filter {
	if drop {
		return &filter{w: w, drop: drop}
	}

	return &filter{w: w}
}

func (f *filter) Write(p []byte) (n int, err error) {
	if f.drop {
		return 0, nil
	}

	return f.w.Write(p)
}

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
}
