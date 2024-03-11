package models

type Group struct {
	Id          string  `json:"id"`
	GroupId          string  `json:"groupid"`
	BrancheId          string  `json:"brancheid"`
	Teacher        string  `json:"teacher"`
	Type       string  `json:"type"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetAllGroupsResponse struct {
	Groups  []Group `json:"branche"`
	Count int64 `json:"count"`
}


