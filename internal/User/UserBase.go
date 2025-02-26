package user

import "errors"

type UserType int

const (
	Student UserType = iota
	Teacher
)

type UserBase interface {
	GetUserType() UserType
	GetName() string
	GetID() int
}

type UserBuilder interface {
	SetName(name string)
	Build() (UserBase, error)
}

func New(userType UserType) (UserBuilder, error) {
	switch userType {
	case Student:
		return nil, nil
	case Teacher:
		return nil, nil
	default:
		return nil, errors.New(`invalid user type`)
	}
}
