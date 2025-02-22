package userservice

import (
	"net/http"
	"time"

	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	baseutil "github.com/ladmakhi81/learning-management-system/internal/base/util"
	userconstant "github.com/ladmakhi81/learning-management-system/internal/user/constant"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo usercontractor.UserRepository
}

func NewUserServiceImpl(
	userRepo usercontractor.UserRepository,
) UserServiceImpl {
	return UserServiceImpl{
		userRepo: userRepo,
	}
}

func (svc UserServiceImpl) CreateUser(dto userrequestdto.CreateUserReqDTO) (*userentity.User, error) {
	duplicatedPhone, duplicatedPhoneErr := svc.FindUserByPhone(dto.Phone)
	if duplicatedPhoneErr != nil {
		return nil, duplicatedPhoneErr
	}
	if duplicatedPhone != nil {
		return nil, baseerror.NewClientErr(
			userconstant.USER_PHONE_DUPLICATED,
			http.StatusConflict,
		)
	}
	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		return nil, baseerror.NewServerErr(
			hashedPasswordErr.Error(),
			"UserServiceImpl.CreateUser",
		)
	}
	user := userentity.NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Phone,
		string(hashedPassword),
	)
	if createErr := svc.userRepo.CreateUser(user); createErr != nil {
		return nil, baseerror.NewServerErr(
			createErr.Error(),
			"UserServiceImpl.CreateUser",
		)
	}
	return user, nil
}

func (svc UserServiceImpl) VerifyTeacherByAdmin(adminId uint, dto userrequestdto.VerifyTeacherByAdminReqDTO) error {
	teacher, teacherErr := svc.FindUserById(dto.TeacherId)
	if teacherErr != nil {
		return teacherErr
	}
	if !teacher.IsProfileComplete {
		return baseerror.NewClientErr(
			userconstant.TEACHER_NOT_COMPLETE_PROFILE,
			http.StatusBadRequest,
		)
	}
	now := time.Now()
	teacher.VerifiedByID = &adminId
	teacher.VerifiedDate = &now
	teacher.IsVerified = true
	if updateErr := svc.userRepo.EditUser(teacher); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.VerifyTeacherByAdmin",
		)
	}
	return nil
}

func (svc UserServiceImpl) ChangePassword(executorId uint, dto userrequestdto.ChangePasswordReqDTO) error {
	user, userErr := svc.FindUserById(dto.UserID)
	if userErr != nil {
		return userErr
	}
	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		return baseerror.NewServerErr(
			hashedPasswordErr.Error(),
			"UserServiceImpl.ChangePassword",
		)
	}
	now := time.Now()
	user.Password = string(hashedPassword)
	user.PasswordChangeBy = &executorId
	user.PasswordChangeDate = &now
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.ChangePassword",
		)
	}
	return nil
}

func (svc UserServiceImpl) UpdateBaseInformation(userId uint, dto userrequestdto.UpdateBaseInformationReqBody) error {
	user, userErr := svc.FindUserById(userId)
	if userErr != nil {
		return userErr
	}
	if dto.FirstName != "" {
		user.FirstName = dto.FirstName
	}
	if dto.LastName != "" {
		user.LastName = dto.LastName
	}
	if dto.Phone != "" {
		duplicatedPhone, duplicatedPhoneErr := svc.FindUserByPhone(dto.Phone)
		if duplicatedPhoneErr != nil {
			return duplicatedPhoneErr
		}
		if duplicatedPhone != nil && duplicatedPhone.ID != userId {
			return baseerror.NewClientErr(
				userconstant.USER_PHONE_DUPLICATED,
				http.StatusConflict,
			)
		}
		user.Phone = dto.Phone
	}
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.UpdateBaseInformation",
		)
	}
	return nil
}

func (svc UserServiceImpl) CompleteTeacherProfile(teacherId uint, dto userrequestdto.CompleteTeacherProfileReqDTO) error {
	user, userErr := svc.FindUserById(teacherId)
	if userErr != nil {
		return userErr
	}
	now := time.Now()
	user.Bio = dto.Bio
	user.Email = dto.Email
	user.NationalID = dto.NationalID
	user.ProfileImage = dto.ProfileImage
	user.ResumeFile = dto.ResumeFile
	user.IsProfileComplete = true
	user.CompleteProfileDate = &now
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.CompleteTeacherProfile",
		)
	}
	return nil
}

func (svc UserServiceImpl) BlockUser(blockById uint, dto userrequestdto.BlockUserReqDTO) error {
	user, userErr := svc.FindUserById(*dto.UserID)
	if userErr != nil {
		return userErr
	}
	now := time.Now()
	user.BlockByID = &blockById
	user.BlockDate = &now
	user.BlockReason = dto.BlockReason
	user.IsBlock = true
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.BlockUser",
		)
	}
	return nil
}

func (svc UserServiceImpl) UnBlockUser(dto userrequestdto.UnBlockUserReqDTO) error {
	user, userErr := svc.FindUserById(*dto.UserID)
	if userErr != nil {
		return userErr
	}
	user.BlockByID = nil
	user.BlockDate = nil
	user.BlockReason = ""
	user.IsBlock = false
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.UnBlockUser",
		)
	}
	return nil
}

func (svc UserServiceImpl) GetUsers(page, limit int) ([]userentity.User, error) {
	users, usersErr := svc.userRepo.GetUsers(page, limit)
	if usersErr != nil {
		return nil, baseerror.NewServerErr(
			usersErr.Error(),
			"UserServiceImpl.GetUsers",
		)
	}
	return users, nil
}

func (svc UserServiceImpl) GetUsersPaginationMetadata(currentPage, limit int) (*basetype.PaginationMetadata, error) {
	totalCount, countErr := svc.userRepo.GetUsersCount()
	if countErr != nil {
		return nil, baseerror.NewServerErr(
			countErr.Error(),
			"UserServiceImpl.GetUsersCount",
		)
	}
	return basetype.NewPaginationMetadata(
		currentPage,
		baseutil.CalculateTotalPage(totalCount, limit),
		totalCount,
	), nil
}

func (svc UserServiceImpl) FindUserById(id uint) (*userentity.User, error) {
	user, userErr := svc.userRepo.FindUserById(id)
	if userErr != nil {
		return nil, baseerror.NewServerErr(
			userErr.Error(),
			"UserRepositoryImpl.GetUserById",
		)
	}
	if user == nil {
		return nil, baseerror.NewClientErr(
			userconstant.USER_NOT_FOUND_ID,
			http.StatusNotFound,
		)
	}
	return user, nil
}

func (svc UserServiceImpl) FindUserByPhone(phone string) (*userentity.User, error) {
	user, userErr := svc.userRepo.FindUserByPhone(phone)
	if userErr != nil {
		return nil, baseerror.NewServerErr(
			userErr.Error(),
			"UserServiceImpl.FindUserByPhone",
		)
	}
	return user, nil
}
