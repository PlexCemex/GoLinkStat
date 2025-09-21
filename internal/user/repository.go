package user

import "projects/GoLinkStat/pkg/db"

type UserRepository struct {
	dataBase *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		dataBase: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.dataBase.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := repo.dataBase.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
