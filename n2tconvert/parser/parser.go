package parser

import (
	"bytes"
	"errors"
	"myself/n2tconvert/dc"
	"regexp"
	"strings"
	"unicode"
)

const (
	curVersion = 1
)

// TextScan 用来遍历文本处理
type TextScan struct {
	tagParsers []TagParser
	textHandle string
}

func initTagParsers() ([]TagParser, error) {
	tagParsers := []TagParser{}
	tagParsers = append(tagParsers, &SoundLinkingParser{})
	tagParsers = append(tagParsers, &ToneParser{})
	tagParsers = append(tagParsers, &StressParser{})
	tagParsers = append(tagParsers, &PronouncePaeser{})
	for _, parser := range tagParsers {
		if err := parser.InitData(); err != nil {
			return nil, err
		}
	}
	return tagParsers, nil
}

// NewTextScan 用来获取TextScan
func NewTextScan() (*TextScan, error) {
	tagParsers, err := initTagParsers()
	if err != nil {
		return nil, err
	}
	return &TextScan{
		tagParsers: tagParsers,
	}, nil
}

// RunScan 用来进行遍历文本处理
func (ts *TextScan) RunScan(input string) (dc.Result, error) {
	var res dc.Result
	var buf bytes.Buffer
	handledLength := 0
	tagLenInAll := 0
	//fmt.Println("总长度：", len(input))
	if len(input) == 0 {
		return res, errors.New("text handled null")
	}
	ts.textHandle = input //save srcInput
	for handledLength < len(input) {
		handleStr := input[handledLength:]
		oneHandledLength, tagLen, err := ts.RunOne(handleStr, handledLength, tagLenInAll, &buf, &res)
		if err != nil {
			return res, err
		}
		tagLenInAll += tagLen
		handledLength += oneHandledLength
		//fmt.Println(">>handledLength循环里=", handledLength)
	}
	//fmt.Println(">>handledLength循环外=", handledLength)
	res.Version = curVersion
	res.DisplayText = buf.String()
	return res, nil
}

func (ts *TextScan) isLastTag(handleStr string) bool {
	var loc []int
	for _, tagParser := range ts.tagParsers {
		loc = tagParser.TagRegex().FindStringSubmatchIndex(handleStr)
		if len(loc) != 0 {
			//fmt.Println("###", loc)
			return false
		}
	}
	return true
}

// RunOne 用来处理一次标记, 返回值string表示还需要处理的文本
func (ts *TextScan) RunOne(handleStr string, diff int, tagLenInAll int, buf *bytes.Buffer, res *dc.Result) (int, int, error) {
	var curParser TagParser
	var curLoc []int
	for _, tagParser := range ts.tagParsers {
		loc := tagParser.TagRegex().FindStringSubmatchIndex(handleStr)
		//fmt.Println(">>>>>", loc)
		if len(loc) == 0 {
			continue
		}
		// 获取最早出现的loc
		if len(curLoc) == 0 {
			curLoc = loc
			curParser = tagParser
		} else {
			if loc[0] < curLoc[0] {
				curLoc = loc
				curParser = tagParser
			}
		}
	}

	// 如果没有匹配
	if len(curLoc) == 0 {
		buf.WriteString(handleStr)
		return len(handleStr), tagLenInAll, nil
	}

	// 查看匹配结果是否存在问题
	tagStr, err := curParser.GetMatchTag(handleStr, curLoc)
	if err != nil {
		return 0, tagLenInAll, err
	}

	var marker dc.TagMarker
	//todo:error judge
	err = curParser.SetMarker(handleStr, diff, tagLenInAll, curLoc, &marker)
	if err != nil {
		return 0, tagLenInAll, errors.New("SetMarker error")
	}
	res.Markers = append(res.Markers, marker)

	handledLength := curParser.GetHandledLength(curLoc)
	buf.WriteString(handleStr[0:curLoc[0]])

	tagLen := 0
	tmpLen := handledLength + curLoc[0] + diff
	//fmt.Println("Runone里面", tmpLen)
	//fmt.Printf("====curLoc=%v\n", curLoc)
	if curParser.TagRegex().String() == soundlinkingRegexTag {
		buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, " ", 1))
		tagLen = len(tagStr) - 1
	}
	if curParser.TagRegex().String() == toneRegexTag {
		buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
		tagLen = len(tagStr)
	}
	if curParser.TagRegex().String() == stressRegexTag {
		//判断是否为最后一个标记
		if tmpLen < len(ts.textHandle) {
			//判断是否为最后一个标记
			if ts.isLastTag(ts.textHandle[tmpLen:]) {
				//fmt.Println("###this is the last tag:tmpLen=", tmpLen)
				//判断此位置的字符是不是标点及空格
				if isPunctuation(ts.textHandle[tmpLen]) {
					buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
					tagLen = len(tagStr)
				} else {
					buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, " ", 1))
					tagLen = len(tagStr) - 1
				}

			} else {
				//判断此位置的字符是不是标点及空格
				if isPunctuation(ts.textHandle[tmpLen]) {
					buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
					tagLen = len(tagStr)
				} else {
					buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, " ", 1))
					tagLen = len(tagStr) - 1
				}
				//buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, " ", 1))
				//tagLen = len(tagStr) - 1
			}

		} else {
			buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
			tagLen = len(tagStr)
		}

	}
	if curParser.TagRegex().String() == pronounceRegexTag {
		//判断是否为最后一个标记
		if tmpLen < len(ts.textHandle) {
			//判断此位置的字符是不是标点及空格
			if unicode.IsPunct(int32(ts.textHandle[tmpLen])) || ts.textHandle[tmpLen] == ' ' {
				buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
				tagLen = len(tagStr)
			} else {
				buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, " ", 1))
				tagLen = len(tagStr) - 1
			}

		} else {
			buf.WriteString(strings.Replace(handleStr[curLoc[0]:curLoc[0]+handledLength], tagStr, "", 1))
			tagLen = len(tagStr)
		}
	}
	return curLoc[0] + handledLength, tagLen, nil
}

func isPunctuation(ch byte) bool {
	switch ch {
	case '?', '.', ',', ';', '!', ' ':
		return true
	}
	return false
}
func CheckTag(in string) (bool, error) {
	reg, err := regexp.Compile(checkRegexTag)
	if err != nil {
		return false, err
	}
	loc := reg.FindStringSubmatchIndex(in)
	//fmt.Printf("loc=%v,tag=%v \n", loc, in[loc[0]:loc[1]])
	if len(loc) == 0 {
		return false, nil
	}
	return true, nil
}
