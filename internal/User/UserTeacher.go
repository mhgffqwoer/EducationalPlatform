package user

import generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"

type UserTeacher struct {
	name string
	id   *generatorid.ID
}

func (u *UserTeacher) GetName() string {
	return u.name
}

func (u *UserTeacher) GetID() *generatorid.ID {
	return u.id
}

type UserTeacherBuilder struct {
	name string
}

func (b *UserTeacherBuilder) SetName(name string) UserBuilder {
	b.name = name
	return b
}

func (b *UserTeacherBuilder) Build() (UserBase, error) {
	return &UserTeacher{name: b.name, id: generator.GenerateID()}, nil
}
