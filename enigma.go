package main

import (
	"github.com/saurbaum/enigma/machine"
	"github.com/saurbaum/enigma/rotor"
)

func main() {
	var rotors []rotor.Rotor

	rotors = append(rotors, rotor.CreateI(0))

	var config = machine.Configuration{rotors, rotor.CreateReflectorB(0)}
	machine.Initialise(config)
	machine.Run("TEST")
}
