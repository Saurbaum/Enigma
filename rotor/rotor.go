package rotor

type Rotor struct {
	table       string
	stepFirst   byte
	stepSecond  byte
	ringSetting int

	position int
}

var RotorI = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
var RotorII = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
var RotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
var RotorIV = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
var RotorV = "VZBRGITYUPSDNHLXAWMJQOFECK"
var RotorVI = "JPGVOUMFYQBENHZRDKASXLICTW"
var RotorVII = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
var RotorVIII = "FKQHTLXOCBJSPDZRAMEWNIUYGV"

func CreateI(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'Q', 0, ringSetting, 0}
	return rotor
}

func CreateII(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'E', 0, ringSetting, 0}
	return rotor
}

func CreateIII(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'V', 0, ringSetting, 0}
	return rotor
}

func CreateIV(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'J', 0, ringSetting, 0}
	return rotor
}

func CreateV(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'Z', 0, ringSetting, 0}
	return rotor
}

func CreateVI(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'Z', 'M', ringSetting, 0}
	return rotor
}

func CreateVII(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'Z', 'M', ringSetting, 0}
	return rotor
}

func CreateVIII(ringSetting int) Rotor {
	rotor := Rotor{RotorI, 'Z', 'M', ringSetting, 0}
	return rotor
}

func (r Rotor) GetPosition() int {
	return r.position
}

func (r Rotor) testStep(stepPoint byte) bool {
	if stepPoint != 0 && byte(r.position+'A') == stepPoint {
		return true
	}

	return false
}

func (r Rotor) Translate(input byte, step bool) (byte, bool) {
	var pushStep = false

	if step {
		if r.testStep(r.stepFirst) || r.testStep(r.stepSecond) {
			pushStep = true
		}
		r.position++
	}

	var convertIndex = int(input-'A') + r.position

	convertIndex = convertIndex % len(r.table)

	var character = byte(r.table[convertIndex])

	return character, pushStep
}
