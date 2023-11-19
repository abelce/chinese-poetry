package main

import "github.com/abelce/chinese-poetry/utils"

// <! ----- 全唐诗  ------ >
// 导入作者

func importTangPoetryAuthor() {

	req := Request{}
	req.SetHeader("tenantId", "96f16846-31f2-489c-9af0-d4ca13e836e4")
	req.SetHeader("operatorId", "96f16846-31f2-489c-9af0-d4ca13e836e4")

	req.Url = "http://127.0.0.1/v1/Author/"
	req.Method = "POST"
	result, err := req.Do()
	if at.Ensure(&err) {
		return nil, err
	}
}

func main() {
	fileNames := utils.ReadJsonFiles(utils.GetRealPath(utils.EntityPath))
	importTangPoetryAuthor()
}
