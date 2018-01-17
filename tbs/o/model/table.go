package model

import (
	"db/mgo"
	"template-builder/tbs/config/cons"
	"template-builder/tbs/x/math"
)

type TableWithBranchCode struct {
	*mgo.Table
}

type TableWithCode struct {
	*mgo.Table
}

type TableWithType struct {
	*mgo.Table
}

func NewTable(name string, idPrefix string) *mgo.Table {
	var db = GetDB()
	var idMaker = math.RandStringMaker{Prefix: idPrefix, Length: 20}
	return mgo.NewTable(db, name, &idMaker)
}

func NewTableWithCode(name string, idPrefix string) *TableWithCode {
	var table = NewTable(name, idPrefix)
	return &TableWithCode{Table: table}
}

func NewTableWithBranchCode(name string, idPrefix string) *TableWithBranchCode {
	var table = NewTable(name, idPrefix)
	return &TableWithBranchCode{Table: table}
}

func NewTableWithType(name string, idPrefix string) *TableWithType {
	var table = NewTable(name, idPrefix)
	return &TableWithType{Table: table}
}

func GetDB() *mgo.Database {
	return mgo.GetDB(cons.ENV_OBJECT_DB)
}
