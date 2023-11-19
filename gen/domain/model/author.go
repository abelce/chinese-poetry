package model

type Author struct {
	//作者ID
	Id string `json:"id"`
	//姓名
	Name string `json:"name"`
	//生平
	Biography string `json:"biography"`
	//朝代
	Dynasty string `json:"dynasty"`
	//标签
	Tags []string `json:"tags"`
	//图片
	Img string `json:"img"`
	//是否删除
	IsDeleted bool `json:"isDeleted"`
	//更新时间
	UpdateTime int64 `json:"updateTime"`
	//创建时间
	CreateTime int64 `json:"createTime"`
	//上一次更新者
	OperatorId string `json:"operatorId"`
	//创建者
	CreatorId string `json:"creatorId"`
	//租户
	TenantId string `json:"tenantId"`
	// <-----------  脚本备用字段----------->
	//
	Desc string `json:"desc"`
}
