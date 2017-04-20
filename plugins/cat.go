package main

import (
	"fmt"
)

var Nick string

type Noise struct {
	Type  string
	Sound string
}

func init() {
	fmt.Println("init!")
}

func Default() *Noise {
	return &Noise{
		Type:  "cat",
		Sound: "meowmeow",
	}
}

func MakeNoise() {
	n := Default()
	fmt.Println(n.Type + " < " + n.Sound + "!!!!!! (" + Nick)
}
