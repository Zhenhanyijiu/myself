package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//bufio.SplitFunc()
	//strings.FieldsFunc()
	var i64 int64
	i64 = 16
	newI64 := atomic.AddInt64(&i64, 1)
	fmt.Printf("oldI64=%x, newI64=%x\n", i64, newI64)
}
