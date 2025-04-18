# Educational Platform

## Описание проекта

Этот проект представляет собой реализацию системы для формирования и редактирования образовательных программ в абстрактном университете. Основная цель проекта — отработать применение порождающих паттернов проектирования, а также принципов ООП, SOLID и GRASP.

Проект написан на языке Go и включает в себя следующие сущности:

- **Пользователи**: Имеют идентификатор и имя, являются авторами создаваемых сущностей.
- **Лабораторные работы**: Содержат идентификатор, наименование, описание, критерии оценивания и количество баллов.
- **Лекционные материалы**: Содержат идентификатор, наименование, краткое описание и контент.
- **Предметы**: Включают идентификатор, название, список лабораторных работ, список лекционных материалов и зачётный формат (экзамен или зачёт).
- **Репозитории**: Отвечают за хранение и поиск сущностей.

## Функциональные требования

1. **Пользователь**:

   - Имеет идентификатор и имя.
   - Привязан к создаваемым сущностям в качестве автора.

2. **Лабораторная работа**:

   - Имеет идентификатор, наименование, описание, критерии оценивания и количество баллов.
   - Может быть создана на основе существующей лабораторной работы.
   - Может быть изменена только автором.

3. **Лекционные материалы**:

   - Имеют идентификатор, наименование, краткое описание и контент.
   - Могут быть созданы на основе существующих материалов.
   - Могут быть изменены только автором.

4. **Предмет**:

   - Имеет идентификатор, название, список лабораторных работ и список лекционных материалов.
   - Зачётный формат: экзамен или зачёт.
   - Суммарное количество баллов должно равняться 100.
   - Может быть создан на основе существующего предмета.
   - Может быть изменён только автором.

## Технологии

- **Язык программирования**: Go
- **Тестирование**: Встроенный пакет `testing`
- **Система контроля версий**: Git (используется GitFlow для управления ветками).
- **Сборка**: `makefile`

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/mhgffqwoer/EducationalPlatform
   ```
2. Пример работы библиотеки (см. `./cmd/main.go`):
   ```bash
   make run
   ```

## Тестирование

Для запуска тестов выполните команду:

```bash
make test
```
