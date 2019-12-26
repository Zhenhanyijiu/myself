package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func print() {
	time.Sleep(8 * time.Second)
	fmt.Println("niho")
}
func gort() string {
	go print()
	return "hudi"
}

const (
	innum = 7 + iota
	innum2
)

func main() {
	//fmt.Println("Hello, World!")
	//out:=gort()
	//fmt.Println("out=", gort())
	//time.Sleep(15 * time.Second)

	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	fmt.Println(">>", valueOfA.Type())
	// 获取interface{}类型的值, 通过类型断言转换
	var getA int = valueOfA.Interface().(int)
	// 获取64位的值, 强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, getA2)
	fmt.Println("innum2=", innum2)
	loc := time.FixedZone("UTC+8", 8*3600)
	layout := "2006-01-02 15:04:05"
	tt := time.Time{}
	//tt := time.Now()
	fmt.Println(tt.String())
	t, _ := time.ParseInLocation(layout, tt.Format(layout), loc)
	fmt.Println(t.String())
	fmt.Println(isZeroUTC8(t))
	ss := "tts/play"
	sss := strings.Split(ss, "/")
	fmt.Println(len(sss), sss[0], sss[1])
	tab := map[int]errBody{
		6: errBody{
			6,
			"unknow error",
		},
	}
	ou, ok := tab[3]
	fmt.Println(ou, ok)

}

type errBody struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func isZeroUTC8(t time.Time) bool {
	//ZeroUTC8 := "0001-01-01 00:00:00 +0800 UTC+8"
	return strings.EqualFold(t.String(), "0001-01-01 00:00:00 +0800 UTC+8") //t.String() == ZeroUTC8
}
