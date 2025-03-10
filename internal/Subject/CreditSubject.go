package subject

import (
	"errors"

	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
	components "github.com/mhgffqwoer/EducationalPlatform/internal/Subject/Components"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

type CreditSubject struct {
	author   user.UserBase
	name     string
	labworks []components.Labwork
	lectures []components.LectureMaterials
	points   int
	id       *generatorid.ID
	basic_id *generatorid.ID
}

func (s *CreditSubject) GetType() SubjectType {
	return CreditSubjectType
}

func (s *CreditSubject) GetName() string {
	return s.name
}

func (s *CreditSubject) GetAuthor() user.UserBase {
	return s.author
}

func (s *CreditSubject) GetLabworks() []components.Labwork {
	return s.labworks
}

func (s *CreditSubject) GetLectures() []components.LectureMaterials {
	return s.lectures
}

func (s *CreditSubject) GetPoints() int {
	return s.points
}

func (s *CreditSubject) GetID() *generatorid.ID {
	return s.id
}

func (s *CreditSubject) GetBasicID() *generatorid.ID {
	return s.basic_id
}

func (s *CreditSubject) Clone() Subject {
	labworks := make([]components.Labwork, len(s.labworks))
	for i, labwork := range s.labworks {
		labworks[i] = labwork.Clone()
	}

	lectures := make([]components.LectureMaterials, len(s.lectures))
	for i, lecture := range s.lectures {
		lectures[i] = lecture.Clone()
	}

	builder := &CreditSubjectBuilder{}
	result, _ := builder.
		SetAuthor(s.author).
		SetName(s.name).
		SetLectures(lectures).
		SetLabworks(labworks).
		Build()

	if result, ok := result.(*CreditSubject); ok {
		result.basic_id = s.id
	}
	return result
}

func (s *CreditSubject) ChangeName(newName string, author user.UserBase) error {
	if author == s.author {
		s.name = newName
		return nil
	}
	return errors.New("user is not author")
}

func (s *CreditSubject) ChangeLectures(newLectures []components.LectureMaterials, author user.UserBase) error {
	if author == s.author {
		s.lectures = newLectures
		return nil
	}
	return errors.New("user is not author")
}

type CreditSubjectBuilderAuthor interface {
	SetAuthor(author user.UserBase) CreditSubjectBuilderName
}

type CreditSubjectBuilderName interface {
	SetName(name string) CreditSubjectBuilderLectures
}

type CreditSubjectBuilderLectures interface {
	SetLectures(lectures []components.LectureMaterials) CreditSubjectBuilderLabworks
}

type CreditSubjectBuilderLabworks interface {
	SetLabworks(labworks []components.Labwork) CreditSubjectBuilderFinish
}

type CreditSubjectBuilderFinish interface {
	Build() (Subject, error)
}

type CreditSubjectBuilder struct {
	author   user.UserBase
	name     string
	labworks []components.Labwork
	lectures []components.LectureMaterials
	id       *generatorid.ID
}

func (b *CreditSubjectBuilder) SetAuthor(author user.UserBase) CreditSubjectBuilderName {
	b.author = author
	return b
}

func (b *CreditSubjectBuilder) SetName(name string) CreditSubjectBuilderLectures {
	b.name = name
	return b
}

func (b *CreditSubjectBuilder) SetLectures(lectures []components.LectureMaterials) CreditSubjectBuilderLabworks {
	b.lectures = lectures
	return b
}

func (b *CreditSubjectBuilder) SetLabworks(labworks []components.Labwork) CreditSubjectBuilderFinish {
	b.labworks = labworks
	return b
}

func (b *CreditSubjectBuilder) Build() (Subject, error) {
	allPoints := 0
	for _, labwork := range b.labworks {
		allPoints += labwork.GetPoints()
	}
	if allPoints != 100 {
		return nil, errors.New("all subject points must be equal to 100")
	}
	return &CreditSubject{
		author:   b.author,
		name:     b.name,
		labworks: b.labworks,
		lectures: b.lectures,
		points:   0,
		id:       generator.GenerateID(),
	}, nil
}
