package page

type Request struct {
	Page    int    `json:"page" default:"1" form:"page" binding:"required,min=1"  example:"1"`
	Limit   int    `json:"limit" default:"10" form:"limit" binding:"required,min=1" example:"10"`
	Keyword string `json:"keyword" default:"" form:"keyword" example:""`
}

type Response struct {
	List  any   `json:"list"`
	Total int64 `json:"total" default:"0"`
}
