package types

// APIResponse ...
type APIResponse struct {
	Status int         `json:"status"`
	Errors []string    `json:"errors"`
	Result interface{} `json:"results"`
}
