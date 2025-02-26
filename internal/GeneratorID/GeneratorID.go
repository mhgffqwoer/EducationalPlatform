package generatorid

type ID int

func (i *ID) Int() int {
	return int(*i)
}

type GeneratorID struct {
	id int
}

func New() *GeneratorID {
	return &GeneratorID{id: 0}
}

func (g *GeneratorID) GenerateID() *ID {
	g.id++
	result := ID(g.id)
	return &result
}
