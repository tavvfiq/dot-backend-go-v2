package domain

type Response struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
