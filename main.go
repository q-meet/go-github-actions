package main

import (
	"fmt"

	"go-github-actions/hello"
)

var version = "0.0.1"

func GetVersion() string {
	return version
}

func main() {
	fmt.Println(hello.Greet())
	fmt.Println(GetVersion())
}

func w1() string {
	return hello.Greet()
}
