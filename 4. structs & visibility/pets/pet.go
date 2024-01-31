package pets

type NoiseMaker interface {
	MakeNoise()
}

type Walk interface {
	Walk() error
}

// Every pet has some attributes that are not unique to a certain pet
type Pet struct {
	name string
}

func (p Pet) GetName() string {
	return p.name
}
