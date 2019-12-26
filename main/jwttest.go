package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type MapInfo map[string]string

func (m MapInfo) String() string {
	out, _ := json.Marshal(m)
	return string(out)
}

func (m MapInfo) Set(arg string) error {
	s := strings.SplitN(arg, "=", 2)
	if len(s) != 2 {
		return fmt.Errorf("error:len(arg)=%v, %v", len(s), "len(arg)!=2")
	}
	m[s[0]] = s[1]
	return nil
}
func main() {
	mapInfo := make(MapInfo)
	flag.Var(mapInfo, "map", "map info ")
	flag.Parse()

	fmt.Printf("++++++++++,%v\n", mapInfo)
	out, err := loadData("go.mod")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Printf("out len:%v\n", len(out))
}
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else if p == "+" {
		return []byte("{}"), nil
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}
