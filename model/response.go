package model

type DefaultResponse struct {
	Message string `json:"message" binding:"required"`
}

type DataResponse struct {
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data" binding:"required"`
}
