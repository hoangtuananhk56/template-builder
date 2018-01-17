package user

import (
	"fmt"
	// "fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func getUser(where map[string]interface{}) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(where, &u)
}

func GetByID(id string) (*User, error) {
	var u User
	return &u, TableUser.ReadByID(id, &u)
}

func GetByUsername(username string) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(map[string]interface{}{
		"username": username,
		"dtime":    0,
	}, &u)
}

func CheckUsernamePassword(username string, password string) (bool, error) {
	u, err := GetByUsername(username)
	if err = u.ComparePassword(password); err != nil {
		return false, err
	}
	return true, err
}

func GetAll() ([]*User, error) {
	var users = []*User{}
	return users, TableUser.UnsafeReadAll(&users)
}

func Compute(fillter []bson.M) ([]*User, error) {
	var res = []*User{}
	var err = TableUser.C().Pipe(fillter).All(&res)
	return res, err
}

func CountUser(start, end time.Time) int {
	var count, err = TableUser.C().Find(bson.M{"mtime": bson.M{"$gte": start.Unix(), "$lt": end.Unix()}}).Count()
	if err != nil {
		fmt.Println(err)
	}
	return count
}
