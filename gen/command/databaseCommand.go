package command

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
)

const (
	sqlPath = "interface-fields"
)

type DatabaseCommand struct {
	BasePath string
	Entities []*model.Author
}

func NewDatabaseCommand(basePath string, entities []*model.Author) DatabaseCommand {
	return DatabaseCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t DatabaseCommand) Execute() {
	utils.RemovePath(t.BasePath)
	utils.Mkdir(t.BasePath)
	for _, entity := range t.Entities {
		fmt.Println("[generate database-------------------]" + entity.Name)
		GenerateDatabase(t.BasePath, entity)
		GenerateDockerfile(t.BasePath, entity)
		GenerateDockerSH(t.BasePath, entity)
	}
}
func (t DatabaseCommand) Add(cm Command) {}

func GenerateDatabase(databaseDir string, entity *model.Author) {
	GenerateSql(databaseDir, entity)
}

// 生成productdb.sql
func GenerateSql(databaseDir string, entity *model.Author) {

	fmt.Println("[generate entity create table sql-------------------]" + entity.Name)

	path, err := filepath.Abs("./assets/template/db/sql.tpl")
	if err != nil {
		panic(err)
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("sql.tpl", path, entity, funcMaps)

	entityDatabaseDir := databaseDir + "/" + entity.Name
	utils.MkdirAll(entityDatabaseDir)
	err = utils.WriteFile(entityDatabaseDir+"/"+utils.LowerCase(entity.Name)+"db.sql", result)
	if err != nil {
		panic(err)
	}
}

// 生成Dockerfile
func GenerateDockerfile(databaseDir string, entity *model.Author) {

	fmt.Println("[generate entity Dockerfile-------------------]" + entity.Name)

	path, err := filepath.Abs("./assets/template/db/Dockerfile.tpl")
	if err != nil {
		panic(err)
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("Dockerfile.tpl", path, entity, funcMaps)

	entityDatabaseDir := databaseDir + "/" + entity.Name
	utils.MkdirAll(entityDatabaseDir)
	err = utils.WriteFile(entityDatabaseDir+"/"+"Dockerfile", result)
	if err != nil {
		panic(err)
	}
}

// 生成docker.sh
func GenerateDockerSH(databaseDir string, entity *model.Author) {

	fmt.Println("[generate entity Dockerfile-------------------]" + entity.Name)

	path, err := filepath.Abs("./assets/template/db/dockerSH.tpl")
	if err != nil {
		panic(err)
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("dockerSH.tpl", path, entity, funcMaps)

	entityDatabaseDir := databaseDir + "/" + entity.Name
	utils.MkdirAll(entityDatabaseDir)
	err = utils.WriteFile(entityDatabaseDir+"/"+"docker.sh", result)
	if err != nil {
		panic(err)
	}
}
