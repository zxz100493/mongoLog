package controller

import (
	"app-log/app/repo"
	"app-log/app/service"
)

var logService repo.ILogRepo

func init() {
	logService = service.NewMongoSvc{logService}
}
