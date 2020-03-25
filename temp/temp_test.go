package temp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/guonaihong/gout"
	"gopkg.in/go-playground/validator.v8"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

//带上测试源文件
//https://github.com/DaveGamble/cJSON
//env GOPATH="/Users/ricky/go/src/n2t" go test -v -test.run Test_StartEn
//env GOPATH=/D_QAwYqH3be/yangyanan/pigai/pigai:/D_QAwYqH3be/yangyanan/gopath/ go test -v -test.run TestHandleInput api_test.go api.go
//env GOPATH=/D_QAwYqH3be/yangyanan/pigai/pigai:/D_QAwYqH3be/yangyanan/gopath/ go test -v -test.run TestMarkServer_Paragraph api_test.go api.go
//go test -v -bench=. benchmark_test.go
//env GOPATH=/D_QAwYqH3be/yangyanan/pigai/pigai:/D_QAwYqH3be/yangyanan/gopath/ go test -v -test.run=TestMarkServer_Paragraph .//当前文件夹
//env GOPATH=/D_QAwYqH3be/yangyanan/pigai/pigai:/D_QAwYqH3be/yangyanan/gopath/ go test -test.bench=BenchmarkHandlePerParagrah -benchmem api_test.go api.go
//env GOPATH=/D_QAwYqH3be/yangyanan/pigai/pigai:/D_QAwYqH3be/yangyanan/gopath/ go test -test.bench=BenchmarkHandlePerParagrah -benchmem .
//go build -o libinter.so -buildmode=c-shared inter.go
//env GO111MODULE=off GOPATH=/D_QAwYqH3be/yangyanan/asr_http/asr_http:/D_QAwYqH3be/yangyanan/gopath/ go test -v -test.run=TestBeforePuncResumeSplitText .

var bys = []byte{11, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 77, 12, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 34, 12, 23, 21, 34}
var ft = []float32{33.77, 33.77, 33.77, 33.77, 33.77, 33.77, 33.77, 33.77, 33.77, 33.77}

func TestByteToFloat32(t *testing.T) {
	f, _ := ByteToFloat32(bys)
	t.Log("===", f)
}
func BenchmarkByteToFloat32(b *testing.B) {
	//b.ResetTimer()
	//b.ReportAllocs()
	//b.ResetTimer()
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		ByteToFloat32(bys)
	}

}
func BenchmarkBinary2flt(b *testing.B) {
	//b.ResetTimer()
	//b.ReportAllocs()
	//b.ResetTimer()
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Binary2flt(bys)
	}

}

func BenchmarkByteToFloat(b *testing.B) {
	//b.ResetTimer()
	//b.ReportAllocs()
	//b.ResetTimer()
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		ByteToFloat(bys)
	}
}

func BenchmarkFloat32ToByte(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Float32ToByte(ft)
	}
}

func BenchmarkFlt2binary(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Flt2binary(ft)
	}
}

func readPhoto() ([]byte, error) {
	ff, err := os.Open("21.jpg")
	if err != nil {
		//fmt.Println("open error")
		return nil, err
	}
	defer ff.Close()
	out, err := ioutil.ReadAll(ff)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func TestImgToRGB(t *testing.T) {
	imgSrc, err := readPhoto()
	if err != nil {
		t.Errorf("error:%v\n", err)
		return
	}
	t.Logf("len=%v\n", len(imgSrc))

	ImgToRGB(imgSrc)
}

func BenchmarkImgToRGB(b *testing.B) {
	imgSrc, err := readPhoto()
	if err != nil {
		b.Errorf("error:%v\n", err)
		return
	}
	b.Logf("len=%v\n", len(imgSrc))
	b.N = 100
	for i := 0; i < b.N; i++ {

		ImgToRGB(imgSrc)
	}
}

func test(t *testing.T, errDone chan int) {
	select {
	case err := <-errDone:
		fmt.Printf("######err=%v\n", err)

	}
}
func TestTmp(t *testing.T) {
	//fchan := make(chan []float32, 3)
	//fchan <- []float32{}
	////close(fchan)
	//fchan <- nil
	//fchan <- nil
	////fchan <- nil
	//f1 := <-fchan
	//f2 := <-fchan
	//f3 := <-fchan
	//fmt.Println(f1, f2, f3, len(f1), len(f2), len(f3))
	errDone := make(chan int, 10)

	defer func() {
		for {
			select {
			case err, ok := <-errDone:
				if !ok {
					fmt.Printf("chan closed...\n")
					return
				}
				fmt.Printf("######err=%v\n", err)
			case <-time.After(time.Second * 5):
				fmt.Printf("timeout.....\n")
				return
			}
		}

	}()

	errDone <- 1111
	errDone <- 3333
	errDone <- 3333
	errDone <- 3333
	close(errDone)
	time.Sleep(time.Second * 5)
}

func testIORead(t *testing.T) {
	t.Logf("===len(%v)\n", len(bys))
	rd := bytes.NewReader(bys)
	buf := make([]byte, 11)
	for i := 0; i < 5; i++ {
		n, err := io.ReadFull(rd, buf)
		t.Logf("i=[%v],n=%v,buf=%v,err=%v\n", i, n, buf, err)
	}
}
func TestIORead(t *testing.T) {
	bys := make([]byte, 20)
	t.Logf("===len(%v)\n", len(bys))
	//rd := bytes.NewReader(bys)
	bw := bytes.NewBuffer(bys)
	buf := make([]byte, 3)
	for i := 0; i < 10; i++ {
		n, err := io.ReadFull(bw, buf)
		t.Logf("i=[%v],n=%v,buf=%v,err=%v\n", i, n, buf, err)
	}
	//ctx, cancel := context.WithCancel(context.Background())
}
func BenchmarkIORead(b *testing.B) {
	b.N = 100000
	t := testing.T{}
	for i := 0; i < b.N; i++ {
		testIORead(&t)
	}
}

func TestServer_BindCheck(t *testing.T) {
	//g:=gin.Default()
	s := server{}
	r := func() *gin.Engine {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("bookabledate", bookableDate)
		}
		g := gin.Default()
		g.POST("/test", func(c *gin.Context) {
			s.BindCheck(c)
		})

		g.GET("/bookable", func(c *gin.Context) { getBookable(c) })

		return g
	}()
	sv := httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
	defer sv.Close()
	//reqInfo := req{
	//	Url:    "1234",
	//	Imgs:   []string{"77"},
	//	Images: []photo{},
	//}
	//Json, _ := json.Marshal(&reqInfo)
	//fmt.Printf("Json=%v\n", string(Json))
	//rd := bytes.NewReader(Json)
	//rsp, _ := http.Post(sv.URL+"/test", "application/json", rd)
	//defer rsp.Body.Close()

	//?check_in=2018-03-08&check_out=2018-03-09
	//?check_in=2018-04-16&check_out=2018-04-17
	rsp, _ := http.Get(sv.URL + "/bookable?check_in=2019-10-08&check_out=2019-10-09")
	res, _ := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	t.Logf("===%v\n", string(res))
}

func TestServer_BindCheck_Req(t *testing.T) {
	//g:=gin.Default()
	s := server{}
	r := func() *gin.Engine {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("bookabledate", bookableDate)
		}
		g := gin.Default()
		g.POST("/test", func(c *gin.Context) {
			s.BindCheck(c)
		})

		g.GET("/bookable", func(c *gin.Context) { getBookable(c) })

		return g
	}()
	sv := httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
	defer sv.Close()
	reqInfo := req{
		Tm:      time.Now().Add(time.Second * 2),
		Flg:     true,
		Url:     "123",
		Num:     7,
		Urls:    []string{},
		Persons: [][]string{[]string{"a"}, []string{"b"}},
		Imgs:    []string{"11"},
		//Images:  []photo{{}, {}},
	}
	Json, _ := json.Marshal(&reqInfo)
	fmt.Printf("Json=%v\n", string(Json))
	rd := bytes.NewReader(Json)
	rsp, _ := http.Post(sv.URL+"/test", "application/json", rd)
	res, _ := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()

	//?check_in=2018-03-08&check_out=2018-03-09
	//?check_in=2018-04-16&check_out=2018-04-17
	//rsp, _ := http.Get(sv.URL + "/bookable?check_in=2019-10-08&check_out=2019-10-09")
	//res, _ := ioutil.ReadAll(rsp.Body)
	//defer rsp.Body.Close()
	t.Logf("===%v\n", string(res))
}

func TestSplit(t *testing.T) {
	str := "1.jpg,2.jpg,3.jpg"
	slcie := strings.Split(str, ",")
	t.Logf("len=%v,val=%v\n", len(slcie), slcie)
}

func TestChannel(t *testing.T) {
	cn := make(chan int)
	//ctx, cancel := context.WithCancel(contex.Background())
	ctx, _ := context.WithCancel(context.Background())
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	go func() {
		//defer wg.Done()

		select {
		case <-ctx.Done():
			fmt.Printf("ctx.done.......\n")
			//cancel()
			return
		case cn <- 77:
			fmt.Printf("777777.......\n")
			//cancel()
		}

	}()

	//cancel()
	//fmt.Printf("1.....\n")
	//cancel()
	//fmt.Printf("2.....\n")
	time.Sleep(time.Second * 3)
	<-cn
	//go func() {
	//	cn <- 77
	//	fmt.Printf("##########..........\n")
	//}()
	//
	//select {
	//case <-cn:
	//	fmt.Printf("##########\n")
	//}
	//fmt.Printf("after cn。。。\n")
	//cancel()
	//wg.Wait()
	time.Sleep(time.Second * 5)
}

func TestStrings(t *testing.T) {
	str := "ni hao? nihao?"
	sl := strings.Split(str, "?")
	fmt.Printf("out:%v,len:%v\n", sl, len(sl))
}

const str = "ni hao ？  nihao; nihao。\n"

func TestRegex(t *testing.T) {
	RegexTag := `(?:\s*([\?？\.。;；!！])\s*)`
	rgp, err := regexp.Compile(RegexTag)
	if err != nil {
		fmt.Printf("error(%v)\n", rgp)
		return
	}
	out := rgp.Split(str, len(str))
	fmt.Printf("out:%v,len:%v,len(str):%v,\n", out, len(out), len(str))
	strDeal := str
	count := 0
	for {
		loc := rgp.FindStringSubmatchIndex(strDeal)
		fmt.Printf("loc:%v\n", loc)
		if len(loc) != 0 {
			fmt.Printf("s:%v\n", strDeal[:loc[3]])
			count += loc[1]
			strDeal = str[count:]
			fmt.Printf("count:%v\n", count)
		} else {
			break
		}
	}
	ot := strings.Split(str, "\n")
	fmt.Printf("%v,%v\n", ot, len(ot))
	res := strings.Replace(str, "\r", "", -1)
	fmt.Printf("res:%v,len:%v\n", res, len(res))
}
func TestAppend(t *testing.T) {
	Append()
}
func BenchmarkAppend(b *testing.B) {
	b.N = 50000
	for i := 0; i < b.N; i++ {
		Append()
	}
}
func BenchmarkSlice(b *testing.B) {
	b.N = 10000
	for i := 0; i < b.N; i++ {
		Slice()
	}
}

//func TestMarkServer_Paragraph(t *testing.T) {
//	m := newMarkServer()
//	markGrpc := "192.168.5.25:1905"
//	conn, err := grpctest.Dial(markGrpc, grpctest.WithInsecure())
//	m.markCli = mark.NewMarkClient(conn)
//	if err != nil {
//		fmt.Printf("%s\n", err)
//		return
//	}
//	router := func() *gin.Engine {
//		router := gin.Default()
//		router.POST("/mark/paragraph", func(c *gin.Context) {
//			m.Paragraph(c)
//		})
//		return router
//	}()
//
//	ts := httptest.NewServer(http.HandlerFunc(router.ServeHTTP))
//	defer ts.Close()
//	func() {
//		req := MarkAPI{
//			Src:    "EN",
//			Trg:    "CN",
//			Inputs: PARAGRAPH,
//		}
//		out, _ := json.Marshal(&req)
//		rd := bytes.NewReader(out)
//		reqJson, err := http.NewRequest("POST", ts.URL+"/mark/paragraph", rd)
//		if err != nil {
//			t.Errorf("NewRequest failed,error(%v)\n", err)
//			return
//		}
//		cli := http.Client{}
//		rsp, err := cli.Do(reqJson)
//		if err != nil {
//			t.Errorf("client do failed,error(%v)\n", err)
//			return
//		}
//		defer rsp.Body.Close()
//		res, _ := ioutil.ReadAll(rsp.Body)
//		t.Logf("result:(%v)\n", string(res))
//	}()
//}

func TestBase64(t *testing.T) {
	ff := `{"name":"hhh","age":"33"}`
	in := []byte(ff)
	fmt.Printf("len str=%v,len byte=%v\n", len(ff), len(in))
	//oaKjpKWmpw==
	out := base64.StdEncoding.EncodeToString(in)
	fmt.Printf("out= %v, %v\n", out, len(out))
}

func TestFieldsFunc(t *testing.T) {
	str := "good。 good.hello，hi？"
	//str := "hi hi; hi;"
	out := strings.FieldsFunc(str, func(r rune) bool {
		switch r {
		case '.', '。', '?', '？', '!', '！', ';', '；':
			return true
		}
		return false
	})
	fmt.Printf("len(out)=%v,out:%v\n", len(out), out)
}

func splitTmp(input string) []string {
	inputH := input
	var res []string
	ind := 0
	//count := 0
	for len(inputH[ind:]) != 0 {
		loc := strings.Index(inputH[ind:], "...")
		if loc == -1 {
			res = append(res, inputH[ind:])
			break
		}
		res = append(res, inputH[ind:ind+loc+3])
		ind = ind + loc + 3
		//count++
	}
	//if count == 0 {
	//	res = append(res, input)
	//}
	return res
}

func TestHandleOmit(t *testing.T) {
	input := "hhhhh"
	res := strings.Split(input, "...")
	//fmt.Printf(">>%v,%v\n", input[6:], len(input[6:]))
	//res := handleOmit(input)
	fmt.Printf("%v, len=%v\n", res, len(res))
}
func TestBufioScanner(t *testing.T) {
	// Comma-separated list; last entry is empty.
	const input = "1,2,34..."
	loc := strings.Index(input, "...")
	fmt.Printf("loc=%v\n", loc)
	//f := NewFormat()
	//inputForm := f.Format(input)
	inputForm := input

	fmt.Printf("===%v|len=%v\n", inputForm, len(inputForm))
	scanner := bufio.NewScanner(strings.NewReader(inputForm))
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)
		for i := 0; i < dataLen; i++ {
			//if data[i] == ',' || data[i] == '?' {
			//	return i + 1, data[:i+1], nil
			//}
			switch data[i] {
			case '.', '?', '!', ';':
				return i + 1, data[:i+1], nil
			}

		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	sents := []string{}
	for scanner.Scan() {
		fmt.Printf("==%v\n", len(scanner.Text()))
		str := scanner.Text()
		if len(str) != 0 {
			sents = append(sents, str)
		}
	}
	fmt.Printf("%v,len=%v\n", sents, len(sents))
	for _, v := range sents {
		fmt.Printf("===>%v\n", v)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func replaceChineseCharTmp(input string) string {
	res := strings.Replace(input, "，", ",", -1)
	res = strings.Replace(res, "。", ".", -1)
	res = strings.Replace(res, "；", ";", -1)
	res = strings.Replace(res, "？", "?", -1)
	res = strings.Replace(res, "！", "!", -1)
	res = strings.Replace(res, "...", "\x00", -1) //先将省略号替换成0
	return res
}

func TestJpg(t *testing.T) {
	in := "/9j/"

	out, _ := base64.StdEncoding.DecodeString(in)
	fmt.Printf("===%x", out)
}

func TestFlt2binary(t *testing.T) {
	fa := []float32{7.3, 3.7, 6.7, 7.7}
	out, _ := Flt2binary(fa)
	fmt.Printf("%v\n", out)
}

func TestBinary2flt(t *testing.T) {
	ss := []byte{154, 153, 233, 64, 205, 204, 108, 64, 102, 102, 214, 64, 102, 102, 246, 64}
	out, _ := Binary2flt(ss)
	fmt.Printf("%v\n", out)
}

func TestSliceChan(t *testing.T) {
	ch := make(chan []byte, 10)
	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(string(data))
			}
		}
	}()
	data := make([]byte, 5)
	data = append(data, []byte("11111")...)
	ch <- data

	// fmt.Printf("%p\n", data)
	data = data[:0]
	// fmt.Printf("%p\n", data)

	data = append(data, []byte("aaa")...)
	ch <- data

	time.Sleep(time.Second * 5)

}

func testGenImg() ([]byte, error) {
	fl, err := os.Open("./20.jpg")
	if err != nil {
		fmt.Println("open file error")
		return nil, err
	}
	defer fl.Close()
	img, err := ioutil.ReadAll(fl)
	return img, err
}
func TestImgToRGB2(t *testing.T) {
	srcImg, err := testGenImg()
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	rgb, w, h, err := ImgToRGB(srcImg)
	fmt.Printf("len(rgb)=%v,w=%v,h=%v,err=%v\n", len(rgb), w, h, err)
}

func TestImgToRGBTmp(t *testing.T) {
	srcImg, err := testGenImg()
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	rgb, w, h, err := ImgToRGBTmp(srcImg)
	fmt.Printf("len(rgb)=%v,w=%v,h=%v,err=%v\n", len(rgb), w, h, err)
}

func BenchmarkImgToRGB2(b *testing.B) {
	srcImg, err := testGenImg()
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ImgToRGB(srcImg)
		}
	})
}
func BenchmarkImgToRGBTmp(b *testing.B) {
	srcImg, err := testGenImg()
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ImgToRGBTmp(srcImg)
		}
	})
}

func TestTrim(t *testing.T) {
	//in := ",,,122,333,333,,,"
	in := ""
	rs := strings.Split(strings.Trim(in, ","), ",")
	fmt.Printf("%v# %v   #%v#\n", rs, len(rs), rs[0])
	c1 := []byte{0, 49, 2, 222, 221, 81, 106, 154, 45, 14, 207, 249, 2, 78, 96, 11, 234, 237, 205, 198, 95, 47, 92, 68, 207, 250, 79, 35, 83, 205, 73, 91, 208, 22, 7, 254, 63, 31, 243, 215, 71, 68, 32, 177, 184, 165, 126, 153, 103, 222, 171}
	out := base64.StdEncoding.EncodeToString(c1)
	fmt.Printf("%v\n", out)
}

type person struct {
	name string
	age  string
}

func TestSwitchAssert(t *testing.T) { //断言
	s := &person{"yyn", "45"}
	//s = nil
	//var p person
	var tp interface{}
	tp = s
	//out:=tp.(type)
	fmt.Printf("tp=%v,s=%v\n", tp != nil, s)
	switch tp.(type) {
	case person:
		fmt.Printf("person...\n")
	case *person:
		fmt.Printf("ajjajaj\n")
		kk, ok := tp.(*person)
		fmt.Printf("ajjajaj  %v\n", ok)
		if !ok {
			fmt.Printf("jjjj\n")
			return
		}
		fmt.Printf("*person...,%s,%s,%v\n", kk.name, kk.age, ok)
	case nil:
		fmt.Printf("nil nil...\n")

	}

}

func TestContain(t *testing.T) {
	in := "hhh[p: jf]"
	out := strings.Contains(in, "[p:")
	fmt.Printf("=== %v\n", out)
	var inn io.Reader
	inn = os.Stdin
	out1, err := ioutil.ReadAll(inn)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}

	fmt.Printf("=== %v\n", out1)
}

func (p person) Name(n string) person {
	p.name = n
	return p
}
func TestUlog(t *testing.T) {
	p := person{}
	p = p.Name("lily")
	fmt.Printf("===%v\n", p.name)
	out := strings.Split("", ",")
	fmt.Printf("out=%v,len(out)=%v\n", out, len(out))
}
func TestForSwitchBreak(t *testing.T) {
	for i := 0; i < 7; i++ {
		switch i {
		case 0:
			fmt.Printf("i=%v\n", i)
		case 1:
			fmt.Printf("i=%v\n", i)
		case 2:
			fmt.Printf("i=%v\n", i)
			break //作用于switch
		case 3, 4, 5, 6:
			fmt.Printf("i=3,4,5,6\n")
		}
	}
}
func TestArrary(t *testing.T) {
	arr := [7]byte{2, 3, 4, 5, 6, 7, 1}
	length := len(arr)
	fmt.Printf("arrary=%v,length=%v\n", arr[:3], length)
}

func TestBufferReas(t *testing.T) {
	str := []byte("123456789")
	buf := bytes.NewBuffer(str)
	for i := 0; i < 17; i++ {
		p := make([]byte, 3)
		n, err := buf.Read(p)
		if err != nil {
			fmt.Printf("p:%v ,num=%v,err=%v\n", p[:n], n, err)
			return
		}
		fmt.Printf("p:%v num=%v,err=%v\n", p[:n], n, err)
	}
}

func TestGoutHeader(t *testing.T) {
	router := func() *gin.Engine {
		router := gin.Default()
		router.GET("/test/header", func(c *gin.Context) {
			c.Writer.Header().Add("sid", "sid-ok")
			c.Writer.Header().Add("Content-Length", "1878")
		})
		return router
	}()

	ts := httptest.NewServer(http.HandlerFunc(router.ServeHTTP))

	g := gout.New(nil)

	type testHeader struct {
		Sid           string `header:"sid"`
		ContentLength string `header:"Content-Length"`
		Code          int
	}

	var tHeader testHeader
	err := g.GET(ts.URL + "/test/header").BindHeader(&tHeader).Code(&tHeader.Code).Do()
	if err != nil {
		return
	}
	fmt.Printf("Sid:%v\n", tHeader)

}
