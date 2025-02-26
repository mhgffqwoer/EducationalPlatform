package user

type UserStudent struct {
	name string
	id   int // TODO id should be generated
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
	name string
}

func (b *UserStudentBuilder) SetName(name string) {
	b.name = name
}

func (b *UserStudentBuilder) Build() (UserBase, error) {
	return &UserStudent{name: b.name, id: 0}, nil // TODO id should be generated
}
