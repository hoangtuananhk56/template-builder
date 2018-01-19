package template

import (
	"db/mgo"
	"template-builder/tbs/x/mlog"
)

var objectTempLog = mlog.NewTagLog("obj_temp")

type Template struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string `json:"user_id" bson:"user_id"`
	Mode          string `json:"mode" bson:"mode"`
	Data          string `json:"data" bson:"data"`
	Type          string `json:"type" bson:"type"`
	Image         string `json:"image" bson:"image"`
	Name          string `json:"name" bson:"name"`
}
