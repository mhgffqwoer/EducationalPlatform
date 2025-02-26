package user

type UserTeacher struct {
	name string
	id   int // TODO id should be generated
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
	name string
}

func (b *UserTeacherBuilder) SetName(name string) {
	b.name = name
}

func (b *UserTeacherBuilder) Build() (UserBase, error) {
	return &UserTeacher{name: b.name, id: 0}, nil // TODO id should be generated
}
