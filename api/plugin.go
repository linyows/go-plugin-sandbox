package api

import "io"

var NoiseFormat1 string = "%s < %s\n"
var NoiseFormat2 string = "%s (%s)\n"

type Conf struct {
	Stdout io.Writer
}

type Sound interface {
	Name() string
	Noise() string
	MakeNoise() (int, error)
}

type Sounds interface {
	Init(Conf)
	List() map[string]Sound
}
