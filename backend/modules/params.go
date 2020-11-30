package modules

type SignUpParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RateParam struct {
	Id      int     `json:"id" db:"id"`
	Rate    float32 `json:"rate" db:"rating"`
	RateNum int     `db:"rate_num"`
}

type CommentParam struct {
	Comment string  `json:"comment" db:"comment_string"`
	Rate    float32 `json:"rate" db:"rate"`
}
