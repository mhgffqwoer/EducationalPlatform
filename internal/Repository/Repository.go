package repository

import (
	"errors"

	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
)

type Identifiable interface {
	GetID() *generatorid.ID
}

type Repository[T Identifiable] interface {
	Add(element T) error
	GetByID(id int) (T, error)
	Update(element T) error
	Delete(id int) error
}

type RepositoryImpl[T Identifiable] struct {
	elements []T
}

func New[T Identifiable]() *RepositoryImpl[T] {
	return &RepositoryImpl[T]{}
}

func (r *RepositoryImpl[T]) Add(element T) error {
	r.elements = append(r.elements, element)
	return nil
}

func (r *RepositoryImpl[T]) GetByID(id int) (T, error) {
	var zero T
	for _, element := range r.elements {
		if element.GetID().Int() == id {
			return element, nil
		}
	}
	return zero, errors.New("element not found")
}

func (r *RepositoryImpl[T]) Update(element T) error {
	for i, e := range r.elements {
		if e.GetID().Int() == element.GetID().Int() {
			r.elements[i] = element
			return nil
		}
	}
	return errors.New("element not found")
}

func (r *RepositoryImpl[T]) Delete(id int) error {
	for i, element := range r.elements {
		if element.GetID().Int() == id {
			r.elements = append(r.elements[:i], r.elements[i+1:]...)
			return nil
		}
	}
	return errors.New("element not found")
}
