package machine

import (
	"fmt"
	"github.com/saurbaum/enigma/rotor"
)

var rotors []rotor.Rotor

func Run() {
	fmt.Printf("Starting Machine with %d rotors\n", len(rotors))

	for _, r := range rotors {
		fmt.Println(r.GetPosition())
	}
}

func Initialise(wheelCount int) {
	for i := 0; i < wheelCount; i++ {
		rotors = append(rotors, rotor.Create(i))
	}
}
