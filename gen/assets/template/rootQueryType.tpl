package queryType

import (
	atgql "github.com/abelce/chinese-poetry/at/gql"
	"github.com/graphql-go/graphql"
)

func GetRootQueryType(isCheckless bool, endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			{{range $i, $entity := .Entities}}
			"{{$entity.Name}}": &graphql.Field{
				Type: Get{{$entity.Name}}Type(isCheckless, endpoint),
				Args: atgql.Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						result, err := atgql.EntityResolver(p, isCheckless, endpoint, "{{$entity.Name}}", id)
						if err != nil {
							return nil, err
						}
						if !result.Success {
							return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
						}
						return result.Data, err
					}

					return nil, nil
				},
			},
			{{end}}
		},
	})
}
