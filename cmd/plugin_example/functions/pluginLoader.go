package functions

import (
	"os"
	"path/filepath"
	"plugin"
)

type PluginObject struct {
	Plugin        *plugin.Plugin
	PluginFile    os.File
	TestFunc func()
	Interface interface{}
}

var loadedPlugins = make(map[string]PluginObject)

func LoadPlugins(){

	var files []string

	root, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/plugins/")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		var p PluginObject
		p.Plugin, err = plugin.Open(file)
		if err != nil {
			panic(err)
		}

		s, err := p.Plugin.Lookup("TestFunc")
		if err != nil {
			panic(err)
		}
		p.TestFunc = s.(func())

		n, err := p.Plugin.Lookup("GetName")
		if err != nil {
			panic(err)
		}
		NameFunc := n.(func() string)

		//Optional
		i, err := p.Plugin.Lookup("GetInterface")
		if err == nil {
			InterfaceFunc := i.(func() interface{})
			p.Interface = InterfaceFunc()
		}

		loadedPlugins[NameFunc()] = p
	}

	println("Calling TestFunc of first Plugin")
	loadedPlugins[os.Args[1]].TestFunc()

}

func GetPluginInterface(pluginName string) interface{} {
	return loadedPlugins[pluginName].Interface
}
