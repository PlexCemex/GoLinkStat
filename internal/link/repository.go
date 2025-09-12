package link

import "projects/GoLinkStat/pkg/db"

type LinkRepository struct {
	DataBase *db.Db
}

func NewLinkRepository(dataBase *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: dataBase,
	}
}

func (repo *LinkRepository) Create(link *Link) {

}
