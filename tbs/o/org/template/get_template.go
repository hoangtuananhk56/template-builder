package template

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)

func GetByID(id string) (*Template, error){
	var temp Template
	return &temp, TableTemplate.ReadByID(id, &temp)
}

func GetAll() ([]*Template, error) {
	var temp = []*Template{}
	return temp, TableTemplate.UnsafeReadAll(&temp)
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