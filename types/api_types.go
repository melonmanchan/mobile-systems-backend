package types

var (
	// Login errors

	ErrorLoginWrongPassword = APIError{
		Message: "Wrong password",
		Code:    "LOGIN_WRONG_PASSWORD",
	}

	ErrorLoginUserNotFound = APIError{
		Message: "User not found",
		Code:    "LOGIN_USER_NOT_FOUND",
	}

	// Register errors

	ErrorRegisterAlreadyExists = APIError{
		Message: "User already exists",
		Code:    "REGISTER_USER_EXISTS",
	}

	// Generic errors

	ErrorGenericRead = APIError{
		Message: "Error reading request",
		Code:    "ERROR_GENERIC_READ",
	}

	ErrorGenericTokenCreate = APIError{
		Message: "Error creating token",
		Code:    "ERROR_GENERIC_TOKEN_CREATE",
	}

	ErrorGenericTokenMissing = APIError{
		Message: "Authentication token missing",
		Code:    "ERROR_GENERIC_TOKEN_MISSING",
	}

	ErrorGenericTokenInvalid = APIError{
		Message: "Authentication token is invalid",
		Code:    "ERROR_GENERIC_TOKEN_INVALID",
	}

	ErrorGenericServer = APIError{
		Message: "Something went wrong!",
		Code:    "ERROR_GENERIC_SERVER",
	}

	ErrorGenericNotFound = APIError{
		Message: "Matching route not found for method and path",
		Code:    "ERROR_GENERIC_NOT_FOUND",
	}
)

// FormatError ...
func FormatError(message string) APIError {
	return APIError{
		Message: message,
		Code:    "ERROR_GENERIC_BAD_FORMAT",
	}
}

// RequiredError ...
func RequiredError(message string) APIError {
	return APIError{
		Message: message,
		Code:    "ERROR_GENERIC_REQUIRED",
	}
}

// APIResponse ...
type APIResponse struct {
	Status int         `json:"status"`
	Errors []APIError  `json:"errors"`
	Result interface{} `json:"results"`
}

// APIError ...
type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
