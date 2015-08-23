package rotor

type Rotor struct {
	offset int
}

func Create(newOffset int) Rotor {
	rotor := Rotor{newOffset}
	return rotor
}

func (r Rotor) GetPosition() int {
	return r.offset
}
