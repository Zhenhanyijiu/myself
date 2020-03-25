package test

import (
	"fmt"
	"testing"
)
import "github.com/tidwall/gjson"

var jsontext = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func TestGjson(t *testing.T) {
	res := gjson.Get(jsontext, "children.#")
	fmt.Printf("json:%v\n", res.String())
	res = gjson.Get(jsontext, "children")
	for i, v := range res.Array() {
		fmt.Printf("i:%v, v:%v\n", i, v)
	}
}
