package components

import (
	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

var generatorLabwork *generatorid.GeneratorID

func init() {
	generatorLabwork = generatorid.New()
}

type Labwork interface {
	GetAuthor() user.UserBase
	GetName() string
	GetDescription() string
	GetCriterias() string
	GetPoints() int
	GetID() *generatorid.ID
	GetBasicID() *generatorid.ID
	Clone() Labwork
}

type LabworkImpl struct {
	author      user.UserBase
	name        string
	description string
	criterias   string
	points      int
	id          *generatorid.ID
	basic_id    *generatorid.ID
}

func (l *LabworkImpl) GetAuthor() user.UserBase {
	return l.author
}

func (l *LabworkImpl) GetName() string {
	return l.name
}

func (l *LabworkImpl) GetDescription() string {
	return l.description
}

func (l *LabworkImpl) GetCriterias() string {
	return l.criterias
}

func (l *LabworkImpl) GetPoints() int {
	return l.points
}

func (l *LabworkImpl) GetID() *generatorid.ID {
	return l.id
}

func (l *LabworkImpl) GetBasicID() *generatorid.ID {
	if l.basic_id == nil {
		return nil
	}
	return l.basic_id
}

func (l *LabworkImpl) Clone() Labwork {
	result := NewLabworkBuilder().
		SetAuthor(l.author).
		SetName(l.name).
		SetDescription(l.description).
		SetCriterias(l.criterias).
		SetPoints(l.points).
		Build()
	if result, ok := result.(*LabworkImpl); ok {
		result.basic_id = l.id
	}
	return result
}

type LabworkBuilderAuthor interface {
	SetAuthor(author user.UserBase) LabworkBuilderName
}

type LabworkBuilderName interface {
	SetName(name string) LabworkBuilderDescription
}

type LabworkBuilderDescription interface {
	SetDescription(description string) LabworkBuilderCriterias
}

type LabworkBuilderCriterias interface {
	SetCriterias(criterias string) LabworkBuilderPoints
}

type LabworkBuilderPoints interface {
	SetPoints(points int) LabworkBuilderFinish
}

type LabworkBuilderFinish interface {
	Build() Labwork
}

type LabworkBuilder struct {
	author      user.UserBase
	name        string
	description string
	criterias   string
	points      int
	id          *generatorid.ID
	basic_id    *generatorid.ID
}

func NewLabworkBuilder() LabworkBuilderAuthor {
	return &LabworkBuilder{}
}

func (b *LabworkBuilder) SetAuthor(author user.UserBase) LabworkBuilderName {
	b.author = author
	return b
}

func (b *LabworkBuilder) SetName(name string) LabworkBuilderDescription {
	b.name = name
	return b
}

func (b *LabworkBuilder) SetDescription(description string) LabworkBuilderCriterias {
	b.description = description
	return b
}

func (b *LabworkBuilder) SetCriterias(criterias string) LabworkBuilderPoints {
	b.criterias = criterias
	return b
}

func (b *LabworkBuilder) SetPoints(points int) LabworkBuilderFinish {
	b.points = points
	return b
}

func (b *LabworkBuilder) Build() Labwork {
	return &LabworkImpl{
		author:      b.author,
		name:        b.name,
		description: b.description,
		criterias:   b.criterias,
		points:      b.points,
		id:          generatorLabwork.GenerateID(),
	}
}
