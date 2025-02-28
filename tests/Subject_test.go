package tests

import (
	"testing"

	subject "github.com/mhgffqwoer/EducationalPlatform/internal/Subject"
	components "github.com/mhgffqwoer/EducationalPlatform/internal/Subject/Components"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

func TestExamSubject(t *testing.T) {
	user, _ := user.New().Student().SetName("John Doe").Build()
	tests := []struct {
		name      string
		labworks  []*components.Labwork
		points    int
		wantID    int
		wantError bool
	}{
		{
			name: "Create exam subject, with one labwork and all points equal to 100",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(10).
					Build(),
			},
			points:    90,
			wantID:    1,
			wantError: false,
		},
		{
			name: "Create exam subject, with one labwork and all points not equal to 100",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(10).
					Build(),
			},
			points:    100,
			wantID:    -1,
			wantError: true,
		},
		{
			name:      "Create exam subject, without labwork and all points equal to 100",
			labworks:  []*components.Labwork{},
			points:    100,
			wantID:    2,
			wantError: false,
		},
		{
			name:      "Create exam subject, without labwork and all points not equal to 100",
			labworks:  []*components.Labwork{},
			points:    90,
			wantID:    -1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, err := subject.New().
				ExamSubject().
				SetExamPoints(tt.points).
				SetAuthor(user).
				SetName("ExamSubject").
				SetLectures([]*components.LectureMaterials{}).
				SetLabworks(tt.labworks).
				Build()
			if tt.wantError {
				if err != nil {
					return
				}
				t.Errorf("Build(), want error")
			}
			if subject.GetID().Int() != tt.wantID {
				t.Errorf("Subject ID = %v, want %v", subject.GetID().Int(), tt.wantID)
			}
		})
	}
}

func TestCreditSubject(t *testing.T) {
	user, _ := user.New().Student().SetName("John Doe").Build()
	tests := []struct {
		name      string
		labworks  []*components.Labwork
		wantID    int
		wantError bool
	}{
		{
			name: "Create credit subject, with one labwork and all points equal to 100",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(100).
					Build(),
			},
			wantID:    3,
			wantError: false,
		},
		{
			name: "Create credit subject, with one labwork and all points not equal to 100",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(10).
					Build(),
			},
			wantID:    -1,
			wantError: true,
		},
		{
			name:      "Create credit subject, without labworks",
			labworks:  []*components.Labwork{},
			wantID:    -1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, err := subject.New().
				CreditSubject().
				SetAuthor(user).
				SetName("ExamSubject").
				SetLectures([]*components.LectureMaterials{}).
				SetLabworks(tt.labworks).
				Build()
			if tt.wantError {
				if err != nil {
					return
				}
				t.Errorf("Build(), want error")
			}
			if subject.GetID().Int() != tt.wantID {
				t.Errorf("Subject ID = %v, want %v", subject.GetID().Int(), tt.wantID)
			}
		})
	}
}

func TestCopyExamSubject(t *testing.T) {
	user, _ := user.New().Student().SetName("John Doe").Build()
	tests := []struct {
		name      string
		labworks  []*components.Labwork
		points    int
		wantError bool
	}{
		{
			name: "Copy exam subject",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(10).
					Build(),
			},
			points:    90,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, err := subject.New().
				ExamSubject().
				SetExamPoints(tt.points).
				SetAuthor(user).
				SetName("ExamSubject").
				SetLectures([]*components.LectureMaterials{}).
				SetLabworks(tt.labworks).
				Build()
			if err != nil {
				t.Errorf("Don't call Build()")
			}
			subject2 := subject.Clone()
			if subject.GetID() != subject2.GetBasicID() {
				if !tt.wantError {
					t.Errorf("Subject ID = %v, want %v", subject.GetID().Int(), subject2.GetBasicID().Int())
				}
			}
		})
	}
}

func TestCopyCreditSubject(t *testing.T) {
	user, _ := user.New().Student().SetName("John Doe").Build()
	tests := []struct {
		name      string
		labworks  []*components.Labwork
		wantError bool
	}{
		{
			name: "Copy credit subject",
			labworks: []*components.Labwork{
				components.NewLabworkBuilder().
					SetAuthor(user).
					SetName("Lab 1").
					SetDescription("Description").
					SetCriterias("Criterias").
					SetPoints(100).
					Build(),
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, err := subject.New().
				CreditSubject().
				SetAuthor(user).
				SetName("ExamSubject").
				SetLectures([]*components.LectureMaterials{}).
				SetLabworks(tt.labworks).
				Build()
			if err != nil {
				t.Errorf("Don't call Build()")
			}
			subject2 := subject.Clone()
			if subject.GetID() != subject2.GetBasicID() {
				if !tt.wantError {
					t.Errorf("Subject ID = %v, want %v", subject.GetID().Int(), subject2.GetBasicID().Int())
				}
			}
		})
	}
}
