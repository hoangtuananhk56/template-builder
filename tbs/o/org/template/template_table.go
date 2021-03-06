package template

import (
	"db/mgo"
	"template-builder/tbs/o/model"
)

var TableTemplate = model.NewTable("temp", "temp")

func NewTempID() string {
	return TableTemplate.Next()
}

func (b *Template) Create() error {
	return TableTemplate.Create(b)
}

func (v *Template) Update(newValue *Template) error {
	var values = map[string]interface{}{}

	values["mode"] = newValue.Mode
	values["type"] = newValue.Type
	values["data"] = newValue.Data
	values["image"] = newValue.Image
	values["name"] = newValue.Name
	return TableTemplate.UnsafeUpdateByID(v.ID, values)
}

var _ = TableTemplate.EnsureIndex(mgo.Index{
	Key:        []string{"mtime"},
	Background: true,
})
