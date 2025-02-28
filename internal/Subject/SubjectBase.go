package subject

import (
	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
	components "github.com/mhgffqwoer/EducationalPlatform/internal/Subject/Components"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

var generator *generatorid.GeneratorID

func init() {
	generator = generatorid.New()
}

type SubjectType string

const (
	ExamSubjectType   SubjectType = "exam"
	CreditSubjectType SubjectType = "credit"
)

type Subject interface {
	GetType() SubjectType
	GetName() string
	GetAuthor() user.UserBase
	GetLabworks() []components.Labwork
	GetLectures() []components.LectureMaterials
	GetPoints() int
	GetID() *generatorid.ID
	GetBasicID() *generatorid.ID
	Clone() Subject
	ChangeName(newName string, author user.UserBase) error
	ChangeLectures(newLectures []components.LectureMaterials, author user.UserBase) error
}

type SubjectFactory interface {
	ExamSubject() ExamSubjectBuilderPoints
	CreditSubject() CreditSubjectBuilderAuthor
}

type SubjectFactoryImpl struct{}

func (s *SubjectFactoryImpl) ExamSubject() ExamSubjectBuilderPoints {
	return &ExamSubjectBuilder{}
}

func (s *SubjectFactoryImpl) CreditSubject() CreditSubjectBuilderAuthor {
	return &CreditSubjectBuilder{}
}

func New() SubjectFactory {
	return &SubjectFactoryImpl{}
}
