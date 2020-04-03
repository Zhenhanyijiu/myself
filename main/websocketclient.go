package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/guonaihong/clop"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wsurl = "ws://192.168.6.95:7700/ws/test?appkey=zpbfrnoef722gn6q435beiebqm2cgxyhcofwi3il"
var wss = "wss://demo-edu.hivoice.cn:1906/ws/mark/en?appkey=zpbfrnoef722gn6q435beiebqm2cgxyhcofwi3il"

//gurl ws -I -R pcm.lst.l -skey "voice=rf.col.0" "|"  -ac 1 -r -w -merge -H "rate:16000"
// -H "appkey:xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx" -H "closeVad:true"  -H "session-id:`uuidgen`"
// -H "am-speak:" -H "vocab-enhance:7" -H "false-alarm:3" -binary -p "@{voice}" eof -H "eof:eof"
// -send-rate "32000/250ms" -url ws://192.168.5.25:8101/ws/asr/pcm -output "/dev/null" "|"  -O  -wkey "err,last_body,voice"

func readListFile(fname string) ([]string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	out, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("out:\n%v\n,bytes:%v\n", string(out), out)
	fileNames := strings.Split(string(out), "\n")
	return fileNames, nil
}
func readAudioFile(fn string) (*bytes.Reader, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	voice, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(voice), nil
}

type args struct {
	Audio    string `clop:"-a; --audio" usage:"voice file name" default:""`
	ListName string `clop:"-l;--list" usage:"test file of voice list" default:""`
	//Rate  string `clop:"-rt;--rate" usage:"sample rate" default:"16000"` //("rate", "16000", "(opt)sample rate")
	//	Header string `clop:"-H;--header" usage:"header parameter," default:
	//"closeVad:true;rate:16000;appkey:xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx;"`
	Header     []string `clop:"-H;--header" usage:"header parameter"`
	UrlWs      string   `clop:"-u;--wsurl" usage:"websocket url" default:"ws://192.168.5.25:8101/ws/asr/pcm"`
	Debug      string   `clop:"-d;--debug" usage:"test file of voice list" default:""`
	RoutineNum int      `clop:"-n;--threadnum" usage:"the number of thead" default:"1"`
}

func main() {
	ags := args{}
	if err := clop.Bind(&ags); err != nil {
		fmt.Printf("clop error(%v)\n", err)
		return
	}
	if ags.Audio == "" && ags.ListName == "" {
		fmt.Printf("-aud, -lst all null\n")
		return
	}
	//head := http.Header{}
	heads := http.Header{}
	//heads.Set("rate", "16000")
	//heads.Set("appkey", "xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx")
	//heads.Set("closeVad", "true")
	//heads.Set("session-id", "fe945929-dfa4-4960-8e31-f8c057ed2b99")
	//heads.Set("am-speak", "")
	////heads.Set("domains", "1,11")
	//heads.Set("eof", "eof")
	//heads.Set("vocab-enhance", "5")
	//heads.Set("group-id", "5e58e7ec39de65172c496471")
	//heads.Set("vocab-id", "5e707b7639de65172c496476")
	//-H "vocab-enhance:$VocabEnhance" -H "group-id:$gid" -H "vocab-id:$hotid"

	//if *language == "en" {
	//	head.Set("am-speak", "")
	//}
	//if *language == "cn" {
	//	head.Set("vocab-enhance", "7")
	//	head.Set("false-alarm", "3")
	//}
	dialDefault := websocket.DefaultDialer
	for _, v := range ags.Header {
		if len(v) != 0 {
			val := strings.Split(v, ":")
			if len(val) < 2 {
				continue
			}
			heads.Set(val[0], val[1])
		}
	}
	if len(ags.Audio) != 0 {
		rd, err := readAudioFile(ags.Audio)
		if err != nil {
			fmt.Printf("read audio file failed\n")
			return
		}
		c, _, err := dialDefault.Dial(ags.UrlWs, heads)
		if err != nil {
			fmt.Printf("Websocket Dial error(%v)\n", err)
			return
		}
		defer c.Close()
		//wg := sync.WaitGroup{}
		//wg.Add(1)
		go sendRoutine(rd, c)
		readRoutine(c, ags.Audio)
		return
	}

	//list file
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	voiceChan := make(chan string, 1000)
	var pc sync.WaitGroup
	pc.Add(1)
	go func(voiceCh chan string) {
		defer fmt.Printf("exit read list goroutine...\n")
		defer pc.Done()
		defer close(voiceChan)
		voiceSlice, err := readListFile(ags.ListName)
		if err != nil {
			fmt.Printf("read list voice file error(%v)\n", err)
			cancel()
			return
		}
		for _, v := range voiceSlice {
			if len(v) != 0 {
				select {
				case <-ctx.Done():
					return
				case voiceChan <- v:
				}
			}
		}
	}(voiceChan)

	pc.Add(1)
	defer pc.Wait()
	go func() {
		defer pc.Done()
		wg.Wait()
		cancel()
	}()
	wg.Add(ags.RoutineNum)
	for i := 0; i < ags.RoutineNum; i++ {
		go func(voiceCh chan string, id int) {
			fmt.Printf("==id(%v)\n", id)
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case voiceName, ok := <-voiceCh:
					if !ok {
						return
					}
					rd, err := readAudioFile(voiceName)
					if err != nil {
						fmt.Printf("read audio file failed\n")
						return
					}
					c, _, err := dialDefault.Dial(ags.UrlWs, heads)
					if err != nil {
						fmt.Printf("Websocket Dial error(%v)\n", err)
						return
					}
					defer c.Close()
					go sendRoutine(rd, c)
					readRoutine(c, voiceName)
				}
			}
		}(voiceChan, i+1)
	}
}

func sendRoutine(rd *bytes.Reader, c *websocket.Conn) {
	//defer wg.Done()
	buf := make([]byte, 320)
	for i := 0; ; i++ { //
		n, err := io.ReadFull(rd, buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		c.WriteMessage(websocket.TextMessage, buf[:n])
		time.Sleep(time.Millisecond * 10)
	}
	c.WriteMessage(websocket.TextMessage, []byte("eof"))
	c.WriteMessage(websocket.CloseNormalClosure, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func readRoutine(c *websocket.Conn, voice string) {
	//read ws
	puncResultTemp := ""
	var netResult NetResult
	var res Result
	i := 0
	for {
		start := time.Now()
		c.SetReadDeadline(start.Add(time.Second * 20))
		_, out, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("error(%v)\n", err)
			printJson(err.Error(), out, voice)
			return
		}
		//fmt.Printf("TestMsg(%v),msgType(%v)\n", websocket.TextMessage, msgType)
		if i < 30 {
			i++
			//fmt.Printf("===out(%v)\n\n", string(out))
		}
		if err := json.Unmarshal(out, &netResult); err != nil {
			fmt.Printf("1 Unmarshal error(%v)\n", err)
			printJson(err.Error(), out, voice)
			return
		}
		if netResult.ErrCode != 0 {
			//fmt.Printf("Error Packet:%v\n", string(out))
			printJson(err.Error(), out, voice)
			return
		}
		//
		if err := json.Unmarshal(netResult.CurrResult.Text, &res); err != nil {
			fmt.Printf("2 Unmarshal error(%v)\n", err)
			printJson(err.Error(), out, voice)
			return
		}
		if res.Punc_text != "" {
			fmt.Printf("###fixed, punctext:(%v)\n", res.Punc_text)
			puncResultTemp += res.Punc_text
		}
		if netResult.AllResult == nil {
			continue
		}
		//fmt.Printf("%v", string(out))
		if err := json.Unmarshal(netResult.AllResult[0].Text, &res); err != nil {
			fmt.Printf("3 Unmarshal error(%v)\n", err)
			printJson(err.Error(), out, voice)
			break
		}
		printJson("", out, voice)
		puncResultTemp += res.Punc_text
		break
	}
	if string(puncResultTemp) != string(res.PuncResult) {
		fmt.Printf("===error,punc result not equal\n")
		printJson("===error,punc result not equal", nil, voice)
		panic("not equal...")
		return
	}
	fmt.Printf("==punc_result equal\n")
}

type NetResult struct {
	ErrCode    int         `json:"errcode"`
	ErrMsg     string      `json:"errmsg"`
	IsEnd      bool        `json:"isEnd"`
	CurrResult AsrResult   `json:"currResult"`
	AllResult  []AsrResult `json:"allResult"`
}
type AsrResult struct {
	Text json.RawMessage `json:"text"`
	AsrResultAux
}
type AsrResultAux struct {
	StartTime  int64  `json:"startTime"`
	AudioSTime int32  `json:"audioSTime"`
	AudioETime int32  `json:"audioETime"`
	EndTime    int64  `json:"endTime"`
	Sid        string `json:"sid"`
	Area       string `json:"area"`
	Time       string `json:"time"`
	Number     int    `json:"number"`
}
type Result struct {
	CUR_active               int32
	CUR_domain               int32
	CUR_path_time            float64
	CUR_sil_hangout          float64
	CUR_speech_hangout       float64
	CUR_state                string
	CUR_time                 float64
	CUR_residual_length      float32
	CUR_residual_sil_prob    float32
	CUR_vad_state            string
	SET_acoustic             int32
	SET_domains_num          int32
	SET_post_proc            bool
	SET_sample_rate          int32
	SET_scene                int
	SET_scene_info           string
	SET_signal               int32
	Text                     string `json:"text"`
	Type                     string `json:"type"`
	Var_text                 string `json:"var_text"`
	Status                   string `json:"status"`
	Result                   string `json:"result"`
	CM_speech                float32
	CM_silence               float32
	CM_sentence              float32
	CM_final                 int
	CUR_residual_speech_prob float32
	Vcoab_enhance            float32 `json:"vcoab_enhance"`
	Total_audio_sec          float32
	Total_cache_sec          float32
	Punc_text                string `json:"punc_text"`
	PuncResult               string `json:"punc_result"` //punctuation data
}

type EndResult struct {
	Err      string          `json:"err"`
	LastBody json.RawMessage `json:"last_body"`
	Voice    string          `json:"voice"`
}

func printJson(err string, lastBody []byte, voice string) {
	end := EndResult{
		Err:      err,
		LastBody: lastBody,
		Voice:    voice,
	}
	out, _ := json.Marshal(&end)
	fmt.Printf("%v", string(out))
}
