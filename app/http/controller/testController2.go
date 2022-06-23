package controller

import (
	"app-log/app/repo"
)

var logService repo.ILogRepo

// func init() {
// 	logService = service.NewMongoSvc{logService}
// }

// func newService()  {

// }

func GetData() {
	logService.Find()
}
