package service

import (
	"app-log/app/repo"
	"fmt"
)

type mongoLogSvc struct {
	logRepo repo.ILogRepo
}

func NewMongoSvc(a repo.ILogRepo) repo.ILogRepo {
	return &mongoLogSvc{
		logRepo: a,
	}
}

func (a *mongoLogSvc) List() {
	a.logRepo.List()
	fmt.Println("i am svc List")
}

func (a *mongoLogSvc) Find() {
	a.logRepo.Find()
	fmt.Println("i am svc find")
}
