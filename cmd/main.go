package main

import (
	"fmt"

	repository "github.com/mhgffqwoer/EducationalPlatform/internal/Repository"
	subject "github.com/mhgffqwoer/EducationalPlatform/internal/Subject"
	components "github.com/mhgffqwoer/EducationalPlatform/internal/Subject/Components"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

func main() {
	user1, err := user.New().Student().SetName("John Doe").Build()
	if err != nil {
		panic(err)
	}
	lm := components.NewLectureMaterialsBuilder().
		SetAuthor(user1).
		SetName("Lecture Materials").
		SetDescription("Description").
		SetContent("Content").Build()

	fmt.Printf("Author: %s\n", lm.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", lm.GetName())
	fmt.Printf("Description: %s\n", lm.GetDescription())
	fmt.Printf("Content: %s\n", lm.GetContent())
	fmt.Printf("ID: %v\n", lm.GetID().Int())

	lm2 := lm.Clone()

	fmt.Printf("Author: %s\n", lm2.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", lm2.GetName())
	fmt.Printf("Description: %s\n", lm2.GetDescription())
	fmt.Printf("Content: %s\n", lm2.GetContent())
	fmt.Printf("ID: %v\n", lm2.GetID().Int())
	fmt.Printf("Basic ID: %v\n", lm2.GetBasicID().Int())

	lab := components.NewLabworkBuilder().
		SetAuthor(user1).
		SetName("Labwork").
		SetDescription("Description").
		SetCriterias("Content").
		SetPoints(10).Build()

	fmt.Printf("Author: %s\n", lab.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", lab.GetName())
	fmt.Printf("Description: %s\n", lab.GetDescription())
	fmt.Printf("Content: %s\n", lab.GetCriterias())
	fmt.Printf("Points: %d\n", lab.GetPoints())
	fmt.Printf("ID: %v\n", lab.GetID().Int())

	lab2 := lab.Clone()

	fmt.Printf("Author: %s\n", lab2.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", lab2.GetName())
	fmt.Printf("Description: %s\n", lab2.GetDescription())
	fmt.Printf("Content: %s\n", lab2.GetCriterias())
	fmt.Printf("Points: %d\n", lab2.GetPoints())
	fmt.Printf("ID: %v\n", lab2.GetID().Int())
	fmt.Printf("Basic ID: %v\n", lab2.GetBasicID().Int())

	exam, err := subject.New().
		ExamSubject().
		SetExamPoints(90).
		SetAuthor(user1).
		SetName("Exam").
		SetLectures([]components.LectureMaterials{lm}).
		SetLabworks([]components.Labwork{lab}).
		Build()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Author: %s\n", exam.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", exam.GetName())
	fmt.Printf("Points: %d\n", exam.GetPoints())
	fmt.Printf("ID: %v\n", exam.GetID().Int())

	exam2 := exam.Clone()

	fmt.Printf("Author: %s\n", exam2.GetAuthor().GetName())
	fmt.Printf("Name: %s\n", exam2.GetName())
	fmt.Printf("Points: %d\n", exam2.GetPoints())
	fmt.Printf("ID: %v\n", exam2.GetID().Int())
	fmt.Printf("Basic ID: %v\n", exam2.GetBasicID().Int())

	repo := repository.New[subject.Subject]()

	repo.Add(exam)
	repo.Add(exam2)

	fmt.Println(repo.GetByID(1))
	fmt.Println(repo.GetByID(2))
}
