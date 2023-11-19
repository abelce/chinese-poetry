package cmds

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/abelce/chinese-poetry/gen/assets/utils"
	"github.com/abelce/chinese-poetry/gen/domain/model"
	"github.com/abelce/chinese-poetry/gen/request"
)

var dirPath = "../全唐诗"

var tenantId = "63f43ea3-8031-4a5a-83e3-8d2d8b4ebc37"
var operatorId = "96f16846-31f2-489c-9af0-d4ca13e836e4"

func importTangPoetryAuthorRequest(author model.Author) error {

	req := request.Request{}
	req.SetHeader("tenantId", tenantId)
	req.SetHeader("operatorId", operatorId)

	req.Url = "http://127.0.0.1:3078/v1/Author"
	req.Method = "POST"

	req.Payload = author

	_, err := req.Do()

	if err != nil {
		return err
	}

	return nil
}

func getAuthorByName(name string, dynasty string) *model.Author {
	type ListResponseStruct struct {
		Total       int             `json:"total"`
		Data        []*model.Author `json:"data"`
		PageSize    int             `json:"pageSize"`
		PageNo      int             `json:"pageNo"`
		HasNextPage bool            `json:"hasNextPage"`
	}

	type NormalListResponseStruct struct {
		Data    ListResponseStruct `json:"data"`
		ErrMsg  string             `json:"errMsg"`
		Code    int64              `json:"code"`
		Success bool               `json:"success"`
	}

	req := request.Request{}
	req.SetHeader("tenantId", tenantId)
	req.SetHeader("operatorId", operatorId)

	req.Url = "http://127.0.0.1:3078/v1/Author"
	req.Method = "GET"

	params := url.Values{}
	params.Set("filter", "name ValueType.eq '"+name+"' and dynasty ValueType.eq '"+dynasty+"'")

	req.Params = params

	response, err := req.Do()

	if err != nil {
		panic(err)
	}

	result := NormalListResponseStruct{}

	err = json.Unmarshal(response, &result)
	if err != nil {
		panic(err)
	}

	if result.Data.Total == 0 {
		log.Println(name, dynasty)
		panic("没有查找到作者")
	}

	return result.Data.Data[0]
}

func importTangPoetryRequest(author model.Poetry) error {

	req := request.Request{}
	req.SetHeader("tenantId", tenantId)
	req.SetHeader("operatorId", operatorId)

	req.Url = "http://127.0.0.1:3077/v1/Poetry"
	req.Method = "POST"

	req.Payload = author

	_, err := req.Do()

	if err != nil {
		return err
	}

	return nil
}

func importAuthor(fileName string) {
	result := utils.ReadOneJsonFile(utils.GetRealPath(dirPath + "/" + fileName))
	type A = []model.Author
	list := A{}

	err := json.Unmarshal(result, &list)
	if err != nil {
		panic(err)
	}

	for _, item := range list {
		author := model.Author(item)
		log.Println(author.Name)
		author.Biography = author.Desc
		if strings.HasPrefix(fileName, "authors.song") {
			author.Dynasty = "宋"
		} else if strings.HasPrefix(fileName, "authors.tang") {
			author.Dynasty = "唐"
		}
		err := importTangPoetryAuthorRequest(author)
		if err != nil {
			panic(err)
		}
	}
}

func importPoetry(fileName string) {
	// 只执行poet.tang.xxx.json，唐诗三百首.json"
	if !strings.HasPrefix(fileName, "poet") || fileName == "唐诗三百首.json" {
		return
	}
	result := utils.ReadOneJsonFile(utils.GetRealPath(dirPath + "/" + fileName))

	type A = []model.Poetry
	list := A{}

	err := json.Unmarshal(result, &list)
	if err != nil {
		panic(err)
	}

	for _, item := range list {
		poet := model.Poetry(item)
		poet.Name = poet.Title
		log.Println(poet.Name)

		var dynasty string
		if strings.HasPrefix(fileName, "poet.song") {
			dynasty = "宋"
		} else if strings.HasPrefix(fileName, "poet.tang") {
			dynasty = "唐"
		}
		author := getAuthorByName(poet.Author, dynasty)
		poet.AuthorId = author.Id
		fmt.Println("作者名称:", poet.AuthorId)

		err := importTangPoetryRequest(poet)
		if err != nil {
			panic(err)
		}
	}

}

func ImportPoetryTang() {
	// 读取entity中的json文

	fileNames := utils.ReadJsonFiles(utils.GetRealPath(dirPath))
	// 存储所有的entity， 方便后面需要所有的entity一起才能处理的任务使用
	// var entites [][]model.Author

	for _, fileName := range fileNames {
		if strings.HasPrefix(fileName, "authors.") {
			importAuthor(fileName)
		} else if strings.HasPrefix(fileName, "poet.") {
			importPoetry(fileName)
		}

	}
}
