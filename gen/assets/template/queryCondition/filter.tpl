{{define "filter"}}
package queryCondition

import (
    "fmt"
    "log"
	"strings"

	"github.com/abelce/chinese-poetry/at"
	atjsonapi "github.com/abelce/chinese-poetry/at/jsonapi"
)

{{$entityName := .Name}}

func Get{{$entityName}}BaseQuery(filter, sort string, limit, offset uint64) (string, []interface{}, error) {
	queryVars := []interface{}{}

	queryVars = append(queryVars, false)
	where := fmt.Sprintf(" data->>'isDeleted' = $%d ", len(queryVars))
	if filter != "" {
		f, err := atjsonapi.NewFilter(filter)
		if at.Ensure(&err) {
			return "", nil, err
		}
		for _, cond := range f.Conditions {
			switch cond.Name {
            {{range $i, $field := .Fields}}
                case "{{$field.Name}}":
                    {{template "operand" $field}}
			{{end}}
            }
		}
	}

	return where, queryVars, nil
}

func Get{{$entityName}}Query(tableName,filter, sort string, limit, offset uint64) (string, []interface{}, error) {

    where, queryVars, err := Get{{$entityName}}BaseQuery(filter, sort, limit, offset )
    if at.Ensure(&err) {
        return "", nil, err
    }

	query := fmt.Sprintf("SELECT data, count(*) OVER() as total from %s WHERE %s", tableName, where)

    query += get{{$entityName}}OrderBy(sort)
	
	query += get{{$entityName}}Page(offset, limit)
	log.Println("{{$entityName}} query:", query, queryVars)
	return query, queryVars, nil
}

func get{{$entityName}}OrderBy(sort string) string {
    var orderBys []string
	sorts := strings.Split(sort, ",")
	for _, s := range sorts {
		switch s {
        {{range $i, $field := .Fields}}
		case "{{$field.Name}}":
			orderBys = append(orderBys, fmt.Sprintf("data->>'{{$field.Name}}' ASC"))
		case "-{{$field.Name}}":
			orderBys = append(orderBys, fmt.Sprintf("data->>'{{$field.Name}}' DESC"))
        {{end}}
		}
	}
    if len(orderBys) > 0 {
		return " ORDER BY " + strings.Join(orderBys, ",")
	}
    return ""
}

func get{{$entityName}}Page(offset, limit uint64) string {
	if limit > 0 {
		return fmt.Sprintf(" OFFSET %d LIMIT %d", offset, limit)
	}

	if offset > 0 {
		return fmt.Sprintf(" LIMIT %d", offset)
	}

	return ""
}
{{end}}