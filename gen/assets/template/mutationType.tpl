package mutationType

import (
	"encoding/json"

	graphqlutils "github.com/abelce/chinese-poetry/at/graphqlUtils"
	"github.com/abelce/chinese-poetry/at/request"
	"github.com/graphql-go/graphql"
)

{{$entityName := .Name}}
var single{{$entityName}}Type *graphql.Object // 使用单例模式
func Get{{$entityName}}Type(endpoint string) *graphql.Object{
	if single{{$entityName}}Type != nil {
		return single{{$entityName}}Type
	}
	single{{$entityName}}Type = graphql.NewObject(graphql.ObjectConfig{
		Name: "{{$entityName}}",
		Fields: graphql.Fields{
            "create": &graphql.Field{
			Type:        {{$entityName}}Type,
			Description: "Create new {{$entityName}}",
			Args: graphql.FieldConfigArgument{
				"data": &graphql.ArgumentConfig{
					Type:         graphqlutils.Interace,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				data, _ := params.Args["data"];
				req := request.Request{
					Url: endpoint + "/v1/{{lowerCase $entityName}}}s",
					Method: "POST",
					Data: data,
				}
				result, err := req.Do()
				if err != nil {
					return nil, err
				}
				var entity map[string]interface{}
				err = json.Unmarshal(result, &entity)
				if err != nil {
					return nil, nil
				}
				
				return entity, nil	
			},
		},
		},
	})
	return single{{$entityName}}Type
}