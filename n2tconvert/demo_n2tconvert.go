package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"myself/n2tconvert/dc"
	"myself/n2tconvert/parser"
	"os"
	"reflect"
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

type ParseTagResult struct {
	Grammar       string
	GrammarWeight string
	dc.Result
}

func parseTag(txtStr string) (parseRes ParseTagResult, errCode int, err error) {
	ts, _ := parser.NewTextScan()
	res, err := ts.RunScan(txtStr)
	if err != nil {
		return ParseTagResult{}, -1, err
	}

	parseRes = ParseTagResult{
		Grammar:       "",
		GrammarWeight: "",
		Result:        res,
	}
	return parseRes, 0, nil
}

//var mapJson map[string] interface{}

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

	//////////
	mapJson := make(map[string]interface{})
	val := reflect.ValueOf(&res)
	val = val.Elem()

	ty := val.Type()
	for i := 0; i < ty.NumField(); i++ {
		sf := ty.Field(i)
		sv := val.Field(i)
		fmt.Println("===", sf.Type, sf.Name, sf)
		mapJson[sf.Name] = sv.Interface()
		fmt.Println(">>>", sv, mapJson[sf.Name])
	}
	out, _ := json.Marshal(mapJson)
	fmt.Println("out:", string(out))
	////
	map2 := make(map[string]interface{})
	json.Unmarshal(resStr, &map2)
	fmt.Println("map2=", map2)

	v, ok := map2["DisplayText"]
	s, o := v.(int)
	fmt.Println(v, ok, s, o)

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
