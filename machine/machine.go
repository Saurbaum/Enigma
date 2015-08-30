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

	fmt.Println(input)

	var output string

	for _, c := range input {
		var character byte = byte(c)

		fmt.Printf("Transforming %s\n", string(character))

		var nextStep bool = true
		for index := 0; index < len(config.Rotors); index++ {
			character, nextStep = config.Rotors[index].Translate(character, nextStep)
			fmt.Printf("Rotor %d translate to %s\n", index, string(character))
		}

		character, _ = config.Reflector.Translate(character, false)
		fmt.Printf("Reflector translate to %s\n", string(character))

		for i := len(config.Rotors) - 1; i >= 0; i-- {
			character = config.Rotors[i].ReverseTranslate(character)
			fmt.Printf("Rotor %d reverse translate to %s\n", i, string(character))
		}

		output += string(character)
	}

	fmt.Println(output)
}

func Initialise(newConfig Configuration) {
	config = newConfig
}
