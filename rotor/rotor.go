package rotor

type Rotor struct {
	table       string
	stepFirst   byte
	stepSecond  byte
	ringSetting int

	position int

	static bool
}

var rotorI = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
var rotorII = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
var rotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
var rotorIV = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
var rotorV = "VZBRGITYUPSDNHLXAWMJQOFECK"
var rotorVI = "JPGVOUMFYQBENHZRDKASXLICTW"
var rotorVII = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
var rotorVIII = "FKQHTLXOCBJSPDZRAMEWNIUYGV"

var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"

func createRotor(tableData string, firstStepPos byte, secondStepPos byte, ringSettingPos int, pos int) Rotor {
	return Rotor{table: tableData, stepFirst: firstStepPos, stepSecond: secondStepPos, ringSetting: ringSettingPos, position: pos, static: false}
}

func createReflector(tableData string, pos int) Rotor {
	return Rotor{table: tableData, position: pos, static: true}
}

func CreateI(ringSetting int) Rotor {
	return createRotor(rotorI, 'Q', 0, ringSetting, 0)
}

func CreateII(ringSetting int) Rotor {
	return createRotor(rotorII, 'E', 0, ringSetting, 0)
}

func CreateIII(ringSetting int) Rotor {
	return createRotor(rotorIII, 'V', 0, ringSetting, 0)
}

func CreateIV(ringSetting int) Rotor {
	return createRotor(rotorIV, 'J', 0, ringSetting, 0)
}

func CreateV(ringSetting int) Rotor {
	return createRotor(rotorV, 'Z', 0, ringSetting, 0)
}

func CreateVI(ringSetting int) Rotor {
	return createRotor(rotorVI, 'Z', 'M', ringSetting, 0)
}

func CreateVII(ringSetting int) Rotor {
	return createRotor(rotorVII, 'Z', 'M', ringSetting, 0)
}

func CreateVIII(ringSetting int) Rotor {
	return createRotor(rotorVIII, 'Z', 'M', ringSetting, 0)
}

func CreateReflectorB(position int) Rotor {
	return createReflector(reflectorB, position)
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

func (r *Rotor) Translate(input byte, step bool) (byte, bool) {
	var pushStep = false

	if step && !r.static {
		r.position++
	}

	if r.testStep(r.stepFirst) || r.testStep(r.stepSecond) {
		pushStep = true
	}

	var convertIndex = int(input-'A') + r.position

	convertIndex = convertIndex % len(r.table)

	var character = byte(byte(r.table[convertIndex]) - byte(r.position)%byte(len(r.table)))

	return character, pushStep
}

func (r Rotor) ReverseTranslate(input byte) byte {
	input = ((input-'A')+byte(r.position))%byte(len(r.table)) + 'A'

	for index := 0; index < len(r.table); index++ {
		if byte(r.table[index]) == input {
			return 'A' + byte((index-r.position)%len(r.table))
		}
	}

	return 0
}
