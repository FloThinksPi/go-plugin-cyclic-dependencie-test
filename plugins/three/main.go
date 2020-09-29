package main

import (
	"fmt"
	pluginexample "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/cmd/plugin_example/functions"
	four "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/four/interface"
	_interface "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/three/interface"
	"time"
)

func GetName() string {
	return "three"
}

func TestFunc()  {
	fmt.Println("three: testFunc called")
	time.Sleep(1 * time.Second)
	fourInterface := pluginexample.GetPluginInterface("four").(four.FourInterface)
	fourInterface.MyFunction()
}

func GetInterface() interface{} {
	return _interface.ThreeInterface{
		Variable1: "bla",
		MyFunc:    TestFunc,
	}
}