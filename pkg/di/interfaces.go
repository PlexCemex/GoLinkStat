package di

import "projects/GoLinkStat/internal/user"

type IStatRepository interface {
	AddClick(linkID uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
}
