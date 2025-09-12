package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url: url,
		Hash: "",
	}
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGIJKLMNOPRSTUVWXYZ")
func RandStringRunes (n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(result)
}