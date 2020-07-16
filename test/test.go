package test

import (
	"fmt"
	"testing"
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
func Test(t *testing.T) {
	fmt.Printf("====")
}

//func TestProxy(t *testing.T) {
//	ps := testNewProxy(t)
//	router := func() *gin.Engine {
//		router := gin.Default()
//		router.POST("/n2t/en", func(c *gin.Context) {
//			ps.ResolveEn(c)
//		})
//		return router
//	}()
//	ts := httptest.NewServer(http.HandlerFunc(router.ServeHTTP))
//	defer ts.Close()
//
//	data := make(url.Values)
//	data["text"] = []string{`{"Version":1,"DisplayText":"pack","Markers":[{"Type":"phone","Position":{"Start":0,"Length":4},"Value":["ˈp·æ·k"]}]}`} //prəˌnʌnsiˈeɪʃn
//	data["text"] = []string{`a[ps:æ] a[ps:æ] apple`}
//	data["text"] = []string{`pronunciation[p:p r ə ˌn ʌn s i ˈeɪ ʃn]`} //prəˌnʌnsiˈeɪʃn
//	res, err := http.PostForm(ts.URL+"/n2t/en", data)
//	if err != nil {
//		t.Errorf("===error(%v)\n", err)
//		return
//	}
//	defer res.Body.Close()
//	outRes, _ := ioutil.ReadAll(res.Body)
//	t.Logf("res:%v\n", string(outRes))
//
//}
