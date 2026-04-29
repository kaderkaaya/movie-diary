package apperror

var (
	ErrPasswordEmpty      = CreateError(1001, "Password cant be empty")
	ErrEmailEmpty         = CreateError(1002, "Email cant be empty")
	ErrUserEmpty          = CreateError(1003, "User cant be empty")
	ErrEmailAlreadyExists = CreateError(1004, "Email already exists")
	ErrPasswordHashError  = CreateError(1005, "Password hash error")
)
