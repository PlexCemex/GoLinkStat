package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	Modek  gorm.Model
	LinkId uint           `json:"link_id"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date"`
}
