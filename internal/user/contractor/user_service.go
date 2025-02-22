package usercontractor

import (
	"mime/multipart"

	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
)

type UserService interface {
	VerifyTeacherByAdmin(adminId uint, dto userrequestdto.VerifyTeacherByAdminReqDTO) error
	UpdateBaseInformation(userId uint, dto userrequestdto.UpdateBaseInformationReqBody) error
	CreateUser(dto userrequestdto.CreateUserReqDTO) (*userentity.User, error)
	ChangePassword(executorId uint, dto userrequestdto.ChangePasswordReqDTO) error
	CompleteTeacherProfile(teacherId uint, dto userrequestdto.CompleteTeacherProfileReqDTO) error
	BlockUser(blockById uint, dto userrequestdto.BlockUserReqDTO) error
	UnBlockUser(dto userrequestdto.UnBlockUserReqDTO) error
	GetUsers(page, limit int) ([]userentity.User, error)
	GetUsersPaginationMetadata(currentPage, limit int) (*basetype.PaginationMetadata, error)
	FindUserById(id uint) (*userentity.User, error)
	FindUserByPhone(phone string) (*userentity.User, error)
	UploadResumeFile(fileHeader *multipart.FileHeader) (string, error)
	UploadProfileImage(fileHeader *multipart.FileHeader) (string, error)
}
