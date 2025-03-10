package subject

import (
	"errors"

	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
	components "github.com/mhgffqwoer/EducationalPlatform/internal/Subject/Components"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

type ExamSubject struct {
	author   user.UserBase
	name     string
	labworks []components.Labwork
	lectures []components.LectureMaterials
	points   int
	id       *generatorid.ID
	basic_id *generatorid.ID
}

func (s *ExamSubject) GetType() SubjectType {
	return ExamSubjectType
}

func (s *ExamSubject) GetName() string {
	return s.name
}

func (s *ExamSubject) GetAuthor() user.UserBase {
	return s.author
}

func (s *ExamSubject) GetLabworks() []components.Labwork {
	return s.labworks
}

func (s *ExamSubject) GetLectures() []components.LectureMaterials {
	return s.lectures
}

func (s *ExamSubject) GetPoints() int {
	return s.points
}

func (s *ExamSubject) GetID() *generatorid.ID {
	return s.id
}

func (s *ExamSubject) GetBasicID() *generatorid.ID {
	return s.basic_id
}

func (s *ExamSubject) Clone() Subject {
	labworks := make([]components.Labwork, len(s.labworks))
	for i, labwork := range s.labworks {
		labworks[i] = labwork.Clone()
	}

	lectures := make([]components.LectureMaterials, len(s.lectures))
	for i, lecture := range s.lectures {
		lectures[i] = lecture.Clone()
	}

	builder := &ExamSubjectBuilder{}
	result, _ := builder.
		SetExamPoints(s.points).
		SetAuthor(s.author).
		SetName(s.name).
		SetLectures(lectures).
		SetLabworks(labworks).
		Build()

	if result, ok := result.(*ExamSubject); ok {
		result.basic_id = s.id
	}
	return result
}

func (s *ExamSubject) ChangeName(newName string, author user.UserBase) error {
	if author == s.author {
		s.name = newName
		return nil
	}
	return errors.New("user is not author")
}

func (s *ExamSubject) ChangeLectures(newLectures []components.LectureMaterials, author user.UserBase) error {
	if author == s.author {
		s.lectures = newLectures
		return nil
	}
	return errors.New("user is not author")
}

type ExamSubjectBuilderPoints interface {
	SetExamPoints(points int) ExamSubjectBuilderAuthor
}

type ExamSubjectBuilderAuthor interface {
	SetAuthor(author user.UserBase) ExamSubjectBuilderName
}

type ExamSubjectBuilderName interface {
	SetName(name string) ExamSubjectBuilderLectures
}

type ExamSubjectBuilderLectures interface {
	SetLectures(lectures []components.LectureMaterials) ExamSubjectBuilderLabworks
}

type ExamSubjectBuilderLabworks interface {
	SetLabworks(labworks []components.Labwork) ExamSubjectBuilderFinish
}

type ExamSubjectBuilderFinish interface {
	Build() (Subject, error)
}

type ExamSubjectBuilder struct {
	author   user.UserBase
	name     string
	labworks []components.Labwork
	lectures []components.LectureMaterials
	points   int
	id       *generatorid.ID
}

func (b *ExamSubjectBuilder) SetExamPoints(points int) ExamSubjectBuilderAuthor {
	b.points = points
	return b
}

func (b *ExamSubjectBuilder) SetAuthor(author user.UserBase) ExamSubjectBuilderName {
	b.author = author
	return b
}

func (b *ExamSubjectBuilder) SetName(name string) ExamSubjectBuilderLectures {
	b.name = name
	return b
}

func (b *ExamSubjectBuilder) SetLectures(lectures []components.LectureMaterials) ExamSubjectBuilderLabworks {
	b.lectures = lectures
	return b
}

func (b *ExamSubjectBuilder) SetLabworks(labworks []components.Labwork) ExamSubjectBuilderFinish {
	b.labworks = labworks
	return b
}

func (b *ExamSubjectBuilder) Build() (Subject, error) {
	allPoints := 0
	for _, labwork := range b.labworks {
		allPoints += labwork.GetPoints()
	}
	if b.points < 0 {
		return nil, errors.New("too few exam points")
	}
	if b.points+allPoints != 100 {
		return nil, errors.New("all subject points must be equal to 100")

	}
	return &ExamSubject{
		author:   b.author,
		name:     b.name,
		labworks: b.labworks,
		lectures: b.lectures,
		points:   b.points,
		id:       generator.GenerateID(),
	}, nil
}
