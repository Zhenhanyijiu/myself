package test

import (
	"encoding/base64"
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
)

var bs = `iL7m48mbFSIy1Y5xbXWwPTR07ufxu7o+myGUE+AdDeWWISkd5W6Gl44oX/jgXldS` +
	`mL/ntUXoZzQz2WKEYLwssAtSTGF+QgSIMvV5faiP+pLYvWgk0oVr42po00CvADFL` +
	`eDAJC7LgagYifS1l4EAK4MY8RGCHyJWEN5JAr0fc/Haa3WfWZ009kOWAp8MDuYxB` +
	`hQlCKUmnUpXCp5c6jwbjlyinLj8XwzzjZ/rVRsY+t2Z0Vcd5qzR5BV8IJCqbG5Py` +
	`z15/EFgMG2N2eYMsiEKgdXeKW2H5XIoWyun/3pBigWaDnTtiWSt9kz2MplqYfIT7` +
	`F+0XE3gdDGalAeN3YwFPHCkxxBmcI+s6lQG9INmf2/gkJQ+MOZBVXKmGLv6Qis3l` +
	`0eyUz1yZvNzf0zlcUBjiPulLF3peThHMEzhSsATfPomyg5NJ0X7ttd0ybnq+sPe4` +
	`qg2OJ8qNhYrqnx7Xlvj61+B2NAZVHvIioma1FzqX8DxQYrnR5S6DJExDqvzNxEz6` +
	`5VPQlH2Ig4hTvNzla84WgJ6USc/2SS4ehCReiNvfeNG9sPZKQnr/Ss8KPIYsKGcC` +
	`Pz/vEqbWDmJwHb7KixCQKPt1EbD+/uf0YnhskOWM15YiFbYAOZKJ5rcbz2Zu66vg` +
	`GAmqcBsHeFR3s/bObEzjxOmMfSr1vzvr4ActNJWVtfNKZNobSehZiMSHL54AXAZW` +
	`Yj48pwTbf7b1sbF0FeCuwTFiYxM+yiZVO5ciYOfmo4HUg53PjknKpcKtEFSj02P1` +
	`8JRBSb++V0IeMDyZLl12zgURDsvualbJMMBBR8emIpF13h0qdyah431gDhHGBnnC` +
	`J5UDGq21/flFjzz0x/Okjwf7mPK5pcmF+uW7AxtHqws6m93yD5+RFmfZ8cb/8CL8` +
	`jmsQslj+OIE64ykkRoJWpNBKyQjL3CnPnLmAB6TQKxegR94C7/hP1FvRW+W0AgZy` +
	`g2QczKQU3KBQP18Ui1HTbkOUJT0Lsy4FnmJFCB/STPRo6NlJiATKHq/cqHWQUvZd` +
	`d4oTMb1opKfs7AI9wiJBuskpGAECdRnVduml3dT4p//3BiP6K9ImWMSJeFpjFAFs` +
	`AbBMKyitMs0Fyn9AJRPl23TKVQ3cYeSTxus4wLmx5ECSsHRV6g06nYjBp4GWEqSX` +
	`RVclXF3zmy3b1+O5s2chJN6TrypzYSEYXJb1vvQLK0lNXqwxZAFV7Roi6xSG0fSY` +
	`EAtdUifLonu43EkrLh55KEwkXdVV8xneUjh+TF8VgJKMnqDFfeHFdmN53YYh3n3F` +
	`kpYSmVLRzQmLbH9dY+7kqvnsQm8y76vjug3p4IbEbHp/fNGf+gv7KDng1HyCl9A+` +
	`Ow/Hlr0NqCAIhminScbRsZ4SgbRTRgGEYZXvyOtQa/uL6I8t2NR4W7ynispMs0QL` +
	`RD61i3++bQXuTi4i8dg3yqIfe9S22NHSzZY/lAHAmmc3r5NrQ1TM1hsSxXawT5CU` +
	`anWFjbH6YQ/QplkkAqZMpropWn6ZdNDg/+BUjukDs0HZrbdGy846WxQUvE7G2bAw` +
	`IFQ1SymBZBtfnZXhfAXOHoWh017p6HsIkb2xmFrigMj7Jh10VVhdWg==`

var ss = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOiIzMyIsInVzZXIiOiJxeGYifQ.T5HFLi14CKy51_B7PYcDi8Zzje3XmdGAxaE1zOdzOAqtqIe5PcaYZspB2tENZGxhA-BNyzG-6c1DPLRrCGU32qkzk9cxLcpXvftwAl2I9K2YQ_u6XRzgEwMSbn9AaHGJf8kTrGiEiFLFJI-NCrUANC6dpEJwOPziSebkk3cwRd3SATIuQc0eBdzeS5lK7vKBs0wAbQWXmeDgHSxXOTiwrxzLGKrRCcXUd97nD1lYhBfLRPfOO3ro7ptDfQdVExW-E9HGuiQw-uBEEAJS14v3r8OOFGlXFu0xd00f8xEujN00Cq8XpRMfsPtURjT2fn75Kk51sIhyBwYIu6Fo4ZM3Cg`

func TestPrint(t *testing.T) {
	by, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		fmt.Println("======error:", err)
		return
	}

	fmt.Println("======", by, len(by))
}

func TestStringsTrimRight(t *testing.T) {
	s1 := strings.TrimRight("cyeamblog.go", "g.o")
	fmt.Printf("TrimRight:%v\n", s1)
	s2 := strings.TrimSuffix("abcdaaaa", "daaaa")
	fmt.Printf("TrimSuffix:%v\n", s2)
	//s3 := []string{"ni", "hao", "ma", "boy"}
	s3 := []string{"ni", "ni", "ni"}
	s4 := strings.Join(s3, "")
	fmt.Printf("strings.Join: %v\n", s4)
	m := make(map[string]string)
	m["11"] = "aa"
	m["22"] = "bb"
	m["33"] = "cc"
	fmt.Printf("map:%v\n", m)
	var kk []string
	for k := range m {
		kk = append(kk, k)
		fmt.Printf("k:%v\n", k)
	}
	sort.Strings(kk)
	fmt.Printf("kk:%v\n", kk)

	//var resize int
	testSwitch(100)
	testSwitch(150)
	testSwitch(151)
}

func testSwitch(resize int) {
	switch true {
	case resize < 150:
		fmt.Printf("resize<150\n")
	case resize >= 150:
		fmt.Printf("resize>=150\n")
	}
}

func TestFloor(t *testing.T) {
	angle := float64(360)
	angle1 := angle - math.Floor(angle/360)*360
	fmt.Printf("angle:%v, %v\n", angle1, math.Floor(angle/360)*360)
}
