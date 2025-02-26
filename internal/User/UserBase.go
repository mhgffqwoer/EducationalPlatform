package user

import (
	"errors"

	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
)

var generator *generatorid.GeneratorID

func init() {
	generator = generatorid.New()
}

type UserType int

const (
	Student UserType = iota
	Teacher
)

type UserBase interface {
	GetUserType() UserType
	GetName() string
	GetID() *generatorid.ID
}

type UserBuilder interface {
	SetName(name string) UserBuilder
	Build() (UserBase, error)
}

func New(userType UserType) (UserBuilder, error) {
	switch userType {
	case Student:
		return &UserStudentBuilder{generator: generator}, nil
	case Teacher:
		return &UserTeacherBuilder{generator: generator}, nil
	default:
		return nil, errors.New(`invalid user type`)
	}
}
