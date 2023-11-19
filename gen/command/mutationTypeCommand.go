package command

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

const (
	mutationTypePath = ""
)

type MutationTypeCommand struct {
	BasePath string
	Entities []*model.Author
}

func NewMutationTypeCommand(basePath string, entities []*model.Author) MutationTypeCommand {
	return MutationTypeCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t MutationTypeCommand) Execute() {
	//for _, entity := range t.Entities {
	fmt.Println("[generate graphql mutation type-------------------]")
	GenerateMutationType(t.BasePath, t.Entities)
	//}
}

func (t MutationTypeCommand) Add(cm Command) {}

func GenerateMutationType(gqlPath string, entities []*model.Author) {

	path, err := filepath.Abs("./assets/template/mutationType.tpl")
	if err != nil {
		panic(err)
	}
	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"getGraphqlType": getGraphqlType})
	funcMaps = append(funcMaps, template.FuncMap{"proccessFieldName": proccessFieldName})
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	for _, entity := range entities {
		fmt.Println("[generate graphql mutationType-------------------]" + entity.Name)
		result := utils.RenderTemplate("mutationType.tpl", path, entity, funcMaps)
		utils.MkdirAll(gqlPath + "/" + mutationTypePath)
		err := utils.WriteFile(gqlPath+"/"+mutationTypePath+"/"+entity.Name+".go", result)
		if err != nil {
			panic(err)
		}
	}
	// rootQueryType
	// GenerateRootGql(gqlPath, entities)
	// listRootQueryType
	// GenerateListRootGql(gqlPath, entities)
}

// func getUppercaseName(field model.Field) string {
// 	return utils.ProccessFieldName(field.Name)
// }

// type RootQuery struct {
// 	Entities []*model.Author
// }

// // 生成rootQueryType
// func GenerateRootGql(gqlPath string, entities []*model.Author) {

// 	fmt.Println("[generate graphql-------------------] rootQueryType")

// 	path, err := filepath.Abs("./assets/template/rootQueryType.tpl")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rt := RootQuery{
// 		Entities: entities,
// 	}

// 	var funcMaps []template.FuncMap
// 	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

// 	result := utils.RenderTemplate("rootQueryType.tpl", path, rt, funcMaps)
// 	utils.MkdirAll(gqlPath + "/" + mutationTypePath)
// 	err = utils.WriteFile(gqlPath+"/"+mutationTypePath+"/rootQueryType.go", result)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // 生成listRootQueryType
// func GenerateListRootGql(gqlPath string, entities []*model.Author) {

// 	fmt.Println("[generate graphql-------------------] listRootQueryType")

// 	path, err := filepath.Abs("./assets/template/listRootQueryType.tpl")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rt := RootQuery{
// 		Entities: entities,
// 	}

// 	var funcMaps []template.FuncMap
// 	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

// 	result := utils.RenderTemplate("listRootQueryType.tpl", path, rt, funcMaps)
// 	utils.MkdirAll(gqlPath + "/" + mutationTypePath)
// 	err = utils.WriteFile(gqlPath+"/"+mutationTypePath+"/listRootQueryType.go", result)
// 	if err != nil {
// 		panic(err)
// 	}
// }
