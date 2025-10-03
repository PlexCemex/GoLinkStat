package stat

import (
	"fmt"
	"projects/GoLinkStat/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	DB db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		DB: *db,
	}
}

func (repo *StatRepository) AddClick(linkID uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkID, currentDate)
	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkID: linkID,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.DB.Save(&stat)
	}
}

func (repo *StatRepository) GetStat(by string, from, to time.Time) []GetStatResponse {
	var statResponse []GetStatResponse
	var selectQuere string
	switch by {
	case GroupByDay:
		selectQuere = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks) as sum"
	case GroupByMonth:
		selectQuere = "to_char(date, 'YYYY-MM') as period, sum(clicks) as sum"
	}
	result := repo.DB.Table("stats").
		Select(selectQuere).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&statResponse)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return statResponse
}
