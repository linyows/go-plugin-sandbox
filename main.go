package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"

	"github.com/linyows/go-plugin-sandbox/api"
)

type Toy struct {
	sounds map[string]api.Sound
}

func main() {
	toy := &Toy{
		sounds: make(map[string]api.Sound),
	}

	fmt.Fprintln(os.Stdout, "\nPlugins:")
	if err := toy.LoadSounds(); err != nil {
		fmt.Errorf("load error:\n", err)
		os.Exit(1)
	}

	count := len(toy.sounds)
	if count < 1 {
		fmt.Errorf("sounds plugin not found\n")
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "\nNoise1:")
	toy.TurnItOut()
	fmt.Fprintln(os.Stdout, "\nNoise2:")
	toy.MakeNoise()
}

func (t *Toy) LoadSounds() error {
	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.Mode().IsRegular() {
			continue
		}
		n := f.Name()

		if filepath.Ext(n) != ".so" {
			continue
		}

		fmt.Println("./plugins/" + n)
		p, err := plugin.Open("./plugins/" + n)
		if err != nil {
			return err
		}

		soundsSymbol, err := p.Lookup("Sounds")
		if err != nil {
			return err
		}

		sounds := soundsSymbol.(api.Sounds)
		c := api.Conf{
			Stdout: os.Stdout,
		}
		sounds.Init(c)

		for name, f := range sounds.List() {
			t.sounds[name] = f
			fmt.Printf("loaded: %s\n", f.Name())
		}
	}

	return nil
}

func (t *Toy) TurnItOut() {
	for name, f := range t.sounds {
		noise := f.Noise()
		fmt.Printf(api.NoiseFormat1, name, noise)
	}
}

func (t *Toy) MakeNoise() {
	for _, f := range t.sounds {
		f.MakeNoise()
	}
}
