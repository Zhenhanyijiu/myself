package parser

import (
	"errors"
	"fmt"
	"myself/n2tconvert/dc"
	"regexp"
	"strings"
)

const (
	// 连读, 文本+可能的多个空格+[c:]标记+可能的多个空格+文本
	//aa[c:]bb, aa  [c:] bb, aa-bb[c:]dd, a's[c:]bb
	soundlinkingRegexTag = `(?:[-':\w]+\s*(\[c[:：]\])\s*[-':\w]+)`
	// 升降调, 文本+常见标点符号+可能的多个空格+[t:r]标记
	//aa[t:r], aa [t:r], aa?[t:r],aa[t:r]?
	toneRegexTag = `(?:([-':\w]+)[^[]?\s*(\[t[:：][rf]\]))`
	//重读, 文本+常见标点符号+可能的多个空格+[t:r]标记
	//aa[s:]bb, aa[s:] bb, aa [s:] bb, aa?[s:],aa[s:]?
	stressRegexTag = `(?:([-':\w]+)[^[]?\s*(\[s[:：]\]))`
	//音标标签，hell[p:hə'ləʊ]
	//pronounceRegexTag = `(?:([-':\w]+)[^[]?\s*(\[p[:：][\w]+\]))`
	//所有音标
	soundMarkAll      = `[ ɑːæʌɔːaʊəaɪbtʃdðeəeɜːeɪfghɪəɪiːi:dʒklmnŋɒəʊɔɪprsʃtθʊəʊuːvwjzʒtrtsdzdsˈˌ'ɛ]+`
	pronounceRegexTag = `(?:([-':\w]+)[^[]?\s*(\[(p|ps)[:：]` + soundMarkAll + `\]))`
)

//用于检查是否存在这几种标记的正则表达式
const checkRegexTag = `(?:\[c[:：]\]|\[t[:：][rf]\]|\[s[:：]\]|\[(p|ps)[:：]` + soundMarkAll + `\])`

// TagParser 用于一个具体的tag处理
type TagParser interface {
	InitData() error
	TagRegex() *regexp.Regexp
	GetMatchTag(handleStr string, loc []int) (string, error)
	SetMarker(handleStr string, diff int, tagLenInAll int, loc []int, marker *dc.TagMarker) error
	GetHandledLength(loc []int) int
}

// SoundLinkingParser 用于一个SoundLinking的处理
type SoundLinkingParser struct {
	tagRegex *regexp.Regexp
}

// TagRegex 用来返回正则表达式
func (s *SoundLinkingParser) TagRegex() *regexp.Regexp {
	return s.tagRegex
}

// InitData 表示初始化数据成员
func (s *SoundLinkingParser) InitData() error {
	var err error
	s.tagRegex, err = regexp.Compile(soundlinkingRegexTag)
	return err
}

// GetMatchTag 用来获取匹配的tag信息
func (s *SoundLinkingParser) GetMatchTag(handleStr string, loc []int) (string, error) {
	if len(loc) != 4 {
		return "", fmt.Errorf("SoundLinkingParser loc size is not logical, len(loc): %d", len(loc))
	}
	return handleStr[loc[2]:loc[3]], nil
}

// SetMarker 用来设置marker中的值
func (s *SoundLinkingParser) SetMarker(handleStr string, diff int, tagLenInAll int, loc []int, marker *dc.TagMarker) error {
	if len(loc) < 4 {
		return fmt.Errorf("SoundLinkingParser loc size is not logical, len(loc): %d", len(loc))
	}
	marker.Position.Start = loc[0] + diff - tagLenInAll
	//marker.Position.Length = loc[1] - loc[0] - 4 + 1
	marker.Position.Length = loc[1] - loc[3] + loc[2] - loc[0] + 1
	marker.Type = "Linking"
	return nil
}

// GetHandledLength 用来获得已处理的字段长度
func (s *SoundLinkingParser) GetHandledLength(loc []int) int {
	if len(loc) < 4 {
		return 0
	}
	return (loc[3] - loc[0])
}

// ToneParser 用于一个tone的处理
type ToneParser struct {
	tagRegex *regexp.Regexp
}

// TagRegex 用来返回正则表达式
func (t *ToneParser) TagRegex() *regexp.Regexp {
	return t.tagRegex
}

// InitData 表示初始化数据成员
func (t *ToneParser) InitData() error {
	var err error
	t.tagRegex, err = regexp.Compile(toneRegexTag)
	return err
}

// GetMatchTag 用来获取匹配的tag信息
func (t *ToneParser) GetMatchTag(handleStr string, loc []int) (string, error) {
	if len(loc) != 6 {
		return "", fmt.Errorf("ToneParser loc size is not logical, len(loc): %d", len(loc))
	}
	tagStr := handleStr[loc[4]:loc[5]]
	if tagStr != "[t:r]" && tagStr != "[t:f]" && tagStr != "[t：r]" && tagStr != "[t：f]" {
		return "", fmt.Errorf("ToneParser tagStr is not logical, tagStr: %s", tagStr)
	}

	return tagStr, nil
}

// SetMarker 用来设置marker中的值
func (t *ToneParser) SetMarker(handleStr string, diff int, tagLenInAll int, loc []int, marker *dc.TagMarker) error {
	if len(loc) != 6 {
		return fmt.Errorf("ToneParser loc size is not logical, len(loc): %d", len(loc))
	}
	marker.Position.Start = loc[0] + diff - tagLenInAll
	marker.Position.Length = loc[3] - loc[2]
	marker.Type = "Tone"
	tagStr := handleStr[loc[4]:loc[5]]
	if tagStr == "[t:r]" || tagStr == "[t：r]" { // 1 升调
		marker.Value = 1
	} else if tagStr == "[t:f]" || tagStr == "[t：f]" { // 3 降调
		marker.Value = 3
	} else {
		return fmt.Errorf("ToneParser tagStr is not logical, tagStr: %s", tagStr)
	}
	return nil
}

// GetHandledLength 用来获取已处理的字段长度
func (t *ToneParser) GetHandledLength(loc []int) int {
	if len(loc) != 6 {
		return 0
	}
	return (loc[1] - loc[0])
}

type StressParser struct {
	tagRegex *regexp.Regexp
}

func (s *StressParser) TagRegex() *regexp.Regexp {
	return s.tagRegex
}
func (s *StressParser) InitData() error {
	var err error
	s.tagRegex, err = regexp.Compile(stressRegexTag)
	return err
}
func (s *StressParser) GetMatchTag(handleStr string, loc []int) (string, error) {
	if len(loc) != 6 {
		return "", fmt.Errorf("StressParser loc size is not logical, len(loc): %d", len(loc))
	}
	return handleStr[loc[4]:loc[5]], nil
}

func (s *StressParser) SetMarker(handleStr string, diff int, tagLenInAll int, loc []int, marker *dc.TagMarker) error {
	if len(loc) != 6 {
		return fmt.Errorf("StressParser loc size is not logical, len(loc): %d", len(loc))
	}
	marker.Position.Start = loc[0] + diff - tagLenInAll
	marker.Position.Length = loc[3] - loc[2]
	marker.Type = "SentenceStress"
	return nil
}
func (s *StressParser) GetHandledLength(loc []int) int {
	if len(loc) != 6 {
		return 0
	}
	return (loc[1] - loc[0])
}

//音标标签
type PronouncePaeser struct {
	tagRegex *regexp.Regexp
}

func (p *PronouncePaeser) InitData() error {
	var err error
	p.tagRegex, err = regexp.Compile(pronounceRegexTag)
	return err
}

func (p *PronouncePaeser) TagRegex() *regexp.Regexp {
	return p.tagRegex
}

func (p *PronouncePaeser) GetMatchTag(handleStr string, loc []int) (string, error) {
	if len(loc) != 8 {
		return "", fmt.Errorf("PronounceParser loc size is not logical, len(loc): %d", len(loc))
	}
	//[p: a a a] [ps: a a a] [ps： a a a] [ps： a a a] 支持中文冒号
	tagStr := handleStr[loc[4]:loc[5]]
	return tagStr, nil
}

func (p *PronouncePaeser) SetMarker(handleStr string, diff int, tagLenInAll int, loc []int, marker *dc.TagMarker) error {
	if len(loc) != 8 {
		return fmt.Errorf("PronounceParser loc size is not logical, len(loc): %d", len(loc))
	}
	marker.Type = "phone"
	marker.Position.Start = loc[0] + diff - tagLenInAll
	marker.Position.Length = loc[3] - loc[2]
	tagStr := strings.Replace(handleStr[loc[4]:loc[5]], "：", ":", -1)
	//fmt.Printf(">>>>>%v,%v\n", tagStr, len(tagStr))
	tagP := handleStr[loc[4]:loc[7]]
	tagLen := len(tagStr)
	ValueStr := []string{}
	//fmt.Printf("tagP:=%v\n", tagP)
	soundMark := ""
	switch tagP {
	case "[p":
		soundMark = tagStr[3 : tagLen-1]
	case "[ps":
		soundMark = tagStr[4 : tagLen-1]
	default:
		return errors.New("pronounce tag parse error")
	}
	ValueStr = append(ValueStr, strings.Replace(soundMark, " ", "·", -1)) //·
	marker.Value = ValueStr
	return nil
}
func (p *PronouncePaeser) GetHandledLength(loc []int) int {
	if len(loc) != 8 {
		return 0
	}
	return (loc[1] - loc[0])
}
