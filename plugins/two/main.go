package main

import "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/two/functions"

func GetName() string {
	return "one"
}

func TestFunc()  {
	functions.TestFunc()
}