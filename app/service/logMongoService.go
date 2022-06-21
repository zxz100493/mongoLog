package service

import (
	"app-log/app/repo"
	"fmt"
)

type mongoLogSvc struct {
	logRepo repo.ILogRepo
}

func NewMongoSvc(a repo.ILogRepo) *mongoLogSvc {
	return &mongoLogSvc{
		logRepo: a,
	}
}

func (a *mongoLogSvc) List() {
	a.Find()
	fmt.Println("i am svc List")
}

func (a *mongoLogSvc) Find() {
	fmt.Println("i am svc find")
}
