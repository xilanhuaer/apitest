package response

type Page struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}
