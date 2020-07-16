package main

import (
	"fmt"
	"io"

	//"github.txt.com/cao-guang/pinyin"

	"github.com/Chain-Zhang/pinyin"
	cgp "github.com/cao-guang/pinyin"
	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
	"github.com/go-ego/gse"
)

var test = `西雅图都会区; 长夜漫漫, winter is coming!`

func main1() {
	args := gpy.Args{
		Style:     gpy.Tone,
		Heteronym: true}

	py := gpy.Pinyin(test, args)
	fmt.Println("===gpy:", py)

	s := gpy.ToString(py)
	fmt.Println("===gpy string:", s)

	phrase.LoadGseDict()
	go func() {
		fmt.Println("===gpy phrase1:", phrase.Paragraph(test))
	}()
	fmt.Println("===gpy phrase2:", phrase.Paragraph(test))

	seg := gse.New("zh, /D_QAwYqH3be/yangyanan/gopath/pkg/mod/github.txt.com/go-ego/gpy@v0.31.3/examples/dict.txt")
	// phrase.DictAdd["都会区"] = "dū huì qū"
	phrase.AddDict("都会区", "dū huì qū")

	fmt.Println("===gpy phrase:", phrase.Paragraph(test, seg))
	fmt.Println("===pinyin: ", phrase.Pinyin(test))
	fmt.Println("===Initial: ", phrase.Initial("都会区"))
}

func main() {
	hans := "中国话"
	hans = "南京市长江大桥"
	// 默认
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong] [guo] [hua]]

	// 包含声调
	a.Style = gpy.Tone
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhōng] [guó] [huà]]

	// 声调用数字表示
	a.Style = gpy.Tone2
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng] [guo2] [hua4]]

	// 开启多音字模式
	fmt.Printf("============\n\n")
	a = gpy.NewArgs()
	a.Heteronym = true
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zhong zhong] [guo] [hua]]
	a.Style = gpy.Tone3
	fmt.Println(gpy.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [hua4]]
	res := gpy.HanPinyin(hans, a)
	fmt.Printf("=======hanpinyin:%v\n\n", res)
	//fmt.Println(gpy.LazyPinyin(hans, a))
	//fmt.Println(gpy.LazyPinyin(hans, gpy.NewArgs()))
	// [zhong guo hua]

	fmt.Println(gpy.Convert(hans, nil))
	// [[zhong] [guo] [hua]]
	fmt.Println(gpy.LazyConvert(hans, nil))
	// [zhong guo hua]
	//hans1 := "南京市长江大桥"
	//a1 := gpy.NewArgs()
	//a1.Fallback = func(r rune, a gpy.Args) []string {
	//	return []string{string(r + 1)}
	//}
	//fmt.Println(gpy.HanPinyin(hans1, a1))
}
func main3() {
	hans := "中国话アイウ"
	hans = "中国话"
	a := gpy.NewArgs()
	a.Fallback = func(r rune, a gpy.Args) []string {
		data := map[rune][]string{
			'ア': {"a"},
			'イ': {"i"},
			'ウ': {"u"},
		}
		s, ok := data[r]
		if ok {
			return s
		} else {
			return []string{}
		}
	}
	fmt.Println(gpy.HanPinyin(hans, a))
	s := gpy.ToString([][]string{{"zhong", "zong"}, {"guo"}, {"hua"}})
	fmt.Printf("======s:%v\n\n", s)
	//phrase.DictAdd
	a = gpy.NewArgs()
	a.Heteronym = true
	a.Style = gpy.Tone3
	fmt.Printf("=======singlepinyin:%v\n", gpy.SinglePinyin(rune('中'), a))

}

func main4() {
	cgp.LoadingPYFileName("/D_QAwYqH3be/yangyanan/gopath/pkg/mod/github.txt.com/cao-guang/pinyin@v0.0.0-20190927081849-38872b67965d/pinyin.txt") //这里是字典文件路径程序启动调用一次，载入缓存
	//demo
	str1, err := cgp.To_Py("汉字拼音", "", "") //默认造型： hanzipinyin
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str1)
	str2, err := cgp.To_Py("汉字拼音", "", cgp.Tone) //带声调：hànzìpīnyīn
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str2)
	str3, err := cgp.To_Py("汉字拼音", "", cgp.InitialsInCapitals) //首字母大写无声调：HanZiPinYin
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str3)
	str4, err := cgp.To_Py("汉字拼音", "-", cgp.InitialsInCapitals) //首字母大写无声调加-分割：Han-Zi-Pin-Yin
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str4)

}

func main5() {
	str, err := pinyin.New("我是中国人").Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str)
	}

	str, err = pinyin.New("我是中国人").Split(" ").Mode(pinyin.WithoutTone).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str)
	}

	str, err = pinyin.New("我是中国人").Split("-").Mode(pinyin.Tone).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str)
	}

	str, err = pinyin.New("我是中国人").Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str)
	}
}

type Pwr struct {
	pw *io.PipeWriter
	pr *io.PipeReader
}

func main6() {
	pwr := Pwr{}
	pwr.pr, pwr.pw = io.Pipe()
}
