package basebootstrap

import (
	"fmt"

	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
	"golang.org/x/crypto/bcrypt"
)

func (b Bootstrap) createSuperAdmin(storage *basestorage.Storage) error {
	var adminCount int64
	countResult := storage.DB.Model(&userentity.User{}).Where("first_name = ? AND last_name = ?", "Super", "Admin").Count(&adminCount)
	if countResult.Error != nil {
		return fmt.Errorf("unable to check admin count : %v", countResult.Error)
	}

	if adminCount > 0 {
		return nil
	}

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte("nima1381"), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		return fmt.Errorf("password not hashed : %v", hashedPasswordErr)
	}
	admin := userentity.NewUser(
		"Super",
		"Admin",
		"09925087579",
		string(hashedPassword),
	)
	adminResult := storage.DB.Model(&userentity.User{}).Create(admin)
	if adminResult.Error != nil {
		return fmt.Errorf("super admin not created : %v", adminResult.Error)
	}
	superRole := roleentity.NewRole("Super Admin", &admin.ID, roleentity.Permissions{roleentity.SUPER_ADMIN})
	roleResult := storage.DB.Model(&roleentity.Role{}).Create(superRole)
	if roleResult.Error != nil {
		return fmt.Errorf("role not created: %v", roleResult.Error)
	}
	admin.RoleID = &superRole.ID
	assignRoleResult := storage.DB.Save(admin)
	if assignRoleResult.Error != nil {
		return fmt.Errorf("role not assign : %v", assignRoleResult.Error)
	}
	return nil
}
