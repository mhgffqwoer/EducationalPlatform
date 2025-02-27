package user

import generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"

type UserStudent struct {
	name string
	id   *generatorid.ID
}

func (u *UserStudent) GetName() string {
	return u.name
}

func (u *UserStudent) GetID() *generatorid.ID {
	return u.id
}

type UserStudentBuilder struct {
	name      string
	generator *generatorid.GeneratorID
}

func (b *UserStudentBuilder) SetName(name string) UserBuilder {
	b.name = name
	return b
}

func (b *UserStudentBuilder) Build() (UserBase, error) {
	return &UserStudent{name: b.name, id: generator.GenerateID()}, nil
}
