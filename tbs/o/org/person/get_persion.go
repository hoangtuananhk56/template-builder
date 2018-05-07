package person

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetByID(id string) (*Person, error) {
	var temp Person
	return &temp, TablePerson.ReadByID(id, &temp)
}

func GetAll(where map[string]interface{}) ([]*Person, error) {
	var temps = []*Person{}
	return temps, TablePerson.UnsafeReadMany(where, &temps)
}

func Compute(fillter []bson.M) ([]*Person, error) {
	var res = []*Person{}
	var err = TablePerson.C().Pipe(fillter).All(&res)
	return res, err
}

func CountTemp(start, end time.Time) int {
	var count, err = TablePerson.C().Find(bson.M{"mtime": bson.M{"$gte": start.Unix(), "$lt": end.Unix()}}).Count()
	if err != nil {
		fmt.Println(err)
	}
	return count
}
