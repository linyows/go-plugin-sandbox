package main

import (
	"fmt"
	"io"

	"github.com/linyows/go-plugin-sandbox/api"
)

func init() {
	fmt.Println("animal plugin init!")
}

type dog string

func (d dog) Name() string  { return string(d) }
func (d dog) Noise() string { return "bowwow!!!!!!" }
func (d dog) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, d.Noise(), d.Name())
}

type cat string

func (c cat) Name() string  { return string(c) }
func (c cat) Noise() string { return "meow~ mewo~" }
func (c cat) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, c.Noise(), c.Name())
}

type cow string

func (w cow) Name() string  { return string(w) }
func (w cow) Noise() string { return "mu~u~~~~~~~!" }
func (w cow) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, w.Noise(), w.Name())
}

type Animals struct {
	Stdout io.Writer
}

func (a *Animals) Init(c api.Conf) {
	a.Stdout = c.Stdout
}

func (a *Animals) List() map[string]api.Sound {
	return map[string]api.Sound{
		"dog": dog("dog"),
		"cat": cat("cat"),
		"cow": cow("cow"),
	}
}

var Sounds Animals
