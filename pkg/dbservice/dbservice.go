package dbservice

import (
	"shortic/pkg/dbservice/dbaccess"
	"shortic/pkg/dbservice/dblogic"
)

type ConsumerService dblogic.IDbLogic

type DatabaseService struct {
	service ConsumerService
}

func QueueServiceFactory() *DatabaseService {
	return &DatabaseService{
		service: dblogic.NewDb(dbaccess.NewUrlCache()),
	}
}

func (q *DatabaseService) Connect() error {
	return q.service.Connect()
}

func (q *DatabaseService) SaveUrl(fullUrl string, shorterUrl string) error {
	return q.service.SaveUrl(fullUrl, shorterUrl)
}

func (q *DatabaseService) CheckUrlCollision(shortUrlKey string) bool {
	return q.service.CheckUrlCollision(shortUrlKey)
}

func (q *DatabaseService) FindShortenUrl(shortUrlKey string) string {
	return q.service.FindShortenUrl(shortUrlKey)
}
