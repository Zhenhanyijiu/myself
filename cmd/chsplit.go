package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"github.com/go-ego/re/log"
)

func main1() {
	gse.New()

	var seg gse.Segmenter
	seg.LoadDict("zh,testdata/test_dict.txt,testdata/test_dict1.txt")

	text1 := "你好世界, Hello world"
	fmt.Println(seg.String(text1, true))

	segments := seg.Segment([]byte(text1))
	fmt.Println(gse.ToString(segments))
	log.Now("fffff")
}

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	newSeg = gse.New("zh,/D_QAwYqH3be/yangyanan/gopath/pkg/mod/github.txt.com/go-ego/gse@v0.60.0-rc3.0.20200710152819-7763312f64e8/testdata/test_dict3.txt", "alpha")

	//text = "你好世界, Hello world, Helloworld."
	text = "小明硕士毕业于中国科学院计算所后在日本京都大学深造"
)

func cut() {

	fmt.Printf("\n=======newSegmenter:%+v, \n\n", newSeg)
	hmm := newSeg.Cut(text, false)
	fmt.Println("cut use hmm: ===", hmm)

	hmm = newSeg.CutSearch(text, true)
	fmt.Println("cut search use hmm: ===", hmm)

	hmm = newSeg.CutAll(text)
	fmt.Println("cut all:=== ", hmm)
}

func main() {
	cut()

	segCut()
}

func posAndTrim(cut []string) {
	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	posSeg.WithGse(seg)
	po := posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}

func cutPos() {
	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))

	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}

func segCut() {
	// 加载默认字典
	//seg.LoadDict()
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.txt.com/go-ego/gse/data/dict/dictionary.txt")
	seg.LoadDict("/D_QAwYqH3be/yangyanan/myself/src/github.txt.com/go-ego/gse/data/dict/dictionary.txt")

	// 分词文本
	//tb := "山达尔星联邦共和国联邦政府"
	//tb := "央广网合肥7月14日消息（记者王利）14日晚，安徽省防汛抗旱指挥部发出紧急命令，要求安庆、池州、铜陵、芜湖、马鞍山市立即做好长江江心洲和外滩圩人员撤离工作。当前，安徽省长江干流全线持续超警戒水位，预计未来1-2天沿江各主要控制站将陆续出现洪峰水位，部分江心洲、外滩圩将面临更加严峻的洪水威胁。命令要求，相关市防指要对照有人居住的江心洲、外滩圩转移预警表，预报达到预警条件的必须立即组织开展人员撤离工作，并尽快完成全部撤离任务，确保撤离不漏一户、不落一人。"
	tb := "央广网合肥7月14日消息（记者王利）14日晚安徽省防汛抗旱指挥部发出紧急命令" +
		"要求安庆池州铜陵芜湖马鞍山市立即做好长江江心洲和外滩圩人员撤离工作" +
		"当前安徽省长江干流全线持续超警戒水位预计未来1-2天沿江各主要控制站将陆续" +
		"出现洪峰水位部分江心洲外滩圩将面临更加严峻的洪水威胁命令要求" +
		"相关市防指要对照有人居住的江心洲外滩圩转移预警表预报达到预警条件的" +
		"必须立即组织开展人员撤离工作并尽快完成全部撤离任务确保撤离不漏一户不落一人"
	tb = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(tb, true))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb))

	segments := seg.Segment([]byte(tb))
	//fmt.Printf("=========%+v,%+v\n\n", segments, segments[0].Token())
	// 处理分词结果
	fmt.Println("处理分词结果：", gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	fmt.Println("处理分词结果：", gse.ToString(segments1, true))
}
