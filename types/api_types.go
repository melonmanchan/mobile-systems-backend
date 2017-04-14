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

	ErrorUpdateProfileFailed = APIError{
		Message: "Updating user profile failed",
		Code:    "UPDATE_PROFILE_FAILED",
	}

	ErrorRegisterTutorFailed = APIError{
		Message: "Registering tutor profile failed",
		Code:    "REGISTER_TUTOR_FAILED",
	}

	// Tutorship errors
	ErrorNotTutor = APIError{
		Message: "The user is not a tutor!",
		Code:    "ERROR_NOT_TUTOR",
	}

	ErrorCreatingTutorship = APIError{
		Message: "Creating tutorship failed!",
		Code:    "ERROR_CREATING_TUTORSHIP",
	}

	ErrorGetTutorships = APIError{
		Message: "Getting tutorships failed!",
		Code:    "ERROR_GET_TUTORSHIPS",
	}

	// Message errors

	ErrorGetLatest = APIError{
		Message: "Getting latest messages failed!",
		Code:    "ERROR_GET_LATEST",
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

	ErrorGenericUserNotTutor = APIError{
		Message: "You need to be a tutor to do that!",
		Code:    "ERROR_GENERIC_USER_NOT_TUTOR",
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
