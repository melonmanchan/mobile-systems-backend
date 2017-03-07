package handlers

// APIResponse ...
type APIResponse struct {
	Errors []string    `json:"errors"`
	Result interface{} `json:"results"`
}
