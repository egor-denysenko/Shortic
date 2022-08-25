package dblogic

import (
// "context"
)

type IDbLogic interface {
	Connect() error
	CheckUrlCollision(shortUrlKey string) bool
	SaveUrl(fullUrl string, shortenUrl string) error
	FindShortenUrl(shortUrlKey string) string
}

type DbBusinnessLogic struct {
	DbAbs IDbLogic
}

func NewDb(dbaccess IDbLogic) *DbBusinnessLogic {
	return &DbBusinnessLogic{
		DbAbs: dbaccess,
	}
}

func (db *DbBusinnessLogic) Connect() error {
	return db.DbAbs.Connect()
}

func (db *DbBusinnessLogic) SaveUrl(fullUrl string, shorterUrl string) error {
	return db.DbAbs.SaveUrl(fullUrl, shorterUrl)
}

func (db *DbBusinnessLogic) CheckUrlCollision(shortUrlKey string) bool {
	return db.DbAbs.CheckUrlCollision(shortUrlKey)
}

func (db *DbBusinnessLogic) FindShortenUrl(shortUrlKey string) string {
	return db.DbAbs.FindShortenUrl(shortUrlKey)
}
