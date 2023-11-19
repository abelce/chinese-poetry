package command

import (
	"fmt"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

type MainCommand struct {
	Entities []*model.Author
	//CommandList []Command
}

func NewMainCommand(entities []*model.Author) MainCommand {
	return MainCommand{
		Entities: entities,
	}
}

func (t MainCommand) Execute() {
	t.ExecuteRecord()
	t.ExecuteEnums()
}
func (t MainCommand) Add(cm Command) {}

// 实体
func (t MainCommand) ExecuteRecord() {
	var entities []*model.Author
	for _, entity := range t.Entities {
		if entity.Type == "record" {
			entities = append(entities, entity)
		}
	}

	constantCommand := NewConstantCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	modelCommand := NewModelCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	gqlCommand := NewGqlCommand(utils.GetRealPath(utils.GqlPath), entities)
	mutationtypeCommand := NewMutationTypeCommand(utils.GetRealPath(utils.MutationTypePath), entities)
	databaseCommand := NewDatabaseCommand(utils.GetRealPath(utils.DatabasePath), entities)
	queryCondition := NewQueryCondition(utils.GetRealPath(utils.QueryCondition), entities)

	var CommandList []Command
	CommandList = append(CommandList, constantCommand)
	CommandList = append(CommandList, modelCommand)
	CommandList = append(CommandList, gqlCommand)
	CommandList = append(CommandList, mutationtypeCommand)
	// databaseCommand 该命令中的建库脚本、docker.sh脚本、数据库的docker配置可以考虑使用子命令来组合（组合模式）, 暂时在一个脚本集中处理
	CommandList = append(CommandList, databaseCommand)
	CommandList = append(CommandList, queryCondition)

	for _, cmd := range CommandList {
		cmd.Execute()
	}
}

// 枚举
func (t MainCommand) ExecuteEnums() {
	var entities []*model.Author
	for _, entity := range t.Entities {
		if entity.Type == "enum" {
			entities = append(entities, entity)
		}
	}
	fmt.Println(len(entities))

	enumCommand := NewEnumCommand(utils.GetRealPath(utils.CodeGenPath), entities)

	var CommandList []Command
	CommandList = append(CommandList, enumCommand)

	for _, cmd := range CommandList {
		cmd.Execute()
	}
}
