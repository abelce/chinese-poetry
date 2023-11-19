{{define "number"}}
    case atjsonapi.FilterEq:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' = $%d ", len(queryVars))
    case atjsonapi.FilterNe:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' {{unescaped "<>"}} $%d ", len(queryVars))
    case atjsonapi.FilterGt:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' > $%d ", len(queryVars))
    case atjsonapi.FilterLt:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' {{unescaped "<"}} $%d ", len(queryVars))
    case atjsonapi.FilterGe:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' >= $%d ", len(queryVars))
    case atjsonapi.FilterLe:
        queryVars = append(queryVars, {{getCondValue .}})
        where += fmt.Sprintf(" AND data->>'{{.Name}}' {{unescaped "<="}} $%d ", len(queryVars))
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
    case atjsonapi.FilterBetween:
        where += fmt.Sprintf(" AND data->>'{{.Name}}' between ")
        // 需要判断between的长度为2， 以及无值的情况, 要使用大于等于/小于等于来构造
        for _, v := range cond.Value.SS {
            queryVars = append(queryVars, v)
            where += fmt.Sprintf("$%d and ", len(queryVars))
        }
        where = strings.TrimRight(where, "and")
{{end}}