package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type mp3Server struct {
	addr  string
	fname string
}

func (m *mp3Server) HandleMp3(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(m.fname)
	if err != nil {
		http.Error(w, err.Error(), 200)
		return
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 200)
		return
	}
	//bugger := bytes.NewBuffer(buf)
	w.Header().Set("Content-Type", "audio/mpeg")
	rd := bytes.NewReader(buf)
	http.ServeContent(w, r, "", time.Now(), rd)

}
func main() {
	fname := flag.String("fn", "", "must,file name")
	addr := flag.String("sa", "192.168.5.25:8083", "opt,lesten addr")
	flag.Parse()
	if *fname == "" {
		flag.Usage()
		return
	}
	m := mp3Server{}
	m.fname = *fname
	m.addr = *addr
	http.HandleFunc("/WebAudio-1.0-SNAPSHOT/audio/play/", m.HandleMp3)
	ser := &http.Server{

		Addr:    m.addr,
		Handler: nil,
	}
	err := ser.ListenAndServe()
	if err != nil {
		fmt.Printf("start http server failed:%v\n", err)
	}
}
