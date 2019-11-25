package main

import (
	"fmt"
	"github.com/guonaihong/flag"
	"io"
	"io/ioutil"
	"os"
)

func testInputFromStdin() {
	var inn io.Reader
	inn = os.Stdin
	out1, err := ioutil.ReadAll(inn)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}
	fmt.Printf("%v\n", string(out1))
}

type FlagStruct struct {
	Name string `opt:"n,string" usage:"(must)name of test"`
	Age  int    `opt:"a,int" usage:"(must)age of test"`
}

func testFlagStruct() {
	flagStr := FlagStruct{}
	flag.ParseStruct(&flagStr)
	if flagStr.Name == "" || flagStr.Age == 0 {
		flag.Usage()
		return
	}
	fmt.Printf("%v\n", flagStr)
}
func main() {
	//testInputFromStdin()
	testFlagStruct()
}
