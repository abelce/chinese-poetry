package queryType

import (
	"github.com/graphql-go/graphql"
	atgql "github.com/abelce/chinese-poetry/at/gql"
	"github.com/abelce/chinese-poetry/at/structUtils"
)

func GetListRootQueryType(isCheckless bool, endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			{{range $i, $entity := .Entities}}
			"{{$entity.Name}}": &graphql.Field{
				Type: atgql.GetPageType("{{$entity.Name}}", graphql.NewList(Get{{$entity.Name}}Type(isCheckless, endpoint))),
				Args: atgql.Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					result, err := atgql.ListResolver(p, isCheckless, endpoint, "{{$entity.Name}}")
					if err != nil {
						return nil, err
					}
					if !result.Success {
						return nil, atgql.GqlFieldFaildReturn(result.Success, result.Code, result.ErrMsg)
					}
					return structUtils.StructToMap(result.Data)
				},
			},
			{{end}}
		},
	})
}
