package memory

import "github.com/julianojj/microservices/accounts/src/domain/entity"

type UserRepositoryMemory struct {
	users []*entity.User
}

func NewCreateUserRepository() *UserRepositoryMemory {
	return &UserRepositoryMemory{}
}

func (u *UserRepositoryMemory) Save(user *entity.User) {
	u.users = append(u.users, user)
}

func (u *UserRepositoryMemory) Find(userId string) *entity.User {
	var output *entity.User
	for _, user := range u.users {
		if user.Id == userId {
			output = user
			break
		}
	}
	return output
}

func (u *UserRepositoryMemory) FindByEmail(email string) *entity.User {
	var output *entity.User
	for _, user := range u.users {
		if user.Email == email {
			output = user
			break
		}
	}
	return output
}
