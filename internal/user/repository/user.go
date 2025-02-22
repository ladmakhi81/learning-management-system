package userrepository

import (
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	storage *basestorage.Storage
}

func NewUserRepositoryImpl(
	storage *basestorage.Storage,
) UserRepositoryImpl {
	return UserRepositoryImpl{
		storage: storage,
	}
}

func (r UserRepositoryImpl) FindUserById(id uint) (*userentity.User, error) {
	user := new(userentity.User)
	result := r.storage.DB.Model(&userentity.User{}).First(user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (r UserRepositoryImpl) FindUserByPhone(phone string) (*userentity.User, error) {
	user := new(userentity.User)
	result := r.storage.DB.Model(&userentity.User{}).Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (r UserRepositoryImpl) CreateUser(user *userentity.User) error {
	result := r.storage.DB.Model(&userentity.User{}).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserRepositoryImpl) EditUser(user *userentity.User) error {
	result := r.storage.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserRepositoryImpl) GetUsers(page, limit int) ([]userentity.User, error) {
	users := make([]userentity.User, 0)
	result := r.storage.DB.Model(&userentity.User{}).Order("created_at desc").Limit(limit).Offset(page * limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r UserRepositoryImpl) GetUsersCount() (int, error) {
	var count int64
	result := r.storage.DB.Model(&userentity.User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}
