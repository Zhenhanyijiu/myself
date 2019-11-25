package main

import "C"
import "fmt"

//export print
func print(f *C.char) {
	fmt.Println("nihao:" + C.GoString(f))
}

func parseJ() {

}

func main() {
}
