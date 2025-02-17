package api

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type PanRequestBody struct {
	PanNumber string `json:"panNumber"`
}

type PanResponse struct {
	Message string `json:"message"`
}

type GetPanResponse struct {
	Pans []PanDocument `json:"pans"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

var InternalErrorResponse = ErrorResponse{Message: "internal server error"}
