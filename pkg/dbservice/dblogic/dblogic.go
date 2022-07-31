package dblogic

import (
  //"context"
)

type IDbLogic interface{
  Connect() error 
  SaveUrl() error // Not implemented
}

type DbBusinnessLogic struct{
  DbAbs IDbLogic
}

func NewDb(dbaccess IDbLogic) *DbBusinnessLogic{
	return &DbBusinnessLogic{
		DbAbs: dbaccess,
	}
}

func (db *DbBusinnessLogic) Connect() error {
	return db.DbAbs.Connect()
}

func (db *DbBusinnessLogic) Enqueue() error {
	db.DbAbs.SaveUrl()
	return nil
}

