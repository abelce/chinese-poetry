package cmds

// func Run() {
// 	// 读取entity中的json文件
// 	fileNames := utils.ReadJsonFiles(utils.GetRealPath(utils.EntityPath))
// 	// 存储所有的entity， 方便后面需要所有的entity一起才能处理的任务使用
// 	var entites []*model.Author

// 	for _, fileName := range fileNames {
// 		entity := utils.ReadOneJsonFile(utils.GetRealPath(utils.EntityPath + "/" + fileName))
// 		entites = append(entites, entity)
// 	}

// 	runMainCommand(entites)
// }

// // 统一执行所有的命令
// func runMainCommand(entities []*model.Author) {
// 	mainCommand := command.NewMainCommand(entities)
// 	mainCommand.Execute()
// }

func Run() {
	ImportPoetryTang()
}
