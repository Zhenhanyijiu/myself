package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Active string

func (a *Active) Check() bool {
	return *a == "disable"
}

type Action string

func (a *Action) Check() bool {
	return *a == "drop"
}

type rule struct {
	Active
	Action
}

type Checker interface {
	Check() bool
}

func runRule(x interface{}) error {

	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("%T type fail", v.Kind())
	}

	v = v.Elem()

	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {

		f := v.Field(i)

		if ck, ok := f.Addr().Interface().(Checker); ok {
			if ck.Check() {
				fmt.Printf("check fail\n")
				return errors.New("check fail")
			}
		}
	}
	return nil
}

func main() {

	runRule(&rule{Active: "enable", Action: "drop"})
	//runRule(&rule{Active: "enable", Action: "drop"})
}
