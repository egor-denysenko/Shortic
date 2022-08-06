package dbservice

import (
	"shortic/pkg/dbservice/dbaccess"
	"shortic/pkg/dbservice/dblogic"
)

type ConsumerService dblogic.IDbLogic

type QueueService struct {
	service *dblogic.DbBusinnessLogic
}

func QueueServiceFactory() *QueueService {
	return &QueueService{
		service: dblogic.NewDb(dbaccess.NewUrlCache()),
	}
}

func (q *QueueService) Connect() error {
	return q.service.Connect()
}

func (q *QueueService) SaveUrl(fullUrl string,shorterUrl string) error {
	return q.service.SaveUrl(fullUrl,shorterUrl)
}
