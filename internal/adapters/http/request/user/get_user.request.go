package request

type GetUserByIdRequest struct {
	ID int `uri:"id" binding:"required"`
}

type GetUserByIdResponse struct {
	ID int `json:"id"`
}
