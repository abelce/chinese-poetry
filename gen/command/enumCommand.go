package command

import (
	"fmt"
	"strings"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

const (
	enumsPath = "gen_enums"
)

type EnumCommand struct {
	BasePath string
	Entities []*model.Author
}

// 只有type == enum的才执行
func NewEnumCommand(basePath string, entities []*model.Author) EnumCommand {
	return EnumCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t EnumCommand) Execute() {
	for _, entity := range t.Entities {
		fmt.Println("[generate enums-------------------]" + entity.Name)
		GenerateEnum(t.BasePath, entity)
	}
}
func (t EnumCommand) Add(cm Command) {}

func GenerateEnum(codeGenPath string, entity *model.Author) {
	utils.RemovePath(codeGenPath)
	utils.Mkdir(codeGenPath)
	// 生成fields
	generateEnumFields(codeGenPath, entity)
}

func generateEnumFields(codeGenPath string, entity *model.Author) {
	entityName := entity.Name

	var result []string

	result = append(result, "package gen_enums") // package
	result = append(result, "")                  // 空行
	result = append(result, "//"+entity.Title)   // 实体名称
	result = append(result, "const (")           // 使用常量
	result = append(result, "  ENUM_"+entityName+" = \""+entityName+"\"")
	for _, field := range entity.Fields {
		result = append(result, "  //"+field.Title+":"+field.Description)
		result = append(result, "  ENUM_"+entityName+"_"+field.Name+" = \""+field.Value.(string)+"\"")
	}
	result = append(result, ")")

	utils.Mkdir(codeGenPath + "/" + enumsPath)
	err := utils.WriteFile(codeGenPath+"/"+enumsPath+"/"+entityName+".go", strings.Join(result, "\n"))
	if err != nil {
		panic(err)
	}

}
