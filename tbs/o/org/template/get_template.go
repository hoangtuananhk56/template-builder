package template

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetByID(id string) (*Template, error) {
	var temp Template
	return &temp, TableTemplate.ReadByID(id, &temp)
}

func GetAll(where map[string]interface{}) ([]*Template, error) {
	var temps = []*Template{}
	return temps, TableTemplate.UnsafeReadMany(where, &temps)
}

func Compute(fillter []bson.M) ([]*Template, error) {
	var res = []*Template{}
	var err = TableTemplate.C().Pipe(fillter).All(&res)
	return res, err
}

func CountTemp(start, end time.Time) int {
	var count, err = TableTemplate.C().Find(bson.M{"mtime": bson.M{"$gte": start.Unix(), "$lt": end.Unix()}}).Count()
	if err != nil {
		fmt.Println(err)
	}
	return count
}
