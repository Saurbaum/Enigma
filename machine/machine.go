package machine

import (
	"fmt"
	"github.com/saurbaum/enigma/rotor"
)

type Configuration struct {
	Rotors []rotor.Rotor
}

var config Configuration

func Run() {
	fmt.Printf("Starting Machine with %d rotors\n", len(config.Rotors))

	var character byte
	var nextStep = false
	for _, r := range config.Rotors {
		character, nextStep = r.Translate('A', nextStep)
		fmt.Printf("%s", string(character))
		character, nextStep = r.Translate('B', nextStep)
		fmt.Printf("%s", string(character))
	}
}

func Initialise(newConfig Configuration) {
	config = newConfig
}
