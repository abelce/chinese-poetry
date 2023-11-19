package command

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

const (
	queryTypePath = ""
)

type GqlCommand struct {
	BasePath string
	Entities []*model.Author
}

func NewGqlCommand(basePath string, entities []*model.Author) GqlCommand {
	return GqlCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t GqlCommand) Execute() {
	//for _, entity := range t.Entities {
	fmt.Println("[generate graphql mutation type-------------------]")
	GenerateGql(t.BasePath, t.Entities)
	//}
}

func (t GqlCommand) Add(cm Command) {}

func GenerateGql(gqlPath string, entities []*model.Author) {

	path, err := filepath.Abs("./assets/template/queryType.tpl")
	if err != nil {
		panic(err)
	}
	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"hasReferField": hasReferField})
	funcMaps = append(funcMaps, template.FuncMap{"getGraphqlType": getGraphqlType})
	funcMaps = append(funcMaps, template.FuncMap{"proccessFieldName": proccessFieldName})
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})
	funcMaps = append(funcMaps, template.FuncMap{"getReferFieldName": getReferFieldName})

	for _, entity := range entities {
		fmt.Println("[generate graphql-------------------]" + entity.Name)
		result := utils.RenderTemplate("queryType.tpl", path, entity, funcMaps)
		utils.MkdirAll(gqlPath + "/" + queryTypePath)
		err := utils.WriteFile(gqlPath+"/"+queryTypePath+"/"+entity.Name+".go", result)
		if err != nil {
			panic(err)
		}
	}
	// rootQueryType
	GenerateRootGql(gqlPath, entities)
	// listRootQueryType
	GenerateListRootGql(gqlPath, entities)
	// allDataQueryType
	GenerateAllDataGql(gqlPath, entities)
}

// entity是否含有refer字段
func hasReferField(fields []model.Field) bool {
	for _, field := range fields {
		if field.BizType == "refer" {
			return true
		}
	}

	return false
}

func getGraphqlType(field model.Field) string {

	if field.Type == "string" {
		return "graphql.String"
	} else if utils.CoerceInt(field.Type) == "int" {
		// return "graphql.Int"
		// int 也是使用Float， 因为int64在graphql中会被判定为超长，被转成null返回给前端了
		return "graphql.Float"
	} else if utils.CoerceFloat(field.Type) == "float" {
		return "graphql.Float"
	} else if field.Type == "bool" {
		return "graphql.Boolean"
	}

	return ""
}

func getUppercaseName(field model.Field) string {
	return utils.ProccessFieldName(field.Name)
}

type RootQuery struct {
	Entities []*model.Author
}

// 生成rootQueryType
func GenerateRootGql(gqlPath string, entities []*model.Author) {

	fmt.Println("[generate graphql-------------------] rootQueryType")

	path, err := filepath.Abs("./assets/template/rootQueryType.tpl")
	if err != nil {
		panic(err)
	}

	rt := RootQuery{
		Entities: entities,
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("rootQueryType.tpl", path, rt, funcMaps)
	utils.MkdirAll(gqlPath + "/" + queryTypePath)
	err = utils.WriteFile(gqlPath+"/"+queryTypePath+"/rootQueryType.go", result)
	if err != nil {
		panic(err)
	}
}

// 生成listRootQueryType
func GenerateListRootGql(gqlPath string, entities []*model.Author) {

	fmt.Println("[generate graphql-------------------] listRootQueryType")

	path, err := filepath.Abs("./assets/template/listRootQueryType.tpl")
	if err != nil {
		panic(err)
	}

	rt := RootQuery{
		Entities: entities,
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("listRootQueryType.tpl", path, rt, funcMaps)
	utils.MkdirAll(gqlPath + "/" + queryTypePath)
	err = utils.WriteFile(gqlPath+"/"+queryTypePath+"/listRootQueryType.go", result)
	if err != nil {
		panic(err)
	}
}

// 生成 allDataQueryType
func GenerateAllDataGql(gqlPath string, entities []*model.Author) {

	fmt.Println("[generate graphql-------------------] allDataQueryType")

	path, err := filepath.Abs("./assets/template/allDataQueryType.tpl")
	if err != nil {
		panic(err)
	}

	rt := RootQuery{
		Entities: entities,
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("allDataQueryType.tpl", path, rt, funcMaps)
	utils.MkdirAll(gqlPath + "/" + queryTypePath)
	err = utils.WriteFile(gqlPath+"/"+queryTypePath+"/allDataQueryType.go", result)
	if err != nil {
		panic(err)
	}
}
