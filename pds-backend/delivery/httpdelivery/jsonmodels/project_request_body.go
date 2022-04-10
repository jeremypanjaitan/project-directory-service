package jsonmodels

type ProjectRequestBody struct {
	Title       string `json:"title"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Story       string `json:"story"`
	CategoryID  uint   `json:"categoryId"`
}
