package domain

type RegisterRequestData struct {
	Name string `json:"name"`
}

type RegisterSellerRequestData struct {
	UserId int64 `json:"user_id"`
}
