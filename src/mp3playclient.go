package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//func main() {
//	txt := flag.String("input", "", "must")
//	flag.Parse()
//	if *txt == "" {
//		flag.Usage()
//		return
//	}
//	resp, err := http.Get("http://192.168.5.25:7999/tts?appkey=xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx&format=wav&person=xiaofeng&type=chinese&text=" + *txt)
//	if err != nil {
//		fmt.Println("get failed!")
//		return
//	}
//	defer resp.Body.Close()
//	fmt.Println(resp.Header.Get("Content-Type"))
//	fmt.Println(resp.Header.Get("Content-Length"))
//	p := []byte{}
//	fmt.Println(len(p))
//}

//func main() {
//
//	pipereader, pipewriter := io.Pipe()
//	end := make(chan struct{}, 1)
//	go Pipewrite(pipewriter)
//	go PipeRead(pipereader, end)
//	//time.Sleep(1 * time.Second)
//	<-end
//}
//
//func PipeRead(reader *io.PipeReader, end chan struct{}) {
//	buf := make([]byte, 128)
//	defer func() {
//		close(end)
//	}()
//	for {
//		//fmt.Println("读出端开始阻塞五秒...")
//		//time.Sleep(5 * time.Second)
//		//fmt.Println("接收端开始接收")
//		n, err := reader.Read(buf)
//		if err != nil {
//			//fmt.Println("=======", err)
//			return
//		}
//		fmt.Printf("收到字节: %d\n buf内容: %s\n", n, buf[:n])
//	}
//}
//
//func Pipewrite(write *io.PipeWriter) {
//	data := []byte("Go语言小实践")
//	for i := 0; i < 12; i++ { //写入次数 i
//		n, err := write.Write(data)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Printf("写入字节 %d\n", n)
//	}
//
//	write.Close()
//}
func (p *playServer) Play(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stat play........", r.URL.Path)
	if r.Method == "GET" {
		urlValue, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			return
		}
		txt := urlValue.Get("text")
		fmt.Println("text:", txt)
		f, err := os.Open(*p.fileName)
		if err != nil {
			fmt.Println("open file failed")
			return
		}
		content, err := ioutil.ReadAll(f)
		fmt.Println(">>>len:", len(content))
		if err != nil {
			fmt.Println("ReadAll file failed")
			return
		}
		fmt.Println("====1", w.Header().Get("Content-Type"))
		w.Header().Add("Content-Type", "audio/mpeg") //audio/rn-mpeg //audio/mpeg
		w.Header().Set("Cache-Control", "max-age=3036000")
		///////
		//w.Header().Add("Content-Type", "audio/wav")
		//w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
		fmt.Println("====2", w.Header().Get("Content-Type"))

		w.Write(content)
		return
	}
	return
}

type playServer struct {
	fileName *string
}

func main() {
	//input := flag.String("fn", "", "must")
	//sa := flag.String("sa", "192.168.5.25:7986", "must")
	//flag.Parse()
	//if *input == "" {
	//	flag.Usage()
	//	return
	//}
	ur := "http://192.168.5.25:7986/tts"
	//rsp, err := http.Get(url)
	cli := http.Client{}
	rsp, err := cli.Get(ur)
	if err != nil {
		fmt.Println(">>>>error:", err)
		return
	}
	hc := rsp.Header.Get("Content-Type")
	fmt.Println("Content-Type:", hc)
	ln := rsp.ContentLength
	fmt.Println("len:", ln)
	var buf bytes.Buffer
	num, err := io.Copy(&buf, rsp.Body)
	if err != nil {
		fmt.Println("copy error:", err)
		return
	}
	fmt.Println("num:", num)
	f, err := os.Create("tst.mp3")
	if err != nil {
		fmt.Println("Create file error:", err)
		return
	}
	n, err := f.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Write file error:", err)
		return
	}
	fmt.Println("...n:", n)
	//ps := playServer{}
	//ps.fileName = input
	//http.HandleFunc("/tts/", ps.Play)
	//fmt.Println(http.ListenAndServe(*sa, nil))
}
