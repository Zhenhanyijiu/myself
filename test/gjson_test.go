package test

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/core"
	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"
)
import "github.com/tidwall/gjson"

var jsontext = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37.9,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func TestGjson(t *testing.T) {
	res := gjson.Get(jsontext, "ageq")
	if res.Exists() {
		ff := res.Float()
		fmt.Printf("json:%v\n", ff)
	} else {
		fmt.Printf("not exist:%v\n", res.String())
	}

	res = gjson.Get(jsontext, "children")
	for i, v := range res.Array() {
		fmt.Printf("i:%v, v:%v\n", i, v)
	}
	//sjson.SetRaw()
	buf := bytes.NewBuffer([]byte("aaaa"))
	fmt.Printf("===%v\n", len(buf.Bytes()))
	ff := make([]byte, 77)
	n, _ := buf.Read(ff)
	fmt.Printf("n:%v,%v\n", n, len(buf.Bytes()))
}
func TestCron(t *testing.T) {
	//cr := cron.New(cron.WithSeconds())
	//cron.New(cron.WithParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))
	//cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	cr := cron.New()
	cr.AddFunc("1 0 1 2 ?", func() {
		fmt.Printf("=========\n")
	})
	cr.Start()
	defer cr.Stop()
	time.Sleep(time.Second * 6000)
}
func TestPrintq(t *testing.T) {
	//fmt.Printf("=======\n")
	//res := time.Now().Format("20060102")
	//fmt.Printf("=======%v\n", res)
	//var UploadData UploadData
	//out, _ := json.MarshalIndent(&UploadData, "", "  ")
	//fmt.Printf("out:\n%v\n", string(out))
	//voiceQuery := VoiceQuery{
	//	Query: []OneQuery{OneQuery{}},
	//}
	//out, _ = json.MarshalIndent(&voiceQuery, "", "  ")
	//fmt.Printf("out:\n%v\n", string(out))
	//
	//var rsp = DownResponse{Data: []DataInfo{DataInfo{}}}
	//out, _ = json.MarshalIndent(&rsp, "", "  ")
	//fmt.Printf("out:\n%v\n", string(out))

	///////////////
	freshTime, _ := time.Parse("15:04:05", "23:58:00")
	fmt.Printf("===>curtime:%v\n", freshTime.String())
	sCurTime := time.Now().Format("15:04:05")
	CurTime, _ := time.Parse("15:04:05", sCurTime)
	fmt.Printf("===curtime:%v\n", CurTime.String())
	outL := freshTime.Sub(CurTime).Hours()
	fmt.Printf("===%v\n", outL)
	//aa := time.Duration(time.Now().Day()*24+23) * time.Hour
	//bb := aa - time.Duration(time.Now().Hour())
	//t1 := time.NewTimer(bb)
	t1 := time.Hour * 24
	fmt.Printf("===??:%v\n", t1)

	//t1 := time.NewTimer(aa)
	//t2 := time.Since(t1)
	//time.Now().Sub(time.Time(t2))
	//time.NewTimer()
	//time.Now()()
	var num int32 = 9
	atomic.AddInt32(&num, 1)
	atomic.AddInt32(&num, 1)
	atomic.AddInt32(&num, -1)
	atomic.AddInt32(&num, -1)
	fmt.Printf("===%v\n", atomic.LoadInt32(&num))

}

type DataInfo struct {
	Result string
	Voice  []byte
}
type DownResponse struct {
	ErrMsg  string
	ErrCode int
	Only    string
	Data    []DataInfo
}

type VoiceQuery struct {
	Query []OneQuery
}
type OneQuery struct {
	Create_time uint64 "create_time"
	Session_id  string "session_id"
}

type LogData struct {
	Ip            string "ip"
	Session_id    string "session_id"
	Server_addr   string "server_addr"
	Server_serial string "server_serial"
	Server_type   string "server_type"
	Time_stamp    string "time_stamp"
	Opt_timeout   int    "opt_timeout"
	Opt_resformat string "opt_resformat"
	Opt_imei      string "opt_imei"
	Opt_key       string "opt_key"
	Opt_secret    string "opt_secret"
	Task_type     string "task_type"
	Oral_text     string "oral_text"
	Result        string "result"
	Pcm_len       uint   "pcm_len"
	Recog_time    int    "recog_time"
	Recv_time     int    "recv_time"
	Error_code    int    "error_code"
	Rectime       int    "rectime"
	Attr          string "attr"
	Create_time   uint64 "create_time"
	FileID        string "file_id"
}

type VoiceData struct {
	Result       string "result"
	Session_id   string "session_id"
	Create_time  uint64 "create_time"
	EncodeFormat string "encodeFormat"
	Session_mark string "session_mark"
	Voice        []byte "voice"
}

type UploadData struct {
	LogInfo   LogData
	VoiceInfo VoiceData
}

func TestTime(t *testing.T) {
	n := time.Second * 5
	timeout := make(chan bool, 1)
	time.AfterFunc(n, func() {
		timeout <- true
	})
	<-timeout
	fmt.Printf("timeout=======\n")
}
func TestD(t *testing.T) {
	fmt.Printf("====")
}

func TestMarker(t *testing.T) {
	//ps := testNewProxy(t)
	//router := func() *gin.Engine {
	//	router := gin.Default()
	//	router.POST("/n2t/en", func(c *gin.Context) {
	//		ps.ResolveEn(c)
	//	})
	//	return router
	//}()
	//ts := httptest.NewServer(http.HandlerFunc(router.ServeHTTP))
	//defer ts.Close()
	data := make(url.Values)
	//data["text"] = []string{`{"Version":1,"DisplayText":"pack","Markers":[{"Type":"phone","Position":{"Start":0,"Length":4},"Value":["ˈp·æ·k"]}]}`} //prəˌnʌnsiˈeɪʃn
	///data["text"] = []string{`a[ps:æ] a[ps:æ] apple`}
	//data["text"] = []string{`pronunciation[p:p r ə ˌn ʌn s i ˈeɪ ʃn]`} //prəˌnʌnsiˈeɪʃn
	data["text"] = []string{"do business[g:], with the city."}
	data["text"] = []string{"sorry to bother you[g:],     but could you take my picture, please?"}
	res, err := http.PostForm("http://192.168.5.25:8080/n2t/en", data)
	if err != nil {
		t.Errorf("===error(%v)\n", err)
		return
	}
	defer res.Body.Close()
	outRes, _ := ioutil.ReadAll(res.Body)
	t.Logf("res:%v\n", string(outRes))
	//fmt.Printf("====n2t return:%v\n", bodyParse)
	var tstMark TestMark
	//json.Unmarshal()
	err = json.Unmarshal(outRes, &tstMark)
	if err != nil {
		fmt.Printf("====err:%v\n", err)
		return
	}
	txt := tstMark.UserText
	for _, v := range tstMark.Markers {
		if v.Texts[0] != txt[v.Start:v.Start+v.Length] {
			fmt.Printf("===error not equal...\n")
			break
		}
	}
	fmt.Printf("===test ok....\n")
}

type TestMark struct {
	UserText string `json:"usertext"`
	Markers  []Mark `json:"markers"`
}
type Mark struct {
	Type   string   `json:"type"`
	Start  int      `json:"start"`
	Length int      `json:"length"`
	Texts  []string `json:"texts"`
}

/*
func TestSync(t *testing.T) {
	//m := syncmap.Map{}
	w := semaphore.NewWeighted(1)

}

func TestTempFunction(t *testing.T) {
	my, err := NewMysql(dbnameTest, dburlTest)
	assert.NoError(t, err)
	defer my.mysqlClient.Close()

	dbC := my.mysqlClient.Begin()
	for i := 0; i < 10; i++ {
		nowTime := time.Now().UnixNano()
		asrlog := &AsrLogDay{
			DateString: time.Unix(0, nowTime).Format("20060102"),
			SessionID:  "3333333",
			OptKey:     "111",
			Result:     `{"CUR_active":1,"CUR_domain":1,"CUR_path_time":8.82,"CUR_sil_hangout":0.39,"CUR_speech_hangout":0,"CUR_state":"normal","CUR_time":9.24,"CUR_residual_length":0.27,"CUR_residual_sil_prob":0,"CUR_vad_state":"speech_end:9.24","SET_acoustic":1,"SET_domains_num":1,"SET_post_proc":false,"SET_sample_rate":16,"SET_scene":1,"SET_scene_info":"","SET_signal":0,"text":"通过对车辆数据的深度挖掘，实现多种以公安应用为核心的业务功能，通过与公安部门车要登记。","type":"fixed","status":"const","result":"通过对车辆数据的深度挖掘，实现多种以公安应用为核心的业务功能，通过与公安部门车要登记。","CM_speech":0,"CM_silence":0,"CM_sentence":0,"CM_final":0,"CUR_residual_speech_prob":0,"vocab_enhance":0,"Total_audio_sec":9.21,"Total_cache_sec":-0.03,"punc_text":"","punc_result":"","word_info":[]}`,
			PcmLen:     377,
			ServerType: "asrWebsocket",
			CreateTime: nowTime,
		}
		err = dbC.Table(asrlog.TableName()).Create(asrlog).Error
		assert.NoError(t, err)
		time.Sleep(time.Second * 2)
	}

	//nowTime = time.Now().UnixNano()
	////dbC := my.mysqlClient.Begin()
	//asrlog = &AsrLogDay{
	//	DateString: time.Unix(0, nowTime).Format("20060102"),
	//	SessionID:  "3333333",
	//	OptKey:     "111",
	//	Result:     `{"CUR_active":1,"CUR_domain":1,"CUR_path_time":8.82,"CUR_sil_hangout":0.39,"CUR_speech_hangout":0,"CUR_state":"normal","CUR_time":9.24,"CUR_residual_length":0.27,"CUR_residual_sil_prob":0,"CUR_vad_state":"speech_end:9.24","SET_acoustic":1,"SET_domains_num":1,"SET_post_proc":false,"SET_sample_rate":16,"SET_scene":1,"SET_scene_info":"","SET_signal":0,"text":"通过对车辆数据的深度挖掘，实现多种以公安应用为核心的业务功能，通过与公安部门车要登记。","type":"fixed","status":"const","result":"通过对车辆数据的深度挖掘，实现多种以公安应用为核心的业务功能，通过与公安部门车要登记。","CM_speech":0,"CM_silence":0,"CM_sentence":0,"CM_final":0,"CUR_residual_speech_prob":0,"vocab_enhance":0,"Total_audio_sec":9.21,"Total_cache_sec":-0.03,"punc_text":"","punc_result":"","word_info":[]}`,
	//	PcmLen:     377,
	//	ServerType: "asrWebsocket",
	//	CreateTime: nowTime,
	//}
	//err = dbC.Table(asrlog.TableName()).Create(asrlog).Error
	//assert.NoError(t, err)

	err = dbC.Commit().Error
	if err != nil {
		dbC.Rollback()
	}
	assert.NoError(t, err)
}

func TestSync(t *testing.T) {
	//wei := semaphore.NewWeighted(2)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//for i := 0; i < 7; i++ {
	//	wei.Acquire(ctx, 1)
	//	go func(index int) {
	//		defer wei.Release(1)
	//	}(i)
	//
	//}

	t.Parallel()
	ctx := context.Background() //Background()
	sem := semaphore.NewWeighted(3)
	fg := sem.Acquire(ctx, 1) == nil
	fmt.Printf("tries:%v\n", fg)
	fg = sem.Acquire(ctx, 1) == nil
	fmt.Printf("tries:%v\n", fg)
	fg = sem.Acquire(ctx, 1) == nil
	fmt.Printf("tries:%v\n", fg)
	sem.Release(1)
	fg = sem.Acquire(ctx, 1) == nil
	fmt.Printf("tries:%v\n", fg)
	//tryAcquire := func(n int64) bool {
	//	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	//	defer cancel()
	//	return sem.Acquire(ctx, n) == nil
	//}
	//
	//tries := []bool{}
	////sem.Acquire(ctx, 1)
	//tries = append(tries, tryAcquire(2))
	//tries = append(tries, tryAcquire(1))
	//
	//sem.Release(3)
	//
	//tries = append(tries, tryAcquire(1))
	////sem.Acquire(ctx, 1)
	//tries = append(tries, tryAcquire(1))

	//fmt.Printf("tries:%v\n", tries)
	//want := []bool{true, false, true, false}
	//for i := range tries {
	//	if tries[i] != want[i] {
	//		t.Errorf("tries[%d]: got %t, want %t", i, tries[i], want[i])
	//	}
	//}
}
*/
func TestPost(t *testing.T) {

	rsp, err := http.PostForm("http://192.168.5.25:9000/", url.Values{
		"name":    []string{"xiao", "aff"},
		"address": []string{"hhhhh"},
	})
	assert.NoError(t, err)
	rs, err := ioutil.ReadAll(rsp.Body)
	assert.NoError(t, err)
	fmt.Printf("%v\n", string(rs))
}

func TestInt16(t *testing.T) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, 0xabff)
	n := binary.LittleEndian.Uint16(buf)
	fmt.Printf("buf:%x, %x\n", buf, n)
	buf1 := []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0xaa, 0xbb, 0xcc, 0xdd}
	rd := bytes.NewReader(buf1)
	out1 := make([]int16, 2)
	for {
		if rd.Len() >= 4 {
			err := binary.Read(rd, binary.LittleEndian, out1)
			fmt.Printf("error:%v, rd.Size():%v,%v\n", err, rd.Size(), rd.Len())
			if err != nil {
				fmt.Printf("=======errror:%v\n", err)
				break
			}
			for i, v := range out1 {
				fmt.Printf("%d===%x\n", i, uint16(v))
			}
		} else {
			fmt.Printf("=========================\n")
			tmp := make([]int16, rd.Len()/2)
			err := binary.Read(rd, binary.LittleEndian, tmp)
			fmt.Printf("error:%v, rd.Size():%v,%v\n", err, rd.Size(), rd.Len())
			if err != nil {
				fmt.Printf("=======errror:%v\n", err)
				break
			}
			for i, v := range tmp {
				fmt.Printf("%d===%x\n", i, uint16(v))
			}
			break
		}

	}
	//out2 := make([]int16, 3)
	//err := binary.Read(rd, binary.LittleEndian, out2)
	//fmt.Printf("error:%v\n", err)
	////assert.NoError(t, err)
	//fmt.Printf("%x\n", out2)
}
func TestCopy(t *testing.T) {
	s1 := make([]byte, 2)
	s1[0] = 99
	s1[1] = 100

	s2 := []byte{0x11, 0x22, 0x33}
	copy(s1, s2)
	fmt.Printf("%v\n", s1)
}

type name string

func TestIoReadFull(t *testing.T) {
	s2 := []byte{0x11, 0x22, 0x33, 0x22, 0x33, 0x22, 0x33, 0x22, 0x33}
	//var s2 []byte
	buf := make([]byte, 4)
	reader := bytes.NewReader(s2)
	for {
		fmt.Printf("===reader len:%v\n", reader.Len())
		//if reader.Len() == 0 {
		//	break
		//}
		n, err := io.ReadFull(reader, buf)
		fmt.Printf("===n:%v, err:%v\n", n, err)
		break
	}
	var nn name
	nn = "asjdi"

	fmt.Printf("===%v\n", nn)
}

func TestDefer(t *testing.T) {

	num := 1
	numP := &num
	fmt.Printf("===before %v\n", numP)
	//defer fmt.Printf("===defer num:%v\n", numP)
	defer func() {
		fmt.Printf("===defer num:%v\n", numP)
	}()
	num2 := 88
	numP = &num2
	//num = 100
	fmt.Printf(">>>%v\n", numP)
}

//func TestShouldBindUri(t *testing.T) {
//	//DefaultWriter = os.Stdout
//	router := gin.New()
//
//	type Person struct {
//		Name string `uri:"name" binding:"required"`
//		Id   string `uri:"id" binding:"required"`
//	}
//	router.Handle("GET", "/rest/:name/:id", func(c *gin.Context) {
//		var person Person
//		assert.NoError(t, c.ShouldBindUri(&person))
//		assert.True(t, "" != person.Name)
//		assert.True(t, "" != person.Id)
//		c.String(http.StatusOK, "ShouldBindUri test OK")
//	})
//
//	//path, _ := exampleFromPath("/rest/:name/:id")
//	//w := performRequest(router, "GET", path)
//	//assert.Equal(t, "ShouldBindUri test OK", w.Body.String())
//	//assert.Equal(t, http.StatusOK, w.Code)
//}
func TestGinFormData(t *testing.T) {
	type Aduio struct {
		Text  string               `form:"text"`
		Voice multipart.FileHeader `form:"voice"`
	}
	rout := func() *gin.Engine {
		r := gin.Default()
		r.POST("/bind", func(c *gin.Context) {
			fmt.Printf("===enter ...\n")
			var audio Aduio
			assert.NoError(t, c.ShouldBind(&audio))
			fmt.Printf("===text:%v\n", audio.Text)
			f, err := audio.Voice.Open()
			if assert.NoError(t, err) == false {
				return
			}
			defer f.Close()
			buf := new(bytes.Buffer)
			io.Copy(buf, f)
			fmt.Printf("voice len:%v\n", buf.Len())
			c.JSON(200, gin.H{"errcode": 0})
		})
		return r
	}()
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	result := "" //:= make(map[string]interface{})
	gout.POST(s.URL + "/bind").SetForm(core.H{"text": "qqqq",
		"voice": core.FormFile("./165599_10.pcm")}).BindBody(&result).Do()
	fmt.Printf("result:%v\n", result)
}

func TestGin1(t *testing.T) {
	type Query struct {
		Text string `form:"text"`
		Num  string `form:"num"`
	}
	rout := func() *gin.Engine {
		r := gin.Default()
		r.POST("/query", func(c *gin.Context) {
			fmt.Printf("===enter ...\n")
			var query Query
			assert.NoError(t, c.ShouldBindQuery(&query))
			fmt.Printf("===text:%v\n", query.Text)
			assert.Equal(t, "www", query.Text)
			assert.Equal(t, "33", query.Num)
			c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
		})
		return r
	}()
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	defer s.Close()
	result := "" //:= make(map[string]interface{})
	err := gout.POST(s.URL + "/query?text=www&num=33").BindBody(&result).Do()
	assert.NoError(t, err)
	fmt.Printf("===result:%v,err:%v\n", result, err)
}
func TestGin2(t *testing.T) {
	type Head struct {
		Name string `header:"name"`
		Age  string `header:"age"`
	}
	rout := func() *gin.Engine {
		r := gin.Default()
		r.POST("/header/pp", func(c *gin.Context) {
			var h Head
			c.ShouldBindHeader(&h)
			fmt.Printf("===Name:%v,Age:%v\n", h.Name, h.Age)
			c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
		})
		return r
	}()
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	defer s.Close()
	result := ""
	err := gout.POST(s.URL + "/header/pp").SetHeader(&Head{Name: "rrr", Age: "77"}).BindBody(&result).Do()
	assert.NoError(t, err)
	fmt.Printf("===result:%v,err:%v\n", result, err)
}

func TestGin3(t *testing.T) {
	type Ur struct {
		Path1 string `uri:"bb"`
		Path2 string `uri:"cc"`
	}
	rout := func() *gin.Engine {
		r := gin.Default()
		r.POST("/:bb", func(c *gin.Context) {
			var u Ur
			err := c.ShouldBindUri(&u)
			assert.NoError(t, err)
			fmt.Printf("===>>>path1:%v, path2:%v\n", u.Path1, u.Path2)
			c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
		})
		return r
	}()
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	defer s.Close()
	//path, ps := exampleFromPath("/")
	//fmt.Printf("==========path:%v,params:%v\n", path, ps)
	result := ""
	err := gout.POST(s.URL + "/wav").BindBody(&result).Do()

	assert.NoError(t, err)
	fmt.Printf("===result:%v,err:%v\n", result, err)
}
func exampleFromPath(path string) (string, gin.Params) {
	output := new(bytes.Buffer)
	params := make(gin.Params, 0, 6)
	start := -1
	for i, c := range path {
		if c == ':' {
			start = i + 1
		}
		if start >= 0 {
			if c == '/' {
				value := fmt.Sprint(rand.Intn(100000))
				params = append(params, gin.Param{
					Key:   path[start:i],
					Value: value,
				})
				output.WriteString(value)
				output.WriteRune(c)
				start = -1
			}
		} else {
			output.WriteRune(c)
		}
	}
	if start >= 0 {
		value := fmt.Sprint(rand.Intn(100000))
		params = append(params, gin.Param{
			Key:   path[start:],
			Value: value,
		})
		output.WriteString(value)
	}

	return output.String(), params
}

func TestGin4(t *testing.T) {
	type Xml struct {
		T1V string `xml:"t1"`
		T2V string `xml:"t2"`
	}
	rout := func() *gin.Engine {
		r := gin.Default()
		return r
	}()
	rout.POST("/aa", func(c *gin.Context) {
		var xv Xml
		err := c.ShouldBindXML(&xv)
		c.Copy()
		assert.NoError(t, err)
		fmt.Printf("t1:%v,t2:%v\n", xv.T1V, xv.T2V)
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
	})
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	defer s.Close()
	result := ""
	err := gout.POST(s.URL + "/aa").SetXML(&Xml{T1V: "zzz", T2V: "vvv"}).BindBody(&result).Do()
	assert.NoError(t, err)
	fmt.Printf("===result:%v,err:%v\n", result, err)
}
func TestGin5(t *testing.T) {
	rout := func() *gin.Engine {
		r := gin.Default()
		return r
	}()
	rout.POST("/aa", func(c *gin.Context) {
		//c.DefaultPostForm()
		//c.ShouldBindYAML()
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
	})
}

func TestGin6(t *testing.T) {
	body := bytes.NewBufferString("foo=bar&page=11&both=&foo=second")
	req := httptest.NewRequest("POST", "/aa/bb", body)
	req.Header.Set("111", "aa")
	req.Header.Set("222", "bb")
	req.Header.Set("Content-Type", gin.MIMEPOSTForm)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req
	assert.Equal(t, "aa", c.GetHeader("111"))
	assert.Equal(t, "bb", c.GetHeader("222"))
	assert.Equal(t, "11", c.DefaultPostForm("page", "none"))
	assert.Equal(t, "bar", c.DefaultPostForm("foo", "none"))
	assert.Equal(t, "11", c.PostForm("page"))
	assert.Equal(t, "bar", c.PostForm("foo"))
	v, ok := c.GetPostForm("page")
	assert.True(t, ok)
	assert.Equal(t, "11", v)
	//assert.Equal(t, "bar", c.GetPostForm("foo"))
}
func TestGin7(t *testing.T) {
	rout := gin.Default()
	flagString := ""
	rout.Use(func(c *gin.Context) {
		flagString += "##"
	})
	rout.Use(func(c *gin.Context) {
		flagString += "22end"
	})
	fmt.Printf("1=======flagString:%v\n", flagString)
	rout.POST("/use", func(c *gin.Context) {
		fmt.Printf("===enter......\n")
		fmt.Printf("2=======flagString:%v\n", flagString)
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
	})
	s := httptest.NewServer(http.HandlerFunc(rout.ServeHTTP))
	defer s.Close()
	var result string
	err := gout.POST(s.URL + "/use").BindBody(&result).Do()
	fmt.Printf("3=======flagString:%v\n", flagString)
	assert.NoError(t, err)
	fmt.Printf("===result:%v,err:%v\n", result, err)
}
func TestTmp22(t *testing.T) {
	buf := []byte{1, 2}
	fmt.Printf("....buf:%v,%v\n", buf[:0], len(buf[:0]))
	m := map[int]string{
		0: "aa",
		1: "bb",
		2: "cc",
	}
	m[0] = "ahgfha"
	fmt.Printf("%v", m)
	for k, v := range m {
		fmt.Printf("%v,%v\n", k, v)
	}
	bf := make([]byte, 0, 100)
	bf2 := bytes.NewBuffer(bf)
	bf2.Write([]byte("agfhah"))
	bf2.Write(nil)
	fmt.Printf("%v,%v", bf2.Bytes(), bf2.String())
}
