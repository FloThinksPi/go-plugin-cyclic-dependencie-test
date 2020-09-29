package main

import (
	"fmt"
	pluginexample "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/cmd/plugin_example/functions"
	"github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/four/interface"
	three "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/three/interface"
	"time"
)

func GetName() string {
	return "four"
}

func TestFunc()  {
	fmt.Println("four: testFunc called")
	time.Sleep(1 * time.Second)
	threeInterface := pluginexample.GetPluginInterface("three").(three.ThreeInterface)
	threeInterface.MyFunc()
}

func GetInterface() interface{} {
	return _interface.FourInterface{
		VariableBla: "test",
		MyFunction: TestFunc,
	}
}