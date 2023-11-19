package model

type ReferInfo struct {
	ReferEntityName string `json:"referEntityName"`
	AliasName       string `json:"aliasName"`
	IsChildren      bool   `json:"isChildren"`
}

type DetailInfo struct {
	DetailEntityName string `json:"detailEntityName"`
}

// type SourceInfo struct {
// 	SourceEntityName string `json:"sourceEntityName"`
// }

type Field struct {
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	DetailInfo *DetailInfo `json:"detailInfo"`
	ReferInfo  *ReferInfo  `json:"referInfo"`
	// SourceInfo      *SourceInfo `json:"sourceInfo"`
	Type            string      `json:"type"`
	Valid           string      `json:"valid"`
	Value           interface{} `json:"value"`
	IsRequired      bool        `josn:"isRequired"`
	BizType         string      `json:"bizType"`
	IsChildrenCount bool        `json:"isChildrenCount"`
	Description     string      `json:"description"`
	IsMutil         bool        `json:"isMutil"`
	AssociateField  string      `json:"associateField"`
}

type Entity struct {
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Fields      []Field `json:"fields"`
	Type        string  `json:"type"`
	Description string  `json:"description"` // 描述文字
}
