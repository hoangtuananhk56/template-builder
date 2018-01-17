package mgo

import (
	"qrcode/pba/x/mlog"
)

var mongoDBLog = mlog.NewTagLog("MongoDB")

type M map[string]interface{}
