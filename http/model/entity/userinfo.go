package entity

type UserInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Token       string `json:"token"`
}
