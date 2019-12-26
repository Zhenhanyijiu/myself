package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type mp3Server struct {
	addr string
	path string
	//fname string
}

func (m *mp3Server) HandleVoice(w http.ResponseWriter, r *http.Request) {

	//val := r.URL.Query()
	urlPath := r.URL.Path
	fmt.Printf("r.URL.Path:%v\n", urlPath)
	//path := val["path"]
	//fmt.Println("path:", path)
	//path0 := path[0]
	if len(urlPath) < 2 {
		fmt.Printf("error path is null\n")
		http.Error(w, "path is null", 404)
		return
	}
	file, err := os.Open(m.path + urlPath)
	if err != nil {
		http.Error(w, err.Error(), 404)
		fmt.Printf("open error (%v)\n", err)
		return
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 404)
		fmt.Printf("io read error (%v)\n", err)
		return
	}
	w.Write(buf)
}
func main() {
	addr := flag.String("sa", "192.168.5.25:8083", "opt,lesten addr,default(192.168.5.25:8083)")
	path := flag.String("path", "", "must,resource path")
	flag.Parse()
	if *path == "" {
		flag.Usage()
		return
	}
	fmt.Printf("=====listen(%v)\n", *addr)
	m := mp3Server{path: *path}
	m.addr = *addr
	http.HandleFunc("/", m.HandleVoice)
	ser := &http.Server{
		Addr:    m.addr,
		Handler: nil,
	}
	err := ser.ListenAndServe()
	if err != nil {
		fmt.Printf("start http server failed:%v\n", err)
	}
}
