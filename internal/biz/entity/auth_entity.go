package entity

type LoginRequest struct {
	Phone    string `json:"phone"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Captcha         string `json:"captcha"`
}
