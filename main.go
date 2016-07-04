package main

import (
	"strings"
	"log"
	"os"
	"io/ioutil"
	"gopkg.in/flosch/pongo2.v3"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)
	templ, err := pongo2.FromString(string(data))
	checkErr(err)
	ctx := pongo2.Context{}
	for _, env := range os.Environ() {
		pos := strings.Index(env, "=")
		if pos > -1 {
			ctx[env[:pos]] = env[pos+1:]
		}
	}
	err = templ.ExecuteWriter(ctx, os.Stdout)
	checkErr(err)
}

func checkErr(err interface{}) {
	if err != nil {
		log.Fatalf("err: %s", err)
	}
}