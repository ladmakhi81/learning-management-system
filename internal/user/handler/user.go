package userhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	baseutil "github.com/ladmakhi81/learning-management-system/internal/base/util"
	userconstant "github.com/ladmakhi81/learning-management-system/internal/user/constant"
	usercontractor "github.com/ladmakhi81/learning-management-system/internal/user/contractor"
	userrequestdto "github.com/ladmakhi81/learning-management-system/internal/user/dto/request"
	userresponsedto "github.com/ladmakhi81/learning-management-system/internal/user/dto/response"
	usermapper "github.com/ladmakhi81/learning-management-system/internal/user/mapper"
)

type UserHandler struct {
	userSvc    usercontractor.UserService
	userMapper usermapper.UserMapper
}

func NewUserHandler(
	userSvc usercontractor.UserService,
	userMapper usermapper.UserMapper,
) UserHandler {
	return UserHandler{
		userSvc:    userSvc,
		userMapper: userMapper,
	}
}

func (h UserHandler) CreateUser(ctx *gin.Context) (*basehandler.Response, error) {
	dto := userrequestdto.NewCreateUserReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	user, userErr := h.userSvc.CreateUser(*dto)
	if userErr != nil {
		return nil, userErr
	}
	mappedUser := h.userMapper.MapUserToUserResponseDTO(user)
	res := userresponsedto.NewCreateUserResDTO(mappedUser)
	return basehandler.NewResponse(res, http.StatusCreated), nil
}

func (h UserHandler) VerifyTeacherByAdmin(ctx *gin.Context) (*basehandler.Response, error) {
	dto := userrequestdto.NewVerifyTeacherByAdminReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	// TODO: replace by real admin id
	adminId := uint(1)
	if verifyErr := h.userSvc.VerifyTeacherByAdmin(adminId, *dto); verifyErr != nil {
		return nil, verifyErr
	}
	res := userresponsedto.NewVerifyTeacherByAdminResDTO("Teacher Verify Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) ChangePassword(ctx *gin.Context) (*basehandler.Response, error) {
	// TODO: replace this id with real executorId from token
	executorId := uint(1)
	dto := userrequestdto.NewChangePasswordReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if changeErr := h.userSvc.ChangePassword(executorId, *dto); changeErr != nil {
		return nil, changeErr
	}
	res := userresponsedto.NewChangePasswordResDTO("Password Changed Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) UpdateBaseInformation(ctx *gin.Context) (*basehandler.Response, error) {
	// TODO: replace real id instead of fake one from token
	userId := uint(1)
	dto := userrequestdto.NewUpdateBaseInformationReqBody()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if err := h.userSvc.UpdateBaseInformation(userId, *dto); err != nil {
		return nil, err
	}
	res := userresponsedto.NewUpdateBaseInformationResDTO("Basic Information Updated Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) CompleteTeacherProfile(ctx *gin.Context) (*basehandler.Response, error) {
	teacherIdParam := ctx.Param("teacher-id")
	teacherId, teacherIdErr := strconv.Atoi(teacherIdParam)
	if teacherIdErr != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_TEACHER_ID,
			http.StatusBadRequest,
		)
	}
	dto := userrequestdto.NewCompleteTeacherProfileReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if err := h.userSvc.CompleteTeacherProfile(uint(teacherId), *dto); err != nil {
		return nil, err
	}
	res := userresponsedto.NewCompleteTeacherProfileResDTO("Teacher Profile Complete Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) BlockUser(ctx *gin.Context) (*basehandler.Response, error) {
	// TODO: replace id with real id from token
	blockById := uint(1)
	dto := userrequestdto.NewBlockUserReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if err := h.userSvc.BlockUser(blockById, *dto); err != nil {
		return nil, err
	}
	res := userresponsedto.NewBlockUserResDTO("User Block Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) UnBlockUser(ctx *gin.Context) (*basehandler.Response, error) {
	dto := userrequestdto.NewUnBlockUserReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if err := h.userSvc.UnBlockUser(*dto); err != nil {
		return nil, err
	}
	res := userresponsedto.NewUnBlockUserResDTO("User UnBlock Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) GetUsers(ctx *gin.Context) (*basehandler.Response, error) {
	paginationParam := baseutil.ExtraPaginationData(ctx.Query("page"), ctx.Query("limit"))
	users, usersErr := h.userSvc.GetUsers(paginationParam.Page, paginationParam.Limit)
	if usersErr != nil {
		return nil, usersErr
	}
	mappedUsers := h.userMapper.MapUsersToUsersResponseDTO(users)
	pagination, paginationErr := h.userSvc.GetUsersPaginationMetadata(paginationParam.Page, paginationParam.Limit)
	if paginationErr != nil {
		return nil, paginationErr
	}
	res := userresponsedto.NewGetUsersRes(mappedUsers, *pagination)
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h UserHandler) UploadTeacherResume(ctx *gin.Context) (*basehandler.Response, error) {
	fileHeader, fileHeaderErr := ctx.FormFile("resume")
	if fileHeaderErr != nil {
		return nil, baseerror.NewClientErr(
			userconstant.NOT_FOUND_RESUME_FILE,
			http.StatusBadRequest,
		)
	}
	if fileHeader.Header.Get("Content-Type") != "application/pdf" {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_FORMAT_RESUME,
			http.StatusBadRequest,
		)
	}
	fileName, fileErr := h.userSvc.UploadResumeFile(fileHeader)
	if fileErr != nil {
		return nil, fileErr
	}
	res := userresponsedto.NewUploadResumeResDTO(fileName)
	return basehandler.NewResponse(res, http.StatusCreated), nil
}

func (h UserHandler) UploadProfileImage(ctx *gin.Context) (*basehandler.Response, error) {
	fileHeader, fileHeaderErr := ctx.FormFile("image")
	if fileHeaderErr != nil {
		return nil, baseerror.NewClientErr(
			userconstant.NOT_FOUND_PROFILE_IMAGE,
			http.StatusBadRequest,
		)
	}
	mediaType := fileHeader.Header.Get("Content-Type")
	if isImage := strings.HasPrefix(mediaType, "image/"); !isImage {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_FORMAT_PROFILE,
			http.StatusBadRequest,
		)
	}
	filename, fileErr := h.userSvc.UploadProfileImage(fileHeader)
	if fileErr != nil {
		return nil, fileErr
	}
	res := userresponsedto.NewUploadProfileImageResDTO(filename)
	return basehandler.NewResponse(res, http.StatusCreated), nil
}

func (h UserHandler) AssignRole(ctx *gin.Context) (*basehandler.Response, error) {
	// TODO: replace with real id from token
	executorId := uint(1)
	dto := userrequestdto.NewAssignRoleReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(
			userconstant.INVALID_REQUEST_BODY,
			http.StatusBadRequest,
		)
	}
	if err := h.userSvc.AssignRole(executorId, *dto); err != nil {
		return nil, err
	}
	res := userresponsedto.NewAssignRoleResDTO("Role Assigned Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}
