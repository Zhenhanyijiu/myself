package main

import (
	"fmt"
	"github.com/eleniums/async"
	"log"
	"time"
)

func main() {
	foo := func() error {
		// do something
		time.Sleep(time.Second * 5)
		return nil
	}

	bar := func() error {
		// do something else
		fmt.Printf("bar ok\n")
		return nil
	}
	errc := async.Run(foo, bar)
	err := async.Wait(errc)
	if err != nil {
		log.Fatalf("task returned an error: %v", err)
	}
	fmt.Printf("===end")
}
