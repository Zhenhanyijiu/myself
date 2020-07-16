package main

import (
	"fmt"
	"github.com/go-ego/gse"
)

func main() {
	text := "小明硕士毕业于中国科学院计算所后在日本京都大学深造"
	text = "南京市长"
	newSeg := gse.New()
	fmt.Printf("精准模式\n\n")
	res := newSeg.Cut(text, true)
	fmt.Printf("hmm(true )===res:%+v\n", res)
	res = newSeg.Cut(text, false)
	fmt.Printf("hmm(false)===res:%+v\n", res)
	fmt.Printf("搜索模式\n\n")
	res = newSeg.CutSearch(text, true)
	fmt.Printf("hmm(true )===res:%+v\n", res)
	res = newSeg.CutSearch(text, false)
	fmt.Printf("hmm(false)===res:%+v\n", res)
	fmt.Printf("全模式\n\n")
	res = newSeg.CutAll(text)
	fmt.Printf("hmm(true )===res:%+v\n", res)
	//res = newSeg.CutAll(text)
	//fmt.Printf("hmm(false)===res:%+v\n", res)
	seg := gse.Segmenter{}
	seg.LoadDict("/D_QAwYqH3be/yangyanan/myself/src/github.txt.com/go-ego/gse/data/dict/dictionary.txt")
	//seg.LoadDict()
	resString := seg.String(text, false)
	fmt.Printf("resString:%v\n\n", resString)
	resSlice := seg.Slice(text, false)
	fmt.Printf("resString seachMode(false):%+v\n\n", resSlice)
	resSlice = seg.Slice(text, true)
	fmt.Printf("resString seachMode(true ):%+v\n\n", resSlice)

	text = "山达尔星联邦共和国"
	segments := seg.Segment([]byte(text))
	//fmt.Printf("===segments:%+v\n\n", segments)
	s1 := gse.ToString(segments, false)
	fmt.Printf("gseToString seachMode(false):%+v\n\n", s1)
	s1 = gse.ToString(segments, true)
	fmt.Printf("gseToString seachMode(true):%+v\n\n", s1)
}
