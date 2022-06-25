package repo

type ILogRepo interface {
	Find()
	List()
	Count()
}
