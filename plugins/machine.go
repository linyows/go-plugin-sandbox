package main

import (
	"fmt"
	"io"

	"github.com/linyows/go-plugin-sandbox/api"
)

func init() {
	fmt.Println("machine plugin init!")
}

type mac string

func (m mac) Name() string  { return string(m) }
func (m mac) Noise() string { return "jyaaaaaaaan" }
func (m mac) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, m.Noise(), m.Name())
}

type win string

func (w win) Name() string  { return string(w) }
func (w win) Noise() string { return "tete------n" }
func (w win) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, w.Noise(), w.Name())
}

type ubuntu string

func (u ubuntu) Name() string  { return string(u) }
func (u ubuntu) Noise() string { return "pecopon...." }
func (u ubuntu) MakeNoise() (int, error) {
	return fmt.Fprintf(Sounds.Stdout, api.NoiseFormat2, u.Noise(), u.Name())
}

type Machines struct {
	Stdout io.Writer
}

func (m *Machines) Init(c api.Conf) {
	m.Stdout = c.Stdout
}

func (m *Machines) List() map[string]api.Sound {
	return map[string]api.Sound{
		"mac":    mac("macos"),
		"win":    win("windows"),
		"ubuntu": ubuntu("ubuntu"),
	}
}

var Sounds Machines
