package types

// CreateMessageRequest ...
type CreateMessageRequest struct {
	Recipient int64  `json:"recipient"`
	Content   string `json:"content"`
}

// IsValid ...
func (req CreateMessageRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Content == "" {
		errs = append(errs, RequiredError("Content is required!"))
	}

	if req.Recipient == 0 {
		errs = append(errs, RequiredError("Recipient not set!"))
	}

	return len(errs) == 0, errs
}
