package queryType

import (
	"github.com/graphql-go/graphql"
	atgql "github.com/abelce/chinese-poetry/at/gql"
	"github.com/abelce/chinese-poetry/at/structUtils"
)
// 查询所有数据，不会分页
func GetAllDataQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			{{range $i, $entity := .Entities}}
			"{{$entity.Name}}": &graphql.Field{
				Type: graphql.NewList(Get{{$entity.Name}}Type(false, endpoint)),
				Args: atgql.Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					result, err := atgql.AllDataResolver(p, endpoint, "{{$entity.Name}}")
					if err != nil {
						return nil, err
					}
					if !result.Success {
						return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
					}
					return structUtils.StructToMapArray(result.Data.Data)
				},
			},
			{{end}}
		},
	})
}
