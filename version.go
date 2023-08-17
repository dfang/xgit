package main

import (
	"io/ioutil"
	"log"
)

var xgitVersion string
var goVersion string
var buildTimestamp string
var repo string

func init() {
	repo = "https://github.com/dfang/xgit"
	goVersion = "go 1.21.0"
	readVersion()
}

func readVersion() {
	content, err := ioutil.ReadFile("version.txt")
	if err != nil {
		log.Fatal(err)
	}
	xgitVersion = string(content)
}
