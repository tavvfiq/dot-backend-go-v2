package domain

type Pagination struct {
	Last    int `json:"last,omitempty"`
	Current int `json:"current"`
	Next    int `json:"next,omitempty"`
}
