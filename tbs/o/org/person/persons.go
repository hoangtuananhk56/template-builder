package person

import (
	"db/mgo"
	"template-builder/tbs/x/mlog"
)

var objectTempLog = mlog.NewTagLog("obj_per")

type Person struct {
	mgo.BaseModel `bson:",inline"`
	Name          string `json:"name" bson:"name"`
	Address       string `json:"address" bson:"address"`
	Age           int    `json:"age" bson:"age"`
	Sex           bool   `json:"sex" bson:"sex"`
	Experience    string `json:"exp" bson:"exp"`
	Position      string `json:"position" bson:"position"`
	Rating        int    `json:"rating" bson:"rating"`
}
