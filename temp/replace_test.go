package temp

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type Format struct {
	//FF          string `"fjakl"`
	r           *strings.Replacer
	insertSpace *regexp.Regexp
	trimSpace   *regexp.Regexp
}

func NewFormatT() *Format {
	space := regexp.MustCompile(`\s+`)
	insert := regexp.MustCompile(`([,.?<>;!\x00])`)
	return &Format{r: strings.NewReplacer(
		"，", ",",
		"。", ".",
		"；", ";",
		"？", "?",
		"！", "!",
		"...", "\x00", //先将省略号替换成0
		"“", "\"",
		"”", "\"",
		"‘", "'",
		"’", "'",
		"《", "<",
		"》", ">",
	), trimSpace: space, insertSpace: insert}
}

func NewFormat() *Format {
	//space := regexp.MustCompile(`\s+`)
	//insert := regexp.MustCompile(`([,.?<>;!\x00])`)
	return &Format{r: strings.NewReplacer(
		"，", ",",
		"。", ".",
		"；", ";",
		"？", "?",
		"！", "!",
		"...", "\x00", //先将省略号替换成0
		"“", "\"",
		"”", "\"",
		"‘", "'",
		"’", "'",
		"《", "<",
		"》", ">",
	)}
}

type testFormat struct {
	needFormat string
	needAfter  string
}

func (f *Format) Format(input string) string {
	rv := f.r.Replace(input)
	//rv = f.insertSpace.ReplaceAllString(rv, " $1 ")
	//rv = f.trimSpace.ReplaceAllString(rv, " ")
	//rv = strings.TrimSpace(rv)
	return rv
}

//func TestFormat(t *testing.T) {
//	test := []testFormat{
//		{needFormat: "                ，？。", needAfter: ", ? ."},
//		{needFormat: "hello          world", needAfter: "hello world"},
//		{needFormat: "hello,?.world", needAfter: "hello , ? . world"},
//	}
//
//	f := NewFormat()
//	for _, v := range test {
//		outFormat := f.Format(v.needFormat)
//		if outFormat != v.needAfter {
//			t.Errorf("got:%s, want:%s\n", outFormat, v.needAfter)
//		}
//	}
//}

func replaceChineseChar(intput string) string {
	res := strings.TrimSpace(intput)
	res = strings.Replace(res, "，", ",", -1)
	res = strings.Replace(res, "。", ".", -1)
	res = strings.Replace(res, "；", ";", -1)
	res = strings.Replace(res, "？", "?", -1)
	res = strings.Replace(res, "！", "!", -1)
	res = strings.Replace(res, "...", "\x00", -1) //先将省略号替换成0
	res = strings.Replace(res, "“", "\"", -1)
	res = strings.Replace(res, "”", "\"", -1)
	res = strings.Replace(res, "‘", "'", -1)
	res = strings.Replace(res, "’", "'", -1)
	res = strings.Replace(res, "《", "<", -1)
	res = strings.Replace(res, "》", ">", -1)
	return res
}
func TestReplaceChineseChar(t *testing.T) {
	strSlice := []string{
		"good morning,“good” moring.‘b’ye bye... bye bye... ok.ok,",
	}
	for _, v := range strSlice {
		res := replaceChineseChar(v)
		fmt.Printf("===%v , len=%v\n", res, len(res))
	}
}

func BenchmarkReplaceR(b *testing.B) {
	str := "good morning,“good” moring.‘b’ye bye... bye bye... ok.ok,"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceChineseChar(str)
	}
}

func replace(f *Format, str string) {
	f.Format(str)
}

//func TestReplace(t *testing.T) {
//	str := "good morning,“good” moring.‘b’ye bye... bye bye... ok.ok,"
//	f := NewFormat()
//	out := f.Format(str)
//	fmt.Printf("out:%v\n", out)
//	fmt.Printf("out:%v\n", replaceChineseChar(str))
//}
func BenchmarkReplace(b *testing.B) {
	str := "good morning,“good” moring.‘b’ye bye... bye bye... ok.ok,"
	f := NewFormat()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replace(f, str)
	}
}
