package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"plugin"
)

func main() {
	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.Mode().IsRegular() {
			continue
		}
		n := f.Name()

		if filepath.Ext(n) == ".so" {
			//plugin, err := loadPlugin(n)
			//if err == nil {
			//execPlugin(plugin)
			fmt.Println("./plugins/" + n)
			p, err := plugin.Open("./plugins/" + n)
			if err != nil {
				panic(err)
			}
			nickSymbol, err := p.Lookup("Nick")
			if err != nil {
				panic(err)
			}
			noiseSymbol, err := p.Lookup("MakeNoise")
			if err != nil {
				panic(err)
			}
			*nickSymbol.(*string) = n
			//fmt.Printf("new: %#v\n", noise)
			//fmt.Printf("make: %#v\n", make.(fnc())())
			makeNoise := noiseSymbol.(func())
			makeNoise()
			//}
		}
	}
}

//func loadPlugin(f string) (*Plugin, err) {
//	p, err := plugin.Open("./plugins/" + f)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return p, err
//}
//
//func execPlugin(p *Plugin) err {
//	name, err := p.Lookup("Name")
//
//	if err == nil {
//		fmt.Println(name)
//	}
//
//	return err
//}
