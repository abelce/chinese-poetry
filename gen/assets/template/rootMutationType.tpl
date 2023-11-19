package queryType

import (
	gen_md "github.com/abelce/chinese-poetry/common/code-gen/models"
	"github.com/abelce/chinese-poetry/at/request"

	"github.com/graphql-go/graphql"
	"encoding/json"
)

func GetRootMutationType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			{{range $i, $entity := .Entities}}
			"{{$entity.Name}}": &graphql.Field{
				Type: Get{{$entity.Name}}Type(endpoint),
                Description: "Create new {{$entity.Name}}",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
                    "data": &graphql.ArgumentConfig{
						Type:         graphql.NewNonNull(graphql.Float),
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/{{$entity.Name}}/" + id,
							Method: "POST",
                            Data: 
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			{{end}}
		},
	})
}
