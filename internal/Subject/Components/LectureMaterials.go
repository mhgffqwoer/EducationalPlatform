package components

import (
	generatorid "github.com/mhgffqwoer/EducationalPlatform/internal/GeneratorID"
	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

var generatorLectureMaterials *generatorid.GeneratorID

func init() {
	generatorLectureMaterials = generatorid.New()
}

type LectureMaterials interface {
	GetAuthor() user.UserBase
	GetName() string
	GetDescription() string
	GetContent() string
	GetID() *generatorid.ID
	GetBasicID() *generatorid.ID
	Clone() LectureMaterials
}

type LectureMaterialsImpl struct {
	author      user.UserBase
	name        string
	description string
	content     string
	id          *generatorid.ID
	basic_id    *generatorid.ID
}

func (l *LectureMaterialsImpl) GetAuthor() user.UserBase {
	return l.author
}

func (l *LectureMaterialsImpl) GetName() string {
	return l.name
}

func (l *LectureMaterialsImpl) GetDescription() string {
	return l.description
}

func (l *LectureMaterialsImpl) GetContent() string {
	return l.content
}

func (l *LectureMaterialsImpl) GetID() *generatorid.ID {
	return l.id
}

func (l *LectureMaterialsImpl) GetBasicID() *generatorid.ID {
	if l.basic_id == nil {
		return nil
	}
	return l.basic_id
}

func (l *LectureMaterialsImpl) Clone() LectureMaterials {
	result := NewLectureMaterialsBuilder().
		SetAuthor(l.author).
		SetName(l.name).
		SetDescription(l.description).
		SetContent(l.content).
		Build()
	if result, ok := result.(*LectureMaterialsImpl); ok {
		result.basic_id = l.id
	}
	return result
}

type LectureMaterialsBuilderAuthor interface {
	SetAuthor(author user.UserBase) LectureMaterialsBuilderName
}

type LectureMaterialsBuilderName interface {
	SetName(name string) LectureMaterialsBuilderDescription
}

type LectureMaterialsBuilderDescription interface {
	SetDescription(description string) LectureMaterialsBuilderContent
}

type LectureMaterialsBuilderContent interface {
	SetContent(content string) LectureMaterialsBuilderFinish
}

type LectureMaterialsBuilderFinish interface {
	Build() LectureMaterials
}

type LectureMaterialsBuilder struct {
	author      user.UserBase
	name        string
	description string
	content     string
	id          *generatorid.ID
	basic_id    *generatorid.ID
}

func NewLectureMaterialsBuilder() LectureMaterialsBuilderAuthor {
	return &LectureMaterialsBuilder{}
}

func (b *LectureMaterialsBuilder) SetAuthor(author user.UserBase) LectureMaterialsBuilderName {
	b.author = author
	return b
}

func (b *LectureMaterialsBuilder) SetName(name string) LectureMaterialsBuilderDescription {
	b.name = name
	return b
}

func (b *LectureMaterialsBuilder) SetDescription(description string) LectureMaterialsBuilderContent {
	b.description = description
	return b
}

func (b *LectureMaterialsBuilder) SetContent(content string) LectureMaterialsBuilderFinish {
	b.content = content
	return b
}

func (b *LectureMaterialsBuilder) Build() LectureMaterials {
	return &LectureMaterialsImpl{
		author:      b.author,
		name:        b.name,
		description: b.description,
		content:     b.content,
		id:          generatorLectureMaterials.GenerateID(),
	}
}
