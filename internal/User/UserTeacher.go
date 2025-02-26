package user

type UserTeacher struct {
	name string
	id   int
}

func (u *UserTeacher) GetUserType() UserType {
	return Student
}

func (u *UserTeacher) GetName() string {
	return u.name
}

func (u *UserTeacher) GetID() int {
	return u.id
}

type UserTeacherBuilder struct {
	name      string
	generator *generatorID
}

func (b *UserTeacherBuilder) SetName(name string) UserBuilder {
	b.name = name
	return b
}

func (b *UserTeacherBuilder) Build() (UserBase, error) {
	return &UserTeacher{name: b.name, id: generator.generateID()}, nil
}
