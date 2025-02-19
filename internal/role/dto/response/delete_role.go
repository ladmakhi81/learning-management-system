package roleresponsedto

type DeleteRoleRes struct {
	Message string `json:"message"`
}

func NewDeleteRoleRes(message string) DeleteRoleRes {
	return DeleteRoleRes{
		Message: message,
	}
}
