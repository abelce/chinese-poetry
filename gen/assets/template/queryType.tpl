package queryType

import (
	// gen_md "github.com/abelce/chinese-poetry/common/code-gen/models"

	"github.com/graphql-go/graphql"
	{{if hasReferField .Fields}}
	atgql "github.com/abelce/chinese-poetry/at/gql"
	{{end}}
)

{{$entityName := .Name}}
var single{{$entityName}}Type *graphql.Object // 使用单例模式
var single{{$entityName}}ChecklessType *graphql.Object // 使用单例模式
func Get{{$entityName}}Type(isCheckless bool, endpoint string) *graphql.Object{
	
	if isCheckless {
		if single{{$entityName}}ChecklessType != nil {
			return single{{$entityName}}ChecklessType
		}
	} else if single{{$entityName}}Type != nil {
		return single{{$entityName}}Type
	}
	
	result := graphql.NewObject(graphql.ObjectConfig{
		Name: "{{$entityName}}",
		// 使用closure 进行懒加载，这样在typechecking阶段就不会进行检查，避免queryType循环引用时出现堆栈溢出错误
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
			{{range $i, $field := .Fields}}
				//1: 如果是外键
				{{if and (eq $field.BizType "refer") $field.ReferInfo}}
				    {{if $field.ReferInfo.IsChildren}}
						"{{$field.Name}}": &graphql.Field {
							Type: graphql.NewList(Get{{$entityName}}Type(isCheckless, endpoint)),
							Description: "{{$field.Title}}",
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								result, err := atgql.ChildrenListResolver(p, isCheckless, endpoint, "{{$entityName}}")
								if err != nil {
									return nil, err
								}
								if !result.Success {
									return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
								}
								return result.Data, err
							},
						},
					{{else}}
						// 1.2: 如果是普通外键，则查询外键实体数据
						{{if $field.IsMutil}}
							"{{getReferFieldName $field}}": &graphql.Field {
								Type: graphql.NewList(Get{{$field.ReferInfo.ReferEntityName}}Type(isCheckless, endpoint)),
								Resolve: func(p graphql.ResolveParams) (interface{}, error) {
									if p.Source == nil {
										return nil, nil
									}
									var tmp []interface{}
									if val, exist := p.Source.(map[string]interface{})["{{$field.Name}}"]; exist {
										ids := atgql.InterfaceToStringArray(val)
										if len(ids) > 0 {
											return atgql.MutilResolver(p, endpoint, "{{$field.ReferInfo.ReferEntityName}}", ids)
										}
										return tmp, nil
									}
									return tmp, nil
								},
							},
						{{else}}
							"{{getReferFieldName $field}}": &graphql.Field {
								Type: Get{{$field.ReferInfo.ReferEntityName}}Type(isCheckless, endpoint),
								Resolve: func(p graphql.ResolveParams) (interface{}, error) {
									if p.Source == nil {
										return nil, nil
									}
									if val, exist := p.Source.(map[string]interface{})["{{$field.Name}}"]; exist {
										if val.(string) != "" {
											result, err := atgql.EntityResolver(p, isCheckless, endpoint, "{{$field.ReferInfo.ReferEntityName}}", val.(string))
												if err != nil {
													return nil, err
												}
												if !result.Success {
													return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
												}
												return result.Data, err
										}
									}
									return nil, nil
								},
							},
						{{end}}
						// 1.2.1: 普通外键也添加对应的原始字段名称
						"{{$field.Name}}": &graphql.Field {
							Type: {{getGraphqlType $field}},
							Description: "{{$field.Title}}",
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								if p.Source == nil {
									return nil, nil
								}
								if val, exist := p.Source.(map[string]interface{})["{{$field.Name}}"]; exist {
									return val, nil
								}
								return nil, nil
							},
						},
					{{end}}
				{{else if and (eq $field.BizType "items") $field.ReferInfo}}
					//1.1: 如果是子表，则关联查询子表的列表数据
					"{{$field.Name}}": &graphql.Field {
						Type: graphql.NewList(Get{{$field.ReferInfo.ReferEntityName}}Type(isCheckless, endpoint)),
						Description: "{{$field.Title}}",
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							result, err := atgql.DetailListResolver(p, isCheckless, endpoint, "{{$field.ReferInfo.ReferEntityName}}", "{{$field.AssociateField}}")
							if err != nil {
								return nil, err
							}
							if !result.Success {
								return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
							}
							return result.Data.Data, err
						},
					},
				{{else}}
					// 2: 普通属性
					{{if $field.IsMutil}}
						"{{$field.Name}}": &graphql.Field {
							Type: graphql.NewList({{getGraphqlType $field}}),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								if p.Source == nil {
									return nil, nil
								}
								var tmp []interface{}
								if val, exist := p.Source.(map[string]interface{})["{{$field.Name}}"]; exist {
									return val, nil
								}
								return tmp, nil
							},
						},
					{{else}}
						"{{$field.Name}}": &graphql.Field {
							Type: {{getGraphqlType $field}},
							Description: "{{$field.Title}}",
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								if p.Source == nil {
									return nil, nil
								}
								// 2.1: 如如果是子级的数量
								{{if $field.IsChildrenCount}}
									result, err := atgql.ChildrenListCountResolver(p, endpoint, "{{$entityName}}")
									if err != nil {
										return nil, err
									}
									if !result.Success {
										return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
									}
									return result.Data.Data, err
								{{else}} 
									// 2.2: 默认
									if val, exist := p.Source.(map[string]interface{})["{{$field.Name}}"]; exist {
										return val, nil
									}
									return nil, nil
								{{end}}
							},
						},
					{{end}}
				{{end}}
			{{end}}
		}
		}),
	})

	if isCheckless {
		single{{$entityName}}ChecklessType = result
		return single{{$entityName}}ChecklessType
	}

	single{{$entityName}}Type = result
	return single{{$entityName}}Type
}