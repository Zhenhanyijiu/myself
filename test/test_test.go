package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
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
	parallel(0, 5, mm)
}

func mm(ys <-chan int) {
	for y := range ys {
		fmt.Printf("===%v\n", y)
		//i := y * dst.Stride
		//src.scan(0, y, src.w, y+1, dst.Pix[i:i+size])
		for i := 0; i < 3; i++ {
			fmt.Printf("///%v\n", i)
		}
	}
}

func parallel(start, stop int, f func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(c)
		}()
	}
	wg.Wait()
}

const (
	aa1 uint32 = 2 * (iota)
	aa2
	aa3
	aa4
	aa5
)

func TestIota(t *testing.T) {
	fmt.Printf("aa1(%v)\naa2(%v)\naa3(%v)\naa4(%v)\naa5(%v)\n", aa1, aa2, aa3, aa4, aa5)
}

func TestBase64(t *testing.T) {
	//in := "123/45678909876543211c"
	//http://wiki.it.yzs.io:8090/display/edu/faceID-v2-doc#faceID-v2-doc
	//in :=[]byte{}
	in := make([]byte, 256)
	for i := 0; i < 256; i++ {
		in[i] = byte(i)
	}
	out := base64.StdEncoding.EncodeToString([]byte(in))
	fmt.Printf("base64 std:\n%v\n", out)

	out = base64.URLEncoding.EncodeToString([]byte(in))
	fmt.Printf("base64 Url:\n%v\n", out)
}

func TestSlice(t *testing.T) {
	in := []byte{1, 2, 3, 4, 5, 6, 7}
	in1 := in[3:7:7]
	fmt.Printf("in1:=%v\n", in1)
	mp := make(map[int]string) //{"11": "a", "22": "b", "33": "c", "44": "a", "55": "b", "66": "c"}
	for i := 0; i < 7; i++ {
		mp[i] = strconv.Itoa(i)
	}
	for k, v := range mp {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	fmt.Printf("%v\n", mp)
	out, _ := json.Marshal(mp)
	fmt.Printf("out:%v\n", string(out))
	fmt.Printf("second:%v\n", time.Now().Unix())
	fmt.Printf("ns:%v\n", time.Now().UnixNano())
	tm := time.Unix(1578461949, 0)
	fmt.Printf("time:%v, %v \n", tm.String(), time.Now().String())
	mp1 := make(map[int]string)
	mp1[0] = "88"
	fmt.Printf("===%v\n", mp1[1])
}

type Infomation struct {
	Name string
	Age  string
}
type Person struct {
	Info string
	Num  int
}

func TestJson(t *testing.T) {
	info := Infomation{Name: "qq", Age: "13"}
	out, _ := json.Marshal(&info)
	p := Person{Info: string(out), Num: 1000}
	out, _ = json.Marshal(&p)
	fmt.Printf("===%v\n", len(out))
	sh := sha256.New()
	sh.Write(out)
	st := sh.Sum(nil)
	In := big.NewInt(0)
	In = In.SetBytes(st)
	he := hex.EncodeToString(st)
	fmt.Printf("====%v,,%v,,%v,,%v,,%v\n", string(out), In.Text(16), len(st), he, len(he))
	info2 := Infomation{}
	err := json.Unmarshal([]byte(p.Info), &info2)
	if err != nil {
		fmt.Printf("error (%v)\n", err)
		return
	}
	fmt.Printf("====%v\n", info2)
}
func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func TestCrypt(t *testing.T) {
	key := []byte{0x93, 0xf4, 0x56, 0x4b, 0x12, 0x2c, 0x95, 0x47,
		0x5c, 0x44, 0x10, 0x2e, 0xac, 0xb6, 0xb9, 0x13}
	iv := []byte{0x39, 0xf4, 0x56, 0x4b, 0x12, 0x2c, 0x95, 0x47,
		0x5c, 0x44, 0x10, 0x2e, 0xac, 0xb6, 0xb9, 0x14}
	ci, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	blockSize := ci.BlockSize()
	src := []byte{0x33, 0xf4, 0x56, 0x4b, 0x12, 0x2c, 0x95, 0x47,
		0x5c, 0x44, 0x10, 0x2e, 0xac, 0xb6, 0xb9, 0x11}
	fmt.Printf("blockSize:%v\n", blockSize)
	src = padding(src, 16)
	md := cipher.NewCBCEncrypter(ci, iv)
	dst := make([]byte, len(src))
	md.CryptBlocks(dst, src)
	fmt.Printf("cipher:%x, %v\n", dst, len(dst))
	//dec
	mdd := cipher.NewCBCDecrypter(ci, iv)
	mdd.CryptBlocks(dst, dst)
	dst = unpadding(dst)
	fmt.Printf("plain :%x, %v\n", dst, len(dst))
	//cipher.NewCBCEncrypter()
	/////////////////
	src1 := src[:16]
	ci.Encrypt(src1, src1)
	fmt.Printf("cipher:%x, %v\n", src1, len(src1))
	ci.Decrypt(src, src)
	fmt.Printf("plain :%x, %v\n", src1, len(src1))
}
