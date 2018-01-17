package model

import (
	"db/mgo"
	"template-builder/tbs/x/math"
)

type TableID struct {
	*mgo.UnsafeTable
}

func NewTableID(name string, idPrefix string) *mgo.UnsafeTable {
	var db = GetDB()
	var idMaker = math.RandStringMaker{Prefix: idPrefix, Length: 20}
	return mgo.NewUnsafeTable(db, name, &idMaker)
}
