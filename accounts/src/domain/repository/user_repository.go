package repository

import "github.com/julianojj/microservices/accounts/src/domain/entity"

type UserRepository interface {
	Save(user *entity.User)
	Find(userId string) *entity.User
	FindByEmail(email string) *entity.User
}
