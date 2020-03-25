package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"
)

const CnEngineContainEnglishText = `(?:\s*[\w':]+\s*)`

type EngText struct {
	Text  string `json:"text"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func dealReg(inputText string) []EngText {
	reg, err := regexp.Compile(CnEngineContainEnglishText)
	if err != nil {
		panic(err)
	}
	diffLen := 0
	engText := make([]EngText, 0)
	//fmt.Println("需要处理的字符串长度为len(input):", len(inputText))
	for diffLen < len(inputText) {
		tmp, eTxt := dealRegOne(reg, diffLen, inputText[diffLen:])
		diffLen += tmp
		if eTxt.Text != "" {
			engText = append(engText, eTxt)
		}
		//fmt.Println("diffLen==", diffLen)
	}
	//fmt.Println(engText, len(engText))
	return engText
}
func dealRegOne(reg *regexp.Regexp, diff int, in string) (int, EngText) {
	loc := reg.FindStringSubmatchIndex(in)
	if len(loc) == 0 {
		return len(in), EngText{}
	}
	//fmt.Println(loc)
	eTxt := EngText{
		Text:  in[loc[0]:loc[1]],
		Start: loc[0] + diff,
		End:   loc[1] + diff,
	}
	return loc[1], eTxt
}

//func (a *asrClient) resolveWordWithBlank(engText []EngText, old string) string {
//	var resStr strings.Builder
//	num := len(engText)
//	var r *strings.Replacer
//	for i := 0; i < num+1; i++ {
//		n1 := 0
//		if i > 0 {
//			n1 = engText[i-1].End
//		}
//		n2 := len(old)
//		if i < num {
//			n2 = engText[i].Start
//		}
//		if n1 < n2 {
//			//去掉中文里的空格
//			r = strings.NewReplacer("<eps>", "", "</s>", "，", "<s>", "", "<unk>", "", "<s", "", "<eps", "", " ", "")
//			//resStr += r.Replace(old[n1:n2])
//			resStr.WriteString(r.Replace(old[n1:n2]))
//		}
//		if i < num {
//			//英文单词直接保存
//			resStr.WriteString(engText[i].Text)
//		}
//
//	}
//	//空格保留，去掉<eps></s><s><unk>等
//	r = strings.NewReplacer("<eps>", "", "</s>", "，", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//	//将结果返回
//	return r.Replace(resStr.String())
//}
func TestRegexp(t *testing.T) {
	old := "htc g 十一 跟 三星 i 九零 零一 哪个 好 </s>"
	eng := dealReg(old)
	jsBytes, _ := json.Marshal(&eng)
	fmt.Printf("text:%v\n", string(jsBytes))
	var tmp string = "jsadjg"
	fmt.Printf("tmp:%v\n", tmp)
}

func TestUpper(t *testing.T) {
	eu := englishUpper{
		firstFix:    true,
		lastFixFlag: false,
	}
	s1 := "hello world,the day is sunny. everyone is happy"
	eu.toUpper(&s1)
	fmt.Printf("=======s1(%v)\n", s1)
	stringSlice := []string{
		//"hello world.",
		//"the day is sunny,",
		//"hello boy?   everyone. is happy,",
		//"are you kidding?",
		//"ok,everybody.",
		"it can be in the toilet operational firearms.",
		"command of specialist protection was the plan british transport police have put out an alert on a possible suicide bomber are attempting to board a london,",
		"bound service we've been ordered to stop the train at barnet,",
		"shed its derelict depot out in the sticks.",
	}
	eu2 := englishUpper{
		firstFix:    true,
		lastFixFlag: false,
	}
	for i := 0; i < len(stringSlice); i++ {
		fmt.Printf("====i=%v\n", i)
		eu2.toUpper(&stringSlice[i])
		fmt.Printf("&stringSlice[i]=%v\n", stringSlice[i])

	}
	res := strings.Join(stringSlice, " ")
	fmt.Printf("===res(%v)\n", res)
}

type englishUpper struct {
	firstFix    bool
	lastFixFlag bool
}

func onceHandle(text *string) {
	srcStr := []byte(*text)
	srcLen := len(srcStr)
	tag := false
	for i := 0; i < srcLen; i++ {
		if srcStr[i] == '.' || srcStr[i] == '?' {
			tag = true
			continue
		}
		if srcStr[i] != ' ' && tag == true {
			srcStr[i] = bytes.ToUpper([]byte(srcStr)[i : i+1])[0]
			tag = false
		}
	}
	*text = string(srcStr)
}

func (e *englishUpper) toUpper(text *string) {
	srcStr := []byte(strings.TrimSpace(*text))
	srcLen := len(srcStr)
	if srcLen == 0 {
		return
	}
	tag := false
	for i := 0; i < srcLen; i++ {
		if srcStr[i] == '.' || srcStr[i] == '?' {
			tag = true
			continue
		}
		if srcStr[i] != ' ' && tag == true {
			srcStr[i] = bytes.ToUpper([]byte(srcStr)[i : i+1])[0]
			tag = false
		}
	}
	if e.firstFix == true {
		srcStr[0] = bytes.ToUpper([]byte(srcStr)[:1])[0]
		e.firstFix = false
	}
	if e.lastFixFlag == true {
		srcStr[0] = bytes.ToUpper([]byte(srcStr)[:1])[0]
		e.lastFixFlag = false
	}
	if srcStr[srcLen-1] == '.' || srcStr[srcLen-1] == '?' {
		e.lastFixFlag = true
	}
	*text = string(srcStr)
}

func TestReplace(t *testing.T) {
	r := strings.NewReplacer("a", "FF")
	res := r.Replace("")
	fmt.Printf("res:%v,len(%v),%v\n", res, len(res), res == "")
	ss := ""
	rs := strings.Split(ss, "</s>")
	fmt.Printf("rs:%v,len:%v\n", rs, len(rs[0]))
	var str strings.Builder
	fmt.Printf("str:%v,len:%v,%v\n", str.String(), len(str.String()), 17/2+1)
}

func TestString(t *testing.T) {
	tmpText := "你</s>好</s>朋</s>友</s>嗯</s>大</s>家</s>好</s>"
	tmpText = "生</s>"
	increment := 2
	if increment < 1 {
		increment = 1
	}
	fmt.Printf("index:%v,len:%v\n", strings.IndexAny(tmpText[3+4:], "</s>"), len(tmpText))
	//tmpText := strings.TrimSpace(*text)
	//replace := strings.NewReplacer("<eps>", "", "</s>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
	//tmpText = replace.Replace(tmpText)
	textSections := strings.Split(tmpText, "</s>")
	fmt.Printf("textSections:%v,%v\n", textSections, len(textSections))
	lenTexts := len(textSections)
	texts := make([]string, lenTexts/increment+1)
	var str strings.Builder
	j := 0
	find := 1
	for i, v := range textSections {
		replace := strings.NewReplacer("<eps>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
		tmpText = replace.Replace(v)
		//engText := dealReg(tmpText)
		//s1 := a.dealWordWithBlank(engText, tmpText)
		str.WriteString((tmpText))
		if find == increment {
			texts[j] = str.String()
			j++
			str.Reset()
			find = 1
			continue
		}
		find++
		if i == lenTexts-1 {
			texts[j] = str.String()
			j++
		}
	}
	fmt.Printf("texts:%v, %v,\n", texts, len(texts))
}

//func (a *asrClient) beforePuncResumeSplitText(text *string, increment int) []string {
//	if increment < 1 {
//		increment = 1
//	}
//	tmpText := strings.TrimSpace(*text)
//	//replace := strings.NewReplacer("<eps>", "", "</s>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//	//tmpText = replace.Replace(tmpText)
//	if a.eType == "chinese" {
//		//for i, v := range tmpText {
//		//	utf8.RuneLen()
//		//}
//		textSections := strings.Split(tmpText, "</s>")
//		lenTexts := len(textSections)
//		texts := make([]string, lenTexts/increment+1)
//		var str strings.Builder
//		j := 0
//		find := 1
//		for i, v := range textSections {
//			tmpText = PuncReplace.Replace(v)
//			engText := dealReg(tmpText)
//			str.WriteString(dealWordWithBlank(engText, tmpText))
//			if find == increment {
//				texts[j] = str.String()
//				str.Reset()
//				find = 1
//				j++
//				continue
//			}
//			find++
//			if i == lenTexts-1 {
//				texts[j] = str.String()
//				j++
//			}
//		}
//		return texts
//	}
//	//replace := strings.NewReplacer("<eps>", "", "</s>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//	tmpText = EnglilshReplace.Replace(tmpText)
//	return []string{strings.TrimSpace(tmpText)}
//
//}

//func (a *asrClient) beforePuncResumeSplitTextTmp(text *string, increment int) []string {
//	if increment < 1 {
//		increment = 1
//	}
//	tmpText := strings.TrimSpace(*text)
//	//replace := strings.NewReplacer("<eps>", "", "</s>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//	//tmpText = replace.Replace(tmpText)
//	//tmpText := "你</s>好</s>朋</s>友</s>嗯</s>大</s>家</s>好</s>"
//	if a.eType == "chinese" {
//		start := 0
//		for {
//			start := strings.IndexAny(tmpText[start:], "</s>")
//			if start == -1 {
//				break
//			}
//			start = start + 4
//			//start = strings.IndexAny(tmpText[start:], "</s>")
//		}
//
//		textSections := strings.Split(tmpText, "</s>")
//		lenTexts := len(textSections)
//		texts := make([]string, lenTexts/increment+1)
//		var str strings.Builder
//		j := 0
//		find := 1
//		for i, v := range textSections {
//			replace := strings.NewReplacer("<eps>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//			tmpText = replace.Replace(v)
//			engText := dealReg(tmpText)
//			//s1 := a.dealWordWithBlank(engText, tmpText)
//			str.WriteString(a.dealWordWithBlank(engText, tmpText))
//			if find == increment {
//				texts[j] = str.String()
//				str.Reset()
//				find = 1
//				j++
//				continue
//			}
//			find++
//			if i == lenTexts-1 {
//				texts[j] = str.String()
//				j++
//			}
//		}
//		return texts
//	} else {
//		replace := strings.NewReplacer("<eps>", "", "</s>", "", "<s>", "", "<unk>", "", "<s", "", "<eps", "")
//		tmpText = replace.Replace(tmpText)
//		return []string{strings.TrimSpace(tmpText)}
//	}
//}
func TestAssert(t *testing.T) {
	bl := assert.Equal(t, "aa", "aa")
	if bl {
		fmt.Printf("===ok...\n")
	}
	bl = assert.NoError(t, nil)
	fmt.Printf("bl=%v\n", bl)
	bl = assert.True(t, true, "www")
	fmt.Printf("bl=%v\n", bl)
	bl = assert.False(t, false, "www")
	fmt.Printf("bl=%v\n", bl)
	bl = assert.Error(t, errors.New("err"), "www")
	fmt.Printf("bl=%v\n", bl)

}

func TestSplit(t *testing.T) {
	str := "世界"
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
	str = "hello世界"
	buf := bytes.Buffer{}
	for i, v := range str {

		//r, size := utf8.DecodeRuneInString(string(v))
		//fmt.Printf("rsize:%v\n", utf8.RuneLen(v))
		buf.WriteString(string(v))
		fmt.Printf("i:%v,r:%c,size:%v\n", i, v, utf8.RuneLen(v))
	}
	fmt.Printf("output:%v,len:%v\n", buf.String(), buf.Len())
}

func haveSigned(bs string, start int) bool {
	if start == -1 {
		return false
	}

	var ok int

	for i := start; i < len(bs); i++ {
		if bs[i] == '<' {
			ok++
			continue
		}

		if bs[i] == '>' {
			ok--
			continue
		}
	}

	return ok != 0
}

func TestAAA(t *testing.T) {
	var buf strings.Builder
	//s := "我是中国人在a云知声"
	s := "a"
	exit := false
	var allLine []string

	for k, v := range s {
		fmt.Printf("1.%c\n", v)
		l := utf8.RuneLen(v)

		exit = k+l == len(s)

		if buf.Len()+l <= 6 {
			if l == 1 {
				buf.WriteByte(byte(v))
			} else {
				buf.WriteRune(v)
			}
			continue
		}
		fmt.Printf("2.%c\n", v)

		bs := buf.String()
		buf.Reset()

		allLine = append(allLine, bs)

		if l == 1 {
			buf.WriteByte(byte(v))
		} else {
			buf.WriteRune(v)
		}

	}

	if exit {
		if buf.Len() > 0 {
			allLine = append(allLine, buf.String())
		}
	}

	for i, v := range allLine {
		fmt.Printf("%d:%s\n", i, v)
	}
}

func TestSplitForPunc(t *testing.T) {
	text := "我是中国人在云知声"
	//text = "你</s>好</s>朋</s>友</s>嗯</s>大</s>家</s>好</s>"
	res := splitTextForPunc(text, 9)
	fmt.Printf("res:%v\n", res)
	var ss []string
	for i, v := range ss {
		fmt.Printf("%v,%v\n", i, v)
	}
}
func splitTextForPunc(text string, stepSize int) []string {
	var buf strings.Builder
	exit := false
	var allLine []string
	for k, v := range text {
		fmt.Printf("1.%c\n", v)
		l := utf8.RuneLen(v)
		exit = k+l == len(text)
		if buf.Len()+l <= stepSize {
			if l == 1 {
				buf.WriteByte(byte(v))
			} else {
				buf.WriteRune(v)
			}
			continue
		}
		fmt.Printf("2.%c\n", v)
		bs := buf.String()
		buf.Reset()
		allLine = append(allLine, bs)
		if l == 1 {
			buf.WriteByte(byte(v))
		} else {
			buf.WriteRune(v)
		}
	}
	if exit {
		if buf.Len() > 0 {
			allLine = append(allLine, buf.String())
		}
	}
	return allLine
}

type zerolog struct {
	w io.Writer
}

func New(w ...io.Writer) *zerolog {
	return &zerolog{w: io.MultiWriter(w...)}
}

func (z *zerolog) Msg(s string) {
	//fmt.Printf("222222\n")
	n, err := z.w.Write([]byte(s))
	fmt.Printf("n:%v,err:%v\n", n, err)
}

type filter struct {
	w    io.Writer
	drop bool
}
type filter1 struct {
	w    io.Writer
	drop bool
}

func NewFilter1(w io.Writer, drop bool) *filter1 {
	if drop {
		return &filter1{w: w, drop: drop}
	}

	return &filter1{w: w}
}
func (f *filter1) Write(p []byte) (n int, err error) {
	//fmt.Printf("1111111\n")
	if f.drop {
		return f.w.Write(p)

	}
	return 0, nil
}
func NewFilter(w io.Writer, drop bool) *filter {
	if drop {
		return &filter{w: w, drop: drop}
	}

	return &filter{w: w}
}

func (f *filter) Write(p []byte) (n int, err error) {
	//fmt.Printf("33333\n")
	if f.drop {
		return 0, nil
	}

	return f.w.Write(p)
}

func TestFilter(t *testing.T) {
	//New(os.Stdout).Msg("hello\n")
	//New(NewFilter(os.Stdout, true)).Msg("drop hello\n")
	//New(NewFilter(os.Stdout, false)).Msg("accept hello\n")
	//, NewFilter1(os.Stdout, true)
	New(NewFilter(os.Stdout, true), NewFilter(os.Stdout, false)).Msg("==\n")
}
