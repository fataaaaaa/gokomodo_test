package model

type FilterListModel struct {
	Limit  int
	Offset int
	Search string
	UserId string
}

type BaseResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

type TotalRow struct {
	Total int `db:"total"`
}
