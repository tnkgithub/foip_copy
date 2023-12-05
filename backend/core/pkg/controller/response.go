package controller

type ErrorResponse struct {
	Code    int
	Message string
}

type GetConnectionTokenRespone struct {
	Token string `json:"token"`
}
