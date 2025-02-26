package user

type UserStudent struct {
	name string
	id   int
}

func (u *UserStudent) GetUserType() UserType {
	return Student
}

func (u *UserStudent) GetName() string {
	return u.name
}

func (u *UserStudent) GetID() int {
	return u.id
}

type UserStudentBuilder struct {
	name      string
	generator *generatorID
}

func (b *UserStudentBuilder) SetName(name string) UserBuilder {
	b.name = name
	return b
}

func (b *UserStudentBuilder) Build() (UserBase, error) {
	return &UserStudent{name: b.name, id: generator.generateID()}, nil
}
