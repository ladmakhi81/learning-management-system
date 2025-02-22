package userrequestdto

type BlockUserReqDTO struct {
	BlockReason string `json:"reason"`
	UserID      *uint  `json:"userId"`
}

func NewBlockUserReqDTO() *BlockUserReqDTO {
	return new(BlockUserReqDTO)
}
