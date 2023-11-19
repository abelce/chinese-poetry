package model

type Poetry struct {
	//ID
	Id string `json:"id"`
	//名称
	Name string `json:"name"`
	//部分
	Section string `json:"section"`
	//标签
	Tags []string `json:"tags"`
	//内容
	Paragraphs []string `json:"paragraphs"`
	//作者
	AuthorId string `json:"authorId"`
	//是否删除
	IsDeleted bool `json:"isDeleted"`
	//更新时间
	UpdateTime int64 `json:"updateTime"`
	//创建时间
	CreateTime int64 `json:"createTime"`
	//更新者
	OperatorId string `json:"operatorId"`
	//创建者
	CreatorId string `json:"creatorId"`
	//租户
	TenantId string `json:"tenantId"`

	// <-----------  脚本备用字段----------->
	Title  string `json:"title"`
	Author string `json:"author"`
}
