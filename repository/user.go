package repository

import (
	uuid "github.com/satori/go.uuid"
	"golang-rest-api-validation-example/handler"
	"golang-rest-api-validation-example/model"
)

type UserRepository struct {
	lastID int
	users  map[string]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*model.User),
	}
}

func (userRepository *UserRepository) GetAllUser() []*model.User {
	users := make([]*model.User, 0)
	for _, user := range userRepository.users {
		users = append(users, user)
	}
	return users
}

func (userRepository *UserRepository) SaveUser(user *model.User) *model.User {
	id := uuid.NewV4().String()
	user.ID = id
	userRepository.users[user.ID] = user
	return user
}

func (userRepository *UserRepository) GetUser(id string) (*model.User, error) {
	existUser, found := userRepository.users[id]
	if !found {
		return nil, handler.ResourceNotFoundException("User", "id", id)
	}
	return existUser, nil
}

func (userRepository *UserRepository) UpdateUser(id string, user *model.User) (*model.User, error) {
	existUser, found := userRepository.users[id]
	if !found {
		return nil, handler.ResourceNotFoundException("User", "id", id)
	}

	existUser.Name = user.Name
	existUser.Email = user.Email

	userRepository.users[id] = existUser

	return existUser, nil
}

func (userRepository *UserRepository) DeleteUser(id string) error {
	_, found := userRepository.users[id]
	if !found {
		return handler.ResourceNotFoundException("User", "id", id)
	}
	delete(userRepository.users, id)
	return nil
}
