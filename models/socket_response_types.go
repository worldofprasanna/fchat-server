package models

// ResponseType - used to categorize the websocket messages
type ResponseType string

const(
	// UserRegisteredSuccessfully - User Registered Successfully
	UserRegisterationSuccessfully ResponseType = "USER_REGISTERED_SUCCESSFULLY"
	UserRegisteredFailed ResponseType = "USER_REGISTERED_SUCCESSFULLY"
)