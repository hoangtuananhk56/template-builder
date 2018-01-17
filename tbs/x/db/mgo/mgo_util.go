package mgo

import (
	"template-builder/tbs/x/mlog"
)

var mongoDBLog = mlog.NewTagLog("MongoDB")

type M map[string]interface{}
