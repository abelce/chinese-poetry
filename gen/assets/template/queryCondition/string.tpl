
{{define "string"}}
    case atjsonapi.FilterEq:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' = $%d ", len(queryVars))
    case atjsonapi.FilterNe:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' {{unescaped "<>"}} $%d ", len(queryVars))
    case atjsonapi.FilterLike:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' like $%d ", len(queryVars))
    case atjsonapi.FilterIn:
        where += fmt.Sprintf(" AND data->>'{{.Name}}' in (")
        for _, v := range cond.Value.SS {
            queryVars = append(queryVars, v)
            where += fmt.Sprintf("$%d,", len(queryVars))
        }
        where = strings.TrimRight(where, ",")
        where += ")"
    case atjsonapi.FilterNin:
        where += fmt.Sprintf(" AND data->>'{{.Name}}' not in (")
        for _, v := range cond.Value.SS {
            queryVars = append(queryVars, v)
            where += fmt.Sprintf("$%d,", len(queryVars))
        }
        where = strings.TrimRight(where, ",")
        where += ")"
{{end}}