package main

import "io"

type Pwr struct {
	pw *io.PipeWriter
	pr *io.PipeReader
}

func main() {
	pwr := Pwr{}
	pwr.pr, pwr.pw = io.Pipe()
}
