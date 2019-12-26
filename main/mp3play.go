package main

import (
	"flag"
	"fmt"
	//"io/ioutil"

	"time"

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
	fmt.Println("start play........>", r.URL.Path)
	ran := r.Header.Get("Range")
	fmt.Println("........range:", ran)
	if len(ran) != 0 {
		fmt.Println("......ran[6:len-1]:", ran[6:len(ran)-1])
	}
	if r.Method == "GET" {
		urlValue, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			return
		}
		fmt.Println("r.URL.RawQuery>>", r.URL.RawQuery)
		txt := urlValue.Get("text")
		fmt.Println("text:", txt)
		mp3f, err := os.Open(*p.fileName)
		if err != nil {
			fmt.Println("open file failed")
			return
		}
		defer mp3f.Close()
		//content, err := ioutil.ReadAll(mp3f)
		//re := bytes.NewReader(content)
		w.Header().Set("Content-Type", "audio/mpeg")
		http.ServeContent(w, r, "", time.Now(), mp3f)
		//http.ServeFile(w, r, *p.fileName)
		//mp3f.Seek()
		//if len(ran) == 0 {
		//	content, err := ioutil.ReadAll(mp3f)
		//	fmt.Println(">>>len:", len(content))
		//	if err != nil {
		//		fmt.Println("ReadAll file failed")
		//		return
		//	}
		//	fmt.Println("====1", w.Header().Get("Content-Type"))
		//	w.Header().Add("Content-Type", "audio/mpeg") //audio/rn-mpeg //audio/mpeg
		//	w.Header().Set("Cache-Control", "max-age=30360000")
		//	//w.Header().Set("Connection", "close")
		//	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
		//	w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", 0, len(content)-1, len(content)))
		//	//identity
		//	//w.Header().Add("Transfer-Encoding", "identity")
		//	fmt.Println("====2", w.Header().Get("Content-Type"))
		//	//num, _ := strconv.Atoi(ran[6 : len(ran)-1])
		//	n, _ := w.Write(content)
		//	fmt.Println("n=", n, len(content))
		//	return
		//}

		//content, err := ioutil.ReadAll(mp3f)
		//fmt.Println(">>>len:", len(content))
		//if err != nil {
		//	fmt.Println("ReadAll file failed")
		//	return
		//}
		//fmt.Println("====1", w.Header().Get("Content-Type"))
		//w.Header().Add("Content-Type", "audio/mpeg") //audio/rn-mpeg //audio/mpeg
		//w.Header().Set("Cache-Control", "max-age=30360000")
		////w.Header().Set("Connection", "close")
		//
		////identity
		////w.Header().Add("Transfer-Encoding", "identity")
		//fmt.Println("====2", w.Header().Get("Content-Type"))
		//num, _ := strconv.Atoi(ran[6 : len(ran)-1])
		//w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)-num))
		//w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", num, len(content)-1, len(content)))
		//n, _ := w.Write(content[num:])
		//fmt.Println("n=", n, len(content))
		/////////
		////w.Header().Add("Content-Type", "audio/wav")
		////w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
		//fmt.Println("====2", w.Header().Get("Content-Type"))

		//w.Write(content)
		return
	}
	return
}

type playServer struct {
	fileName *string
}

func main() {
	input := flag.String("fn", "", "must")
	sa := flag.String("sa", "192.168.5.25:7981", "must")
	flag.Parse()
	if *input == "" {
		flag.Usage()
		return
	}
	ps := playServer{}
	ps.fileName = input
	http.HandleFunc("/tts", ps.Play)
	fmt.Println(http.ListenAndServe(*sa, nil))
}
