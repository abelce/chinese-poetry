// package script

// import (
// 	"bytes"
// 	"fmt"
// 	"html/template"
// 	"path/filepath"
// 	"strings"
// )

// const (
// 	modelsPath = "models"
// )

// type GenModel struct {
// 	BasePath string
// 	Entities []*Entity
// }

// func NewGenModel(basePath string, entities []*Entity) GenModel {
// 	return GenModel{
// 		BasePath: basePath,
// 		Entities: entities,
// 	}
// }

// func (t GenModel) Execute() {
// 	for _, entity := range t.Entities {
// 		fmt.Println("[generate constants-------------------]" + entity.Name)
// 		GenerateModel(t.BasePath, entity)
// 	}
// }

// func (t GenConstant) Add() {}

// func GenerateModel(modelPath string, entity *Entity) {
// 	RemovePath(modelPath)
// 	Mkdir(modelPath)
// 	// 生成fields
// 	generateStruct(modelPath, entity)
// }

// func generateStruct(modelPath string, entity *Entity) {
// 	entityName := entity.Name

// 	var result []string

// 	// 添加crud代码
// 	result = append(result, appendEntityCrudCode(entity))

// 	Mkdir(modelPath + "/" + modelsPath)
// 	err := WriteFile(modelPath+"/"+modelsPath+"/"+entityName+".go", strings.Join(result, "\n"))
// 	if err != nil {
// 		panic(err)
// 	}

// }

// // 获取结构体的内容
// func getStructBody(entity *Entity) template.HTML {
// 	var result []string

// 	for _, field := range entity.Fields {
// 		result = append(result, "  //"+field.Title)
// 		result = append(result, "  "+proccessFieldName(field.Name)+" "+field.Type+" "+getTag(field))
// 	}
// 	// 使用template.HTML就不会被转义了
// 	return template.HTML(strings.Join(result, "\n"))
// }

// // 获取字段的tag信息
// func getTag(field Field) string {
// 	var funcs []func(field Field) string

// 	// 添加tag处理函数
// 	funcs = append(funcs, getJsonTag)
// 	funcs = append(funcs, getValid)

// 	var tagArray []string
// 	for _, funcItem := range funcs {
// 		if funcItem != nil {
// 			result := funcItem(field)
// 			if result != "" {
// 				tagArray = append(tagArray, result)
// 			}
// 		}
// 	}

// 	return "`" + strings.Join(tagArray, " ") + "`"
// }

// func getJsonTag(field Field) string {
// 	return `json:"` + field.Name + `"`
// }

// // 添加tag valid处理器
// func getValid(field Field) string {
// 	if field.Valid == "" {
// 		return ""
// 	}
// 	return "valid:\"" + field.Valid + "\""
// }

// // 处理struct的name首字母大写
// func proccessFieldName(name string) string {
// 	if name == "" {
// 		panic("name is can not nil")
// 	}
// 	capture := string([]byte(name)[:1]) // 这里不考虑中文的空间不止一个字节的问题， name通常都是中文
// 	others := string([]byte(name)[1:])
// 	return strings.ToUpper(capture) + others
// }

// // 添加业务对象crud代码
// func appendEntityCrudCode(entity *Entity) string {
// 	var result []string
// 	// result = append(result, appendValid(entity))

// 	// 通过模版来渲染，字符串不好拼接代码
// 	path, err := filepath.Abs("./script/template/model.tpl")
// 	if err != nil {
// 		panic(err)
// 	}
// 	tplStr := renderTemplate(path, entity)
// 	result = append(result, tplStr)

// 	return strings.Join(result, "\n")
// }

// func renderTemplate(tplPath string, entity *Entity) string {
// 	t := template.New("model.tpl")
// 	t.Funcs(template.FuncMap{"getCreateFuncParams": getCreateFuncParams})
// 	t.Funcs(template.FuncMap{"getCreateFuncBody": getCreateFuncBody})
// 	t.Funcs(template.FuncMap{"getStructBody": getStructBody})
// 	t.Funcs(template.FuncMap{"getUpdateParams": getUpdateParams})
// 	t.Funcs(template.FuncMap{"getUpdateBody": getUpdateBody})

// 	file, err := filepath.Abs(tplPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t, err = t.ParseFiles(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var buf bytes.Buffer
// 	err = t.Execute(&buf, entity)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return buf.String()
// }

// func getCreateFuncParams(entity *Entity) string {
// 	var result []string
// 	var excludeFields []string
// 	excludeFields = []string{"isDeleted", "createdTime", "updatedTime"}
// 	for _, field := range entity.Fields {
// 		if !IsIncludeItem(excludeFields, field.Name) {
// 			result = append(result, "  "+field.Name+" "+field.Type+",")
// 		}
// 	}
// 	return strings.Join(result, "\n")
// }

// func getCreateFuncBody(entity *Entity) string {
// 	var result []string
// 	var excludeFields []string
// 	excludeFields = []string{"isDeleted", "createdTime", "updatedTime"}
// 	for _, field := range entity.Fields {
// 		if !IsIncludeItem(excludeFields, field.Name) {
// 			result = append(result, "  "+proccessFieldName(field.Name)+": "+field.Name+",")
// 		}
// 	}
// 	return strings.Join(result, "\n")
// }

// func getUpdateParams(entity *Entity) string {
// 	var result []string
// 	var excludeFields []string
// 	excludeFields = []string{"id", "isDeleted", "createdTime", "updatedTime", "operatorId"}

// 	for _, field := range entity.Fields {
// 		if !IsIncludeItem(excludeFields, field.Name) {
// 			result = append(result, "  "+field.Name+" "+field.Type+",")
// 		}
// 	}
// 	return strings.Join(result, "\n")
// }

// func getUpdateBody(entity *Entity) string {
// 	var result []string
// 	var excludeFields []string
// 	excludeFields = []string{"id", "isDeleted", "createdTime", "updatedTime", "operatorId"}

// 	for _, field := range entity.Fields {
// 		if !IsIncludeItem(excludeFields, field.Name) {
// 			result = append(result, "  entity."+proccessFieldName(field.Name)+"="+field.Name)
// 		}
// 	}

// 	return strings.Join(result, "\n")
// }
