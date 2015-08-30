package machine

import (
	"fmt"
	"github.com/saurbaum/enigma/rotor"
)

type Configuration struct {
	Rotors    []rotor.Rotor
	Reflector rotor.Rotor
}

var config Configuration

func Run(input string) {
	fmt.Printf("Starting Machine with %d rotors\n", len(config.Rotors))

	for _, c := range input {
		var character byte = byte(c)

		fmt.Printf("Transforming %s\n", string(character))

		var nextStep bool = true
		for index := 0; index < len(config.Rotors); index++ {
			character, nextStep = config.Rotors[index].Translate(character, nextStep)
			fmt.Printf("Rotor %d translate to %s\n", index, string(character))

			index = index + 1
		}
	}
}

func Initialise(newConfig Configuration) {
	config = newConfig
}
