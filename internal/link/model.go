package link

import (
	"math/rand"
	"projects/GoLinkStat/internal/stat"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"containts:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

var lenOfHash = 7

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = randStringRunes(lenOfHash)
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGIJKLMNOPRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(result)
}
