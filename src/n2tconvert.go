package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"n2tconvert/parser"
	"os"
	"regexp"
	"strings"
)

var input = flag.String("input", "", "the input to the text parse")

//const eng = `(?:[\w: ]+)`  `(?:(\[c[:：]\])|(\[t[:：][rf]\])|(\[s[:：]\]))`
const checkRegexTag = `(?:\[c[:：]\]|\[t[:：][rf]\]|\[s[:：]\])`

func checkTag(in string) (bool, error) {
	reg, err := regexp.Compile(checkRegexTag)
	if err != nil {
		return false, err
	}
	loc := reg.FindStringSubmatchIndex(in)
	fmt.Println("=====", loc)
	if len(loc) == 0 {
		return false, nil
	}
	return true, nil
}

type EngText struct {
	text  string
	start int
	end   int
}

func dealReg(inputText string) ([]EngText, error) {
	reg, err := regexp.Compile(`(?:\s*[\w':]+\s*)`)
	if err != nil {
		return nil, err
	}
	diffLen := 0
	//diff := 0
	engText := make([]EngText, 0)
	fmt.Println("需要处理的字符串长度为len(input):", len(inputText))
	for diffLen < len(inputText) {
		tmp, eTxt := dealRegOne(reg, diffLen, inputText[diffLen:])
		diffLen += tmp
		if eTxt.text != "" {
			engText = append(engText, eTxt)
		}
		fmt.Println("diffLen==", diffLen)
	}
	fmt.Println(engText, len(engText))
	return engText, nil
}
func dealRegOne(reg *regexp.Regexp, diff int, in string) (int, EngText) {
	loc := reg.FindStringSubmatchIndex(in)
	if len(loc) == 0 {
		return len(in), EngText{}
	}
	fmt.Println(loc)
	eTxt := EngText{
		text:  in[loc[0]:loc[1]],
		start: loc[0] + diff,
		end:   loc[1] + diff,
	}
	return loc[1], eTxt
}

func main() {
	flag.Parse()

	if len(*input) == 0 {
		flag.Usage()
		fmt.Printf("please give an input\n")
		return
	}

	ts, err := parser.NewTextScan()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse.NewTextScan error: %v\n", err)
		os.Exit(1)
	}

	res, err := ts.RunScan(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rs.RunScan error: %v\n", err)
		os.Exit(1)
	}

	resStr, err := json.MarshalIndent(res, "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.MarshalIndent error: %v\n", err)
		return
	}

	fmt.Println("res:", string(resStr), "\n", res.DisplayText, len(res.DisplayText))

	strErr := `They are very beautiful`
	fmt.Println("strErr=", strErr)
	strOk := `   {"ddisplayText:"They are very beautiful","Language":"en"}  `
	//strOk := `{}`
	fmt.Printf("strOk=%c\n", strOk[1])
	var txt PlayTxt
	err = json.Unmarshal([]byte(strOk), &txt)
	fmt.Println("err=", err, "txt=", txt)
	bl := strings.Contains(strOk, `}`) && strings.Contains(strOk, `{`)
	fmt.Println("bl==", bl)
	bll := strings.EqualFold("abs", "AbS")
	fmt.Println("bll==", bll)
	out := strings.ToUpper("sh__jkasjf")
	fmt.Println(out)
	tt("12334")
	//jj := `{"Version":"1","DisplayText":"Jsgf Grammar Tool Generated","GrammarWeight":"{\"weight_struct\":[]}","Grammar":"#JSGF V1.0 utf-8 cn;\ngrammar main;\npublic <main> = \"<s>\"(no|not yet|no the internet is not back)\"</s>\";\n"}`
	jk := `{"Grammar" : "#enumerate \nIt's Tuesday.\nTuesday.\nToday is Tuesday.\nIt's Tuesday today.\n", "Version" : 1, "DiSplAyText" : "adjhh", "GrammarWeight" : "" }`
	err = json.Unmarshal([]byte(jk), &txt)
	//jbl := json.Valid([]byte("dahsd"))
	fmt.Println("jbl:", txt.DisplayText)
	s1 := " {fh df} "
	out = strings.Replace(s1, " ", "", -1)
	fmt.Println(out, len(out), len(s1))
	out = strings.TrimSpace(s1)
	bint := checkJsonFormatErr(s1)
	fmt.Println(len(out), len(s1), bint)
	a := "abc"
	b := "Abc"
	i := strings.Compare(a, b)
	fmt.Println("i=", i)
	i = strings.Count(b, "")
	fmt.Println("i=", i)
	bl = strings.EqualFold(a, b)
	fmt.Println(bl)
	ss := "addgjghhgnm"
	outss := strings.Split(ss, "m")
	fmt.Println(outss, len(outss))
	//patterns := []string{
	//	"y", "25",
	//	"中", "国",
	//	"中工", "家伙",
	//}

	/*
		 patterns := make([]string,270 * 2)
		 for i :=0;i< 270 *2;i++{
			 patterns[i] = fmt.Sprintf("%d",i)
		 }
	*/
	replacer := strings.NewReplacer("<eps>", "", "</s>", "，", "<s>", "", "<unk>", "", "<s", "", "<eps", "", " ", "")
	//replacer := strings.NewReplacer("<eps>", "", "</s>", ",", "<s>", "", "<unks>", "", "<s", "", "<eps", "")
	format := "<。"
	strfmt := replacer.Replace(format)
	fmt.Println("\nmain() replacer.Replace old=", format)
	fmt.Println("main() replacer.Replace new=", strfmt)

	fmt.Println(len("。"))

	dealReg(*input)
	fmt.Println(*input)

	gg := "hsddd"
	hh := gg[0:2]
	hh = "kl"
	fmt.Println(gg, hh)

	var strcop strings.Builder
	strcop.WriteString("hello ")
	strcop.WriteString("everyone")
	dd := strcop.String()
	fmt.Println(dd, len(dd))

	//exp := `{"DisplayText":"They called it the \"ballpoint\"[s:] pen."}`
	//checkTag(exp)
	checkTag(*input)

	//reg, _ := regexp.Compile(`(?:[\w: ]+)`)
	//loc := reg.FindStringSubmatchIndex("kk")
	//fmt.Println("==", loc)
	//out := strings.Replace("far[c:]away", "[c:]", "    ", 1)
	////fmt.Println("out=", out)
	//s1 := "012345"
	//fmt.Println(s1[6-1:])
	//str := "Please get off[s:] at the next bus stop."
	////str := `{"DisplayText":"Please get off[s:] at the next bus stop."}`
	//js := PlayTxt{}
	//err = json.Unmarshal([]byte(str), &js)
	//if err != nil {
	//	fmt.Println("json Unmarshal failed")
	//	fmt.Println("js.DisplayText:", js.DisplayText, js)
	//	return
	//}
	//fmt.Println("json Unmarshal success")
	//fmt.Println("js.DisplayText:", js.DisplayText, js)
}

type PlayTxt struct {
	DisplayText string
	errBody
}
type errBody struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func tt(mode string) {
	mode = "99999"
	mode = "hhhhh"
	fmt.Println(mode)
}
func checkJsonFormatErr(input string) int {
	tmp := strings.TrimSpace(input)
	length := len(tmp)
	if length > 1 {
		if (tmp[0] == '{' && tmp[length-1] == '}') || (tmp[0] == '[' && tmp[length-1] == ']') {
			return 77
		}
	}
	return 0
}
