package main

import (
	"os"

	"github.com/iamlongalong/easyget"
)

func main() {
	os.Setenv("LONG_name", "longalong")
	os.Setenv("LONG_age", "18")
	os.Setenv("LONG_hobbies", "cooking,eating")

	g := easyget.NewEnvGetter("LONG_")
	v, ok := g.Get("name")
	if !ok {
		panic("get env name should be ok")
	}

	if v != "longalong" {
		panic("get env name should be longalong")
	}

	fg := easyget.NewJSONGetterFromJSONFile("tests/files/test.json")

	// cg := easyget.NewJSONGetterFromHTTP()

	gg := easyget.NewSelectGetter(fg)

	v, ok = gg.Get("name")
	if !ok {
		panic("get env name should be ok")
	}

	if v != "longalong" {
		panic("get env name should be longalong")
	}
}
