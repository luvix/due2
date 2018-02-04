package core

type Timely struct {
	name string
}

func (t *Timely) Name() string {
	return t.name
}

func NewTimely(name string) *Timely {
	t := &Timely{
		name: name,
	}

	return t
}
