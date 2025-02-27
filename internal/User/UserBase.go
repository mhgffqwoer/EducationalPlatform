package user

import (
	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
)

var generator *generatorid.GeneratorID

func init() {
	generator = generatorid.New()
}

type UserBase interface {
	GetName() string
	GetID() *generatorid.ID
}

type UserBuilder interface {
	SetName(name string) UserBuilder
	Build() (UserBase, error)
}

type UserFactory interface {
	Teacher() UserBuilder
	Student() UserBuilder
}

type UserFactoryImpl struct{}

func (u *UserFactoryImpl) Teacher() UserBuilder {
	return &UserTeacherBuilder{generator: generator}
}

func (u *UserFactoryImpl) Student() UserBuilder {
	return &UserStudentBuilder{generator: generator}
}

func New() UserFactory {
	return &UserFactoryImpl{}
}
