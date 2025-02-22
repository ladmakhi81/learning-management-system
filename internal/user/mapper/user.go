package usermapper

import (
	userresponsedto "github.com/ladmakhi81/learning-management-system/internal/user/dto/response"
	userentity "github.com/ladmakhi81/learning-management-system/internal/user/entity"
)

type UserMapper struct{}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (m UserMapper) MapUserToUserResponseDTO(user *userentity.User) userresponsedto.UserResDTO {
	dto := userresponsedto.NewUserResDTO()
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.ID = user.ID
	dto.CreatedAt = &user.CreatedAt
	dto.UpdatedAt = &user.UpdatedAt
	dto.DeletedAt = user.DeletedAt
	dto.Phone = user.Phone
	return dto
}

func (m UserMapper) MapUsersToUsersResponseDTO(users []userentity.User) []userresponsedto.UserResDTO {
	dtos := make([]userresponsedto.UserResDTO, len(users))
	for index, user := range users {
		dtos[index] = m.MapUserToUserResponseDTO(&user)
	}
	return dtos
}
