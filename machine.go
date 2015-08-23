package machine

import (
	"fmt"
	"github.com/saurbaum/enigma/rotor"
)

func Run() {
	fmt.Println("test")
	rotorA := rotor.Create(0)
	fmt.Println(rotorA.GetPosition())
}
