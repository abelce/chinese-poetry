package gen_models

import (
	"time"
	structUtils "github.com/abelce/chinese-poetry/at/structUtils"
	"github.com/asaskevich/govalidator"
)

// {{.Title}}
type {{.Name}} struct {
	{{getStructBody .}}
}

func (entity *{{.Name}}) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func New{{.Name}}(id string, data interface{}) (*{{.Name}}, error) {
	entity := new({{.Name}})
	
	structUtils.SetStructValueFromMap(data, entity)
	entity.Id = id

    entity.IsDeleted = false
    entity.CreateTime = time.Now().UnixNano() / int64(time.Millisecond)
	entity.UpdateTime = time.Now().UnixNano() / int64(time.Millisecond)
	//entity.CreatorId = entity.OperatorId
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *{{.Name}}) Delete() {
	entity.IsDeleted = true
	entity.UpdateTime = time.Now().UnixNano() / int64(time.Millisecond)
}

func (entity *{{.Name}}) Update(data interface{}) error {
	structUtils.SetStructValueFromMap(data, entity)
	entity.UpdateTime = time.Now().UnixNano() / int64(time.Millisecond)

	return entity.Valid()
}
