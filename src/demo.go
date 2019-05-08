package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "12345667"
	s := str[4:]
	s = strconv.Itoa(9)

	n, _ := strconv.Atoi("11")
	fmt.Println(s, n)
	fmt.Println(s, n)
	fmt.Println(s, n)

	tst := "a  a"
	out := strings.TrimPrefix(tst, "  ")
	fmt.Println(">>", out, len(tst), len(out))
	n = 16
	out = strconv.FormatInt(16, 16)
	fmt.Println("out:", out)

}

//func GetCureve(curveType ...interface{}) string {
//	switch curveType {
//	case "P256":
//		return "P256"
//	case "P224":
//		return "P224"
//	case "P384":
//		return "P384"
//	case "P521":
//		return "P384"
//	default:
//		return "P256"
//	}
//}
