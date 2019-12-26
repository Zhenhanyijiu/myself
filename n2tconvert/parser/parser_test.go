package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"myself/n2tconvert/dc"
	"testing"
)

//soundlinkingRegexTag = `(?:\w+\s*(\[c:\])\s*\w+)`
const (
	linkInputString      = "far[c:]away"
	toneRiseInputString  = "far away[t:r]"
	toneFallInputString  = "far away[t:f]"
	stressInputString    = "far[s:]away"
	pronounceInputString = "a [ps:æ] a [ps:æ] opportunity [ps：ˌ ɒ p ə ˈ t juː n ə t i] she's [p:ʃ i: z] record [p:ˈ r e k ə r d] family[p：ˈ f æ m ə l ɪ] boy[p:b ɔ ɪ] "
)
const (
	link = iota + 1
	toneRise
	toneFall
	stress
	pronounce
)

func TestNewTextScan(t *testing.T) {
	textScan, err := NewTextScan()
	if err != nil {
		t.Errorf("run NewTextScan error\n")
		return
	}
	for _, v := range textScan.tagParsers {
		if v == nil {
			t.Errorf("init tagParser failed\n")
			return
		}
	}
}

func (ts *TextScan) testTextScan_RunOne(input string, startIndex, fg int) (bool, error) {
	var res dc.Result
	var buf bytes.Buffer
	switch fg {
	case link:
		fmt.Println("link RunOne test start")
		oneHandledLength, _, _ := ts.RunOne(input[startIndex:], startIndex, 0, &buf, &res)
		curloc := make([]int, 4)
		for _, v := range ts.tagParsers {
			loc := v.TagRegex().FindStringSubmatchIndex(input[startIndex:])
			if len(loc) == 0 {
				continue
			}
			curloc = loc
			//fmt.Println(">>", loc)
		}
		//fmt.Println("cur", curloc, "buf=", buf.String(), len(buf.String()))
		if oneHandledLength != curloc[3] {
			fmt.Println(oneHandledLength)
			return false, errors.New("not equal")
		}
	case toneRise:
		fmt.Println("toneRise RunOne test start")
		oneHandledLength, _, _ := ts.RunOne(input[startIndex:], startIndex, 0, &buf, &res)
		curloc := make([]int, 6)
		for _, v := range ts.tagParsers {
			loc := v.TagRegex().FindStringSubmatchIndex(input[startIndex:])
			if len(loc) == 0 {
				continue
			}
			curloc = loc
			//fmt.Println(">>", loc)
		}
		//fmt.Println("cur", curloc, "buf=", buf.String(), len(buf.String()))
		if oneHandledLength != curloc[5] {
			//fmt.Println(oneHandledLength)
			return false, errors.New("not equal")
		}
	case toneFall:
		fmt.Println("toneFall RunOne test start")
		oneHandledLength, _, _ := ts.RunOne(input[startIndex:], startIndex, 0, &buf, &res)
		curloc := make([]int, 6)
		for _, v := range ts.tagParsers {
			loc := v.TagRegex().FindStringSubmatchIndex(input[startIndex:])
			if len(loc) == 0 {
				continue
			}
			curloc = loc
			//fmt.Println(">>", loc)
		}
		//fmt.Println("cur", curloc, "buf=", buf.String(), len(buf.String()))
		if oneHandledLength != curloc[5] {
			fmt.Println(oneHandledLength)
			return false, errors.New("not equal")
		}
	case stress:
		fmt.Println("stress RunOne test start")
		oneHandledLength, _, _ := ts.RunOne(input[startIndex:], startIndex, 0, &buf, &res)
		curloc := make([]int, 6)
		for _, v := range ts.tagParsers {
			loc := v.TagRegex().FindStringSubmatchIndex(input[startIndex:])
			if len(loc) == 0 {
				continue
			}
			curloc = loc
			///fmt.Println(">>", loc)
		}
		//fmt.Println("cur", curloc, "buf=", buf.String(), len(buf.String()))
		if oneHandledLength != curloc[1] {
			fmt.Println(oneHandledLength)
			return false, errors.New("not equal")
		}
	case pronounce:
		fmt.Println("pronounce RunOne test start")
		oneHandledLength, _, _ := ts.RunOne(input[startIndex:], startIndex, 0, &buf, &res)
		curloc := make([]int, 6)
		for _, v := range ts.tagParsers {
			loc := v.TagRegex().FindStringSubmatchIndex(input[startIndex:])
			if len(loc) == 0 {
				continue
			}
			curloc = loc
			//fmt.Println(">>", loc)
		}
		//fmt.Println("cur", curloc, "buf=", buf.String(), len(buf.String()))
		if oneHandledLength != curloc[1] {
			fmt.Println(oneHandledLength)
			return false, errors.New("not equal")
		}
	default:
		return false, errors.New("fg param error")
	}
	return true, nil
}

func TestTextScan_RunOne(t *testing.T) {
	ts, err := NewTextScan()
	if err != nil {
		t.Errorf("NewTextScan \n")
		return
	}
	bl, err := ts.testTextScan_RunOne(linkInputString, 1, link)
	if err != nil || bl == false {
		fmt.Println(err, bl)
		t.Errorf("link RunOne error\n")
	}
	bl, err = ts.testTextScan_RunOne(toneRiseInputString, 0, toneRise)
	if err != nil || bl == false {
		fmt.Println(err, bl)
		t.Errorf("toneRise RunOne error\n")
	}
	bl, err = ts.testTextScan_RunOne(toneFallInputString, 2, toneFall)
	if err != nil || bl == false {
		fmt.Println(err, bl)
		t.Errorf("toneFall RunOne error\n")
	}
	bl, err = ts.testTextScan_RunOne(stressInputString, 0, stress)
	if err != nil || bl == false {
		fmt.Println(err, bl)
		t.Errorf("stress RunOne error\n")
	}

	bl, err = ts.testTextScan_RunOne(pronounceInputString, 0, pronounce)
	if err != nil || bl == false {
		fmt.Println(err, bl)
		t.Errorf("pronounce RunOne error\n")
	}
}

func check(res1, res2 dc.Result) bool {
	if res1.Version != res2.Version {
		return false
	}
	if res1.DisplayText != res2.DisplayText || len(res1.Markers) != len(res2.Markers) {
		return false
	}
	for i, v := range res1.Markers {
		if v.Type != res2.Markers[i].Type || v.Value != res2.Markers[i].Value {
			return false
		}
		if v.Position.Start != res2.Markers[i].Position.Start || v.Position.Length != res2.Markers[i].Position.Length {
			return false
		}
	}
	return true
}
func testCheckResult(result dc.Result, fg int) (bool, error) {
	linkRes := dc.Result{
		Version:     1,
		DisplayText: "far away",
		Markers: []dc.TagMarker{
			dc.TagMarker{
				Position: dc.TagPosition{Start: 0, Length: 8},
				Type:     "Linking",
			},
		},
	}

	toneRiseRes := dc.Result{
		Version:     1,
		DisplayText: "far away",
		Markers: []dc.TagMarker{
			dc.TagMarker{
				Position: dc.TagPosition{Start: 4, Length: 4},
				Type:     "Tone",
				Value:    1,
			},
		},
	}

	toneFallRes := dc.Result{
		Version:     1,
		DisplayText: "far away",
		Markers: []dc.TagMarker{
			dc.TagMarker{
				Position: dc.TagPosition{Start: 4, Length: 4},
				Type:     "Tone",
				Value:    3,
			},
		},
	}

	stressRes := dc.Result{
		Version:     1,
		DisplayText: "far away",
		Markers: []dc.TagMarker{
			dc.TagMarker{
				Position: dc.TagPosition{Start: 0, Length: 3},
				Type:     "SentenceStress",
			},
		},
	}
	switch fg {
	case link:
		return check(linkRes, result), nil
	case toneRise:
		return check(toneRiseRes, result), nil
	case toneFall:
		return check(toneFallRes, result), nil
	case stress:
		return check(stressRes, result), nil
	default:
		return false, errors.New("fg param error")
	}
	return true, nil
}

func TestTextScan_RunScan(t *testing.T) {
	ts, err := NewTextScan()
	if err != nil {
		t.Errorf("NewTextScan \n")
		return
	}

	//linking test
	res, err := ts.RunScan(linkInputString)
	if err != nil {
		t.Errorf("linkInputString  RunScan error\n")
		return
	}
	bl, _ := testCheckResult(res, link)
	if !bl {
		t.Errorf("the result is not equal\n")
	}
	out, _ := json.MarshalIndent(&res, "", "  ")
	t.Logf("link(%v)\n", string(out))

	//tone rise test
	res, err = ts.RunScan(toneRiseInputString)
	if err != nil {
		t.Errorf("toneRiseInputString  RunScan error\n")
		return
	}

	bl, _ = testCheckResult(res, toneRise)
	if !bl {
		t.Errorf("the result is not equal\n")
	}
	out, _ = json.MarshalIndent(&res, "", "  ")
	t.Logf("toneRise(%v)\n", string(out))

	//tone fall test
	res, err = ts.RunScan(toneFallInputString)
	if err != nil {
		t.Errorf("toneFallInputString  RunScan error\n")
		return
	}

	bl, _ = testCheckResult(res, toneFall)
	if !bl {
		t.Errorf("the result is not equal\n")
	}
	out, _ = json.MarshalIndent(&res, "", "  ")
	t.Logf("toneFall(%v)\n", string(out))

	//stress test
	res, err = ts.RunScan(stressInputString)
	if err != nil {
		t.Errorf("stressInputString  RunScan error\n")
		return
	}
	bl, _ = testCheckResult(res, stress)
	if !bl {
		t.Errorf("the result is not equal\n")
	}
	out, _ = json.MarshalIndent(&res, "", "  ")
	t.Logf("stress(%v)\n", string(out))

	//pronounce mark test
	res, err = ts.RunScan(pronounceInputString)
	if err != nil {
		t.Errorf("stressInputString  RunScan error\n")
		return
	}
	out, _ = json.MarshalIndent(&res, "", "  ")
	t.Logf("pronounce(%v)\n", string(out))
}

type ResultTmp struct {
	Version     int
	DisplayText string
	Markers     []TagMarker
}

type TagPosition struct {
	Start  int
	Length int
}
type TagMarker struct {
	Position TagPosition
	Type     string
	Value    interface{} `json:"Value,omitempty"`
}

func TestTextScan_RunScan_Pronounce(t *testing.T) {
	//pronounceInputString := "record[p:ˈ r ɛ k ə r d]family[p：ˈ f æ m ə l ɪ]boy[p：b ɔ ɪ]."
	ts, err := NewTextScan()
	if err != nil {
		t.Errorf("NewTextScan \n")
		return
	}
	res, err := ts.RunScan(pronounceInputString)
	if err != nil {
		t.Errorf("pronounceInputString  RunScan error\n")
		return
	}
	out, _ := json.MarshalIndent(&res, "", "  ")
	t.Logf("%v\n", string(out))
	var tmp ResultTmp
	json.Unmarshal(out, &tmp)
	out1, _ := json.Marshal(&tmp)
	t.Logf("%v\n", string(out1))

	ch := "·"
	t.Logf("==%v,%x, %v\n", len(ch), []byte(ch), ch)
}

func (ts *TextScan) benchTextScan_RunScan(b *testing.B, input string, fg int) {
	b.ResetTimer()
	b.ReportAllocs()
	b.ResetTimer()
	switch fg {
	case link:
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ts.RunScan(input)
			}
		})
	case toneRise:
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ts.RunScan(input)
			}
		})
	case toneFall:
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ts.RunScan(input)
			}
		})
	case stress:
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ts.RunScan(input)
			}
		})
	case pronounce:
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ts.RunScan(input)
			}
		})
	default:
		return
	}

}
func BenchmarkTextScan_RunScanLink(b *testing.B) {
	ts, _ := NewTextScan()
	ts.benchTextScan_RunScan(b, linkInputString, link)

}
func BenchmarkTextScan_RunScanToneRise(b *testing.B) {
	ts, _ := NewTextScan()
	ts.benchTextScan_RunScan(b, toneRiseInputString, toneRise)
}
func BenchmarkTextScan_RunScanToneFall(b *testing.B) {
	ts, _ := NewTextScan()
	ts.benchTextScan_RunScan(b, toneFallInputString, toneFall)
}
func BenchmarkTextScan_RunScanStress(b *testing.B) {
	ts, _ := NewTextScan()
	ts.benchTextScan_RunScan(b, stressInputString, stress)
}
func BenchmarkTextScan_RunScanPronounce(b *testing.B) {
	ts, _ := NewTextScan()
	ts.benchTextScan_RunScan(b, pronounceInputString, stress)
}

func TestCheckTag(t *testing.T) {
	fg, err := CheckTag(linkInputString)
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	if fg != true {
		t.Errorf("fg(%v)\n", fg)
		return
	}
	t.Logf("fg(%v)\n", fg)

	fg, err = CheckTag(toneRiseInputString)
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	if fg != true {
		t.Errorf("fg(%v)\n", fg)
		return
	}
	t.Logf("fg(%v)\n", fg)

	fg, err = CheckTag(toneFallInputString)
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	if fg != true {
		t.Errorf("fg(%v)\n", fg)
		return
	}
	t.Logf("fg(%v)\n", fg)

	fg, err = CheckTag(stressInputString)
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	if fg != true {
		t.Errorf("fg(%v)\n", fg)
		return
	}
	t.Logf("fg(%v)\n", fg)

	fg, err = CheckTag(pronounceInputString)
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	if fg != true {
		t.Errorf("fg(%v)\n", fg)
		return
	}
	t.Logf("fg(%v)\n", fg)
}
