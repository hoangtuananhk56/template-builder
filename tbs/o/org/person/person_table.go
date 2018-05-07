package person

import (
	"db/mgo"
	"template-builder/tbs/o/model"
)

var TablePerson = model.NewTable("per", "per")

func NewTempID() string {
	return TablePerson.Next()
}

func (b *Person) Create() error {
	return TablePerson.Create(b)
}

func (v *Person) Update(newValue *Person) error {
	var values = map[string]interface{}{}

	return TablePerson.UnsafeUpdateByID(v.ID, values)
}

var _ = TablePerson.EnsureIndex(mgo.Index{
	Key:        []string{"mtime"},
	Background: true,
})
