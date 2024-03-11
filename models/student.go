package models

type Student struct {
	Id          string  `json:"id"`
	Full_Name        string  `json:"full_name"`
	Email     string  `json:"email"`
	Age     string  `json:"age"`
	Paid_sum     string  `json:"paid_sum"`
	Status     string  `json:"status"`
	Login     string  `json:"login"`
	Password    string  `json:"password"`
	Group_id    string  `json:"group_id"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetAllStudentsResponse struct {
	Students  []Student `json:"student"`
	Count int64 `json:"count"`
}

