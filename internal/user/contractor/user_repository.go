package usercontractor

import userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"

type UserRepository interface {
	CreateUser(user *userentity.User) error
	EditUser(user *userentity.User) error
	FindUserById(id uint) (*userentity.User, error)
	FindUserByPhone(phone string) (*userentity.User, error)
	GetUsers(page, limit int) ([]userentity.User, error)
	GetUsersCount() (int, error)
}
