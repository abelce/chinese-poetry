package command

import (
	"fmt"
	"html/template"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

const (
	queryConditionPath = "queryCondition"
)

type QueryCondition struct {
	BasePath string
	Entities []*model.Author
}

func NewQueryCondition(basePath string, entities []*model.Author) QueryCondition {
	return QueryCondition{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t QueryCondition) Execute() {
	utils.RemovePath(t.BasePath)
	utils.Mkdir(t.BasePath)
	for _, entity := range t.Entities {
		fmt.Println("[generate queryCondition-------------------]" + entity.Name)
		GenerateQueryCondition(t.BasePath, entity)
	}
}
func (t QueryCondition) Add(cm Command) {}

// func GenerateQueryCondition(databaseDir string, entity *model.Author) {
// 	GenerateSql(databaseDir, entity)
// }

// 生成productdb.sql
func GenerateQueryCondition(baseDir string, entity *model.Author) {

	fmt.Println("[generate entity cqueryCondition-------------------]" + entity.Name)

	// path, err := filepath.Abs("./assets/template/queryCondition/filter.tpl")
	// if err != nil {
	// 	panic(err)
	// }

	pathList := []string{
		"./assets/template/queryCondition/filter.tpl",
		"./assets/template/queryCondition/operand.tpl",
		"./assets/template/queryCondition/bool.tpl",
		"./assets/template/queryCondition/string.tpl",
		"./assets/template/queryCondition/number.tpl",
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})
	funcMaps = append(funcMaps, template.FuncMap{"unescaped": utils.Unescaped})
	funcMaps = append(funcMaps, template.FuncMap{"getCondValue": getCondValue})
	funcMaps = append(funcMaps, template.FuncMap{"isNumber": utils.IsNumber})

	result := utils.RenderMutilTemplate("filter", pathList, entity, funcMaps)

	entityDatabaseDir := baseDir
	utils.MkdirAll(entityDatabaseDir)
	err := utils.WriteFile(entityDatabaseDir+"/"+utils.LowerCase(entity.Name)+".go", result)
	if err != nil {
		panic(err)
	}
}

func getCondValue(Field *model.Field) string {
	if Field.Type == "bool" {
		return "cond.Value.B"
	}
	if utils.CoerceInt(Field.Type) == "int" || utils.CoerceFloat(Field.Type) == "float" {
		if Field.IsMutil {
			return "cond.Value.NS"
		}
		return "cond.Value.N"
	}
	if Field.Type == "string" {
		if Field.IsMutil {
			return "cond.Value.SS"
		}
		return "cond.Value.S"
	}
	return "cond.Value.S"
}
