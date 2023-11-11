package vo

type Userinfo struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Token       string `json:"token"`
}
