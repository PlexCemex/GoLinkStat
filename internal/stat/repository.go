package stat

import (
	"gorm.io/datatypes"
	"projects/GoLinkStat/pkg/db"
	"time"
)

type StatRepository struct {
	db db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		db: *db,
	}
}

func (repo *StatRepository) AddClick(linkID uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.db.Find(&stat, "link_id = ? and date = ?", linkID, currentDate)
	if stat.ID == 0 {
		repo.db.Create(&Stat{
			LinkID: linkID,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.db.Save(&stat)
	}
}
