package userservice

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	baseutil "github.com/ladmakhi81/learning-management-system/internal/base/util"
	queuedto "github.com/ladmakhi81/learning-management-system/internal/queue/dto"
	queueservice "github.com/ladmakhi81/learning-management-system/internal/queue/service"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	userconstant "github.com/ladmakhi81/learning-management-system/internal/user/constant"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
	"github.com/nfnt/resize"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo        usercontractor.UserRepository
	config          *baseconfig.Config
	roleSvc         rolecontractor.RoleService
	pdfQueueService *queueservice.PDFQueueService
}

func NewUserServiceImpl(
	userRepo usercontractor.UserRepository,
	config *baseconfig.Config,
	roleSvc rolecontractor.RoleService,
	pdfQueueService *queueservice.PDFQueueService,
) UserServiceImpl {
	return UserServiceImpl{
		userRepo:        userRepo,
		config:          config,
		roleSvc:         roleSvc,
		pdfQueueService: pdfQueueService,
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

func (svc UserServiceImpl) UploadResumeFile(fileHeader *multipart.FileHeader) (string, error) {
	file, fileErr := fileHeader.Open()
	if fileErr != nil {
		return "", baseerror.NewServerErr(
			fileErr.Error(),
			"UserServiceImpl.UploadResumeFile",
		)
	}
	defer file.Close()
	destination := fmt.Sprintf("%s/resumes", svc.config.UploadDirectory)
	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return "", baseerror.NewServerErr(
			err.Error(),
			"UserServiceImpl.UploadResumeFile",
		)
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%d%s",
		time.Now().UnixMicro(),
		rand.Intn(10000000000),
		fileExt,
	)
	outputFile, outputErr := os.Create(path.Join(destination, filename))
	if outputErr != nil {
		return "", baseerror.NewServerErr(
			outputErr.Error(),
			"UserServiceImpl.UploadResumeFile",
		)
	}
	defer outputFile.Close()
	if _, err := io.Copy(outputFile, file); err != nil {
		return "", baseerror.NewServerErr(
			err.Error(),
			"UserServiceImpl.UploadResumeFile",
		)
	}
	svc.pdfQueueService.QueueService.Publish(queuedto.NewPDFCompressMessage(filename, destination))
	return filename, nil
}

func (svc UserServiceImpl) UploadProfileImage(fileHeader *multipart.FileHeader) (string, error) {
	file, fileErr := fileHeader.Open()
	if fileErr != nil {
		return "", baseerror.NewServerErr(
			fileErr.Error(),
			"UserServiceImpl.UploadProfileImage",
		)
	}
	defer file.Close()
	destination := fmt.Sprintf("%s/profiles", svc.config.UploadDirectory)
	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return "", baseerror.NewServerErr(
			err.Error(),
			"UserServiceImpl.UploadProfileImage",
		)
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%d%s",
		time.Now().UnixMicro(),
		rand.Intn(10000000000),
		fileExt,
	)
	decodedImage, decodedImageErr := svc.decodeProfileImage(file, fileExt)
	if decodedImageErr != nil {
		return "", baseerror.NewServerErr(
			decodedImageErr.Error(),
			"UserServiceImpl.UploadProfileImage",
		)
	}
	if decodedImage == nil {
		return "", baseerror.NewClientErr(
			userconstant.INVALID_FORMAT_PROFILE,
			http.StatusBadRequest,
		)
	}
	imageWidth := 200
	imageHeight := 0
	resizedImage := resize.Resize(uint(imageWidth), uint(imageHeight), decodedImage, resize.Lanczos2)
	outputFile, outputErr := os.Create(path.Join(destination, filename))
	if outputErr != nil {
		return "", baseerror.NewServerErr(
			outputErr.Error(),
			"UserServiceImpl.UploadProfileImage",
		)
	}
	defer outputFile.Close()
	if encodeErr := svc.encodeProfileImage(outputFile, resizedImage, fileExt); encodeErr != nil {
		return "", baseerror.NewServerErr(
			encodeErr.Error(),
			"UserServiceImpl.UploadProfileImage",
		)
	}
	return filename, nil
}

func (svc UserServiceImpl) decodeProfileImage(file multipart.File, fileExt string) (image.Image, error) {
	var decodedImage image.Image

	if fileExt == ".jpg" || fileExt == ".jpeg" {
		jpgFile, jpgErr := jpeg.Decode(file)
		if jpgErr != nil {
			return nil, jpgErr
		}
		decodedImage = jpgFile
	}

	if fileExt == ".png" {
		pngFile, pngErr := png.Decode(file)
		if pngErr != nil {
			return nil, pngErr
		}
		decodedImage = pngFile
	}

	return decodedImage, nil
}

func (svc UserServiceImpl) encodeProfileImage(outputFile *os.File, resizedImage image.Image, fileExt string) error {
	if fileExt == ".jpg" || fileExt == ".jpeg" {
		if err := jpeg.Encode(outputFile, resizedImage, nil); err != nil {
			return err
		}
	}
	if fileExt == ".png" {
		if err := png.Encode(outputFile, resizedImage); err != nil {
			return err
		}
	}
	return nil
}

func (svc UserServiceImpl) AssignRole(executorId uint, dto userrequestdto.AssignRoleReqDTO) error {
	user, userErr := svc.FindUserById(*&dto.UserID)
	if userErr != nil {
		return userErr
	}
	role, roleErr := svc.roleSvc.FindRoleById(dto.RoleID)
	if roleErr != nil {
		return roleErr
	}
	now := time.Now()
	user.AssignedRoleByID = &executorId
	user.AssignedRoleDate = &now
	user.RoleID = &role.ID
	if updateErr := svc.userRepo.EditUser(user); updateErr != nil {
		return baseerror.NewServerErr(
			updateErr.Error(),
			"UserServiceImpl.AssignRole",
		)
	}
	return nil
}
