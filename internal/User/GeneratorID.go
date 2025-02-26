package user

type generatorID struct {
	id int
}

func (g *generatorID) generateID() int {
	g.id++
	return g.id
}
