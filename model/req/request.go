package req

type CreateContainerReq struct {
	Host  string `json:"host"`
	Image string `json:"image" binding:"required"`
}

type UserLoginReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
