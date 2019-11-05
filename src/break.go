package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func init() {
	fmt.Println("enter init")
}
func main() {
	num := 0
end:
	for ; num < 10; num++ {
		fmt.Printf("%v\n", num)
		switch num {
		case 3:
			break end
		}
	}
	f, _ := os.Open("fis")
	bytes.NewReader()
	strconv.ParseFloat()
	fmt.Printf("num=%v\n", num)
}
