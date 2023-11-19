package command

import "github.com/abelce/chinese-poetry/gen/domain/model"

type Command interface {
	Add(Command)
	Execute()
}

type GenBase struct {
	BasePath string
	Entities []*model.Author
}
