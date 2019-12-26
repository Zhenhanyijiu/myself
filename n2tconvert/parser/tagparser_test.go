package parser

import (
	"encoding/json"
	"fmt"
	"myself/n2tconvert/dc"
	"testing"
)

func TestSoundLinkingParser_InitData(t *testing.T) {
	var linkParser SoundLinkingParser
	linkParser.InitData()
	if linkParser.tagRegex.String() != soundlinkingRegexTag {
		t.Errorf("InitData error")
		return
	}
}
func TestSoundLinkingParser_TagRegex(t *testing.T) {
	var linkParser SoundLinkingParser
	linkParser.InitData()
	out := linkParser.TagRegex()
	if (out == nil) || (out.String() != soundlinkingRegexTag) {
		t.Errorf("TagRegex error")
		return
	}
}
func TestSoundLinkingParser_GetMatchTag(t *testing.T) {
	var linkParser SoundLinkingParser
	linkParser.InitData()
	inputString := "far[c:]away"
	//loc := []int{0, 11, 3, 7}
	loc := linkParser.TagRegex().FindStringSubmatchIndex(inputString)
	//fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out, err := linkParser.GetMatchTag(inputString, loc)
	if err != nil {
		t.Errorf("GetMatchTag error\n")
		return
	}
	if out != inputString[loc[2]:loc[3]] {
		t.Errorf("GetMatchTag output error\n")
		return
	}
}

//func testCheckMarker(marker dc.TagMarker,loc []int) error {
//	if
//	return nil
//}
func TestSoundLinkingParser_SetMarker(t *testing.T) {
	var linkParser SoundLinkingParser
	linkParser.InitData()
	inputString := " far[c:]away"
	//loc := []int{0, 11, 3, 7}
	loc := linkParser.TagRegex().FindStringSubmatchIndex(inputString)
	fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	marker := dc.TagMarker{}
	err := linkParser.SetMarker(inputString, 0, 0, loc, &marker)
	if err != nil {
		t.Errorf("SetMarker error\n")
		return
	}
	func() {
		if marker.Type != "Linking" || marker.Position.Start != loc[0]+0-0 || marker.Position.Length != loc[1]-loc[0]-3 {
			t.Errorf("SetMarker error\n")
			return
		}
	}()

}

func TestSoundLinkingParser_GetHandledLength(t *testing.T) {
	var linkParser SoundLinkingParser
	linkParser.InitData()
	inputString := "far[c:]away"
	//loc := []int{0, 11, 3, 7}
	loc := linkParser.TagRegex().FindStringSubmatchIndex(inputString)
	fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out := linkParser.GetHandledLength(loc)
	if out != loc[3]-loc[0] {
		t.Errorf("GetHandledLength error\n")
	}
}

func TestToneParser_InitData(t *testing.T) {
	var toneParser ToneParser
	toneParser.InitData()
	if toneParser.tagRegex.String() != toneRegexTag {
		t.Errorf("InitData error")
		return
	}
}
func TestToneParser_TagRegex(t *testing.T) {
	var toneParser ToneParser
	toneParser.InitData()
	if toneParser.TagRegex().String() != toneRegexTag {
		t.Errorf("InitData error")
		return
	}
}

func TestToneParser_GetMatchTag(t *testing.T) {
	var toneParser ToneParser
	toneParser.InitData()
	inputString := "far away[t:r]"
	//loc := []int{4, 13, 4, 8, 8, 13}
	loc := toneParser.TagRegex().FindStringSubmatchIndex(inputString)
	//fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out, err := toneParser.GetMatchTag(inputString, loc)
	if err != nil {
		t.Errorf("GetMatchTag error\n")
		return
	}
	if out != inputString[loc[4]:loc[5]] {
		t.Errorf("GetMatchTag error\n")
		return
	}

}

func TestToneParser_SetMarker(t *testing.T) {
	var toneParser ToneParser
	toneParser.InitData()
	inputString := "far away[t:r]"
	//loc := []int{4, 13, 4, 8, 8, 13}
	loc := toneParser.TagRegex().FindStringSubmatchIndex(inputString)
	fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	var marker dc.TagMarker
	err := toneParser.SetMarker(inputString, 0, 0, loc, &marker)
	if err != nil {
		t.Errorf("SetMarker error\n")
		return
	}
	func() {
		if marker.Type != "Tone" || marker.Position.Start != loc[0]+0-0 || marker.Position.Length != loc[3]-loc[0] {
			t.Errorf("SetMarker error\n")
			return
		}
	}()
}
func TestToneParser_GetHandledLength(t *testing.T) {
	var toneParser ToneParser
	toneParser.InitData()
	inputString := "far away[t:r]"
	//loc := []int{4, 13, 4, 8, 8, 13}
	loc := toneParser.TagRegex().FindStringSubmatchIndex(inputString)
	//fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out := toneParser.GetHandledLength(loc)
	if out != loc[1]-loc[0] {
		t.Errorf("GetHandledLength error\n")
		return
	}
}

func TestStressParser_InitData(t *testing.T) {
	var stressParser StressParser
	stressParser.InitData()
	if stressParser.tagRegex.String() != stressRegexTag {
		t.Errorf("InitData error\n")
		return
	}
}
func TestStressParser_TagRegex(t *testing.T) {
	var stressParser StressParser
	stressParser.InitData()
	if stressParser.TagRegex().String() != stressRegexTag {
		t.Errorf("TagRegex error")
		return
	}
}
func TestStressParser_GetMatchTag(t *testing.T) {
	var stressParser StressParser
	stressParser.InitData()
	inputString := "far[s：]away"
	//loc := []int{0, 7, 3, 7}
	loc := stressParser.TagRegex().FindStringSubmatchIndex(inputString)
	//fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out, err := stressParser.GetMatchTag(inputString, loc)
	if err != nil {
		t.Errorf("GetMatchTag error\n")
		return
	}
	if out != inputString[loc[2]:loc[3]] {
		t.Errorf("GetMatchTag error\n")
		return
	}
}
func TestStressParser_SetMarker(t *testing.T) {
	var stressParser StressParser
	stressParser.InitData()
	inputString := " far[s:]away"
	//loc := []int{0, 7, 3, 7}
	loc := stressParser.TagRegex().FindStringSubmatchIndex(inputString)
	fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	var marker dc.TagMarker
	err := stressParser.SetMarker(inputString, 0, 0, loc, &marker)
	if err != nil {
		t.Errorf("SetMarker error\n")
		return
	}
	func() {
		if marker.Type != "SentenceStress" || marker.Position.Start != loc[0]+0-0 || marker.Position.Length != loc[2]-loc[0] {
			t.Errorf("SetMarker error\n")
			return
		}
	}()
}
func TestStressParser_GetHandledLength(t *testing.T) {
	var stressParser StressParser
	stressParser.InitData()
	inputString := "far[s:]away"
	//loc := []int{0, 7, 3, 7}
	loc := stressParser.TagRegex().FindStringSubmatchIndex(inputString)
	//fmt.Println(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	out := stressParser.GetHandledLength(loc)
	if out != loc[1]-loc[0] {
		t.Errorf("GetHandledLength error\n")
		return
	}
}

func TestPronouncePaeser_InitData(t *testing.T) {
	var pronounceParser PronouncePaeser
	if err := pronounceParser.InitData(); err != nil {
		t.Errorf("initData error(%v)\n", err)
		return
	}
	if pronounceParser.tagRegex.String() != pronounceRegexTag {
		t.Errorf("initData error\n")
		return
	}
	t.Logf("ok ok ...%v\n", pronounceParser.tagRegex.String())
}

func TestPronouncePaeser_TagRegex(t *testing.T) {
	var pronounceParser PronouncePaeser
	if err := pronounceParser.InitData(); err != nil {
		t.Errorf("initData error(%v)\n", err)
		return
	}
	if pronounceParser.TagRegex().String() != pronounceRegexTag {
		t.Errorf("initData error\n")
		return
	}
	t.Logf("ok ok ...%v\n", pronounceParser.TagRegex().String())
}

func TestPronouncePaeser_GetMatchTag(t *testing.T) {
	var pronounceParse PronouncePaeser
	if err := pronounceParse.InitData(); err != nil {
		t.Errorf("initData error(%v)\n", err)
		return
	}
	input := "hi hello [ps:a a a] boy!"
	loc := pronounceParse.TagRegex().FindStringSubmatchIndex(input)
	t.Logf("loc=%v\n", loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	tag, err := pronounceParse.GetMatchTag(input, loc)
	if err != nil {
		t.Errorf("GetMatchTag error(%v)\n", err)
		return
	}
	t.Logf("Tag=%v\n", tag)
	if tag != input[loc[4]:loc[5]] {
		t.Errorf("GetMatchTag error(%v)\n", err)
		return
	}
}
func TestPronouncePaeser_SetMarker(t *testing.T) {
	var pronounceParse PronouncePaeser
	if err := pronounceParse.InitData(); err != nil {
		t.Errorf("initData error(%v)\n", err)
		return
	}
	input := "hello [ps:a a a] boy!"
	loc := pronounceParse.TagRegex().FindStringSubmatchIndex(input)
	t.Logf("loc=%v\n", loc)
	marker := &dc.TagMarker{}
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	err := pronounceParse.SetMarker(input, 0, 0, loc, marker)
	if err != nil {
		t.Errorf("SetMarker error\n")
		return
	}
	func() {
		if marker.Type != "phone" || marker.Position.Start != loc[0]+0-0 || marker.Position.Length != loc[3]-loc[2] {
			t.Errorf("SetMarker error\n")
			return
		}
	}()
	out, _ := json.MarshalIndent(&marker, "", "  ")
	t.Logf("%v\n", string(out))
}

func TestPronouncePaeser_GetHandledLength(t *testing.T) {
	var pronounceParse PronouncePaeser
	if err := pronounceParse.InitData(); err != nil {
		t.Errorf("initData error(%v)\n", err)
		return
	}
	input := "say hello[p:a a a] boy!"
	loc := pronounceParse.TagRegex().FindStringSubmatchIndex(input)
	t.Logf("loc=%v\n", loc)
	out := pronounceParse.GetHandledLength(loc)
	if len(loc) == 0 {
		t.Logf("no tag \n")
		return
	}
	if out != loc[1]-loc[0] {
		t.Errorf("GetHandledLength error)\n")
		return
	}

}
