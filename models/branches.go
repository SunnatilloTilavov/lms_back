package models

type Branche struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Address       string  `json:"address"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetAllBranchesResponse struct {
	Branches  []Branche `json:"branche"`
	Count int64 `json:"count"`
}
