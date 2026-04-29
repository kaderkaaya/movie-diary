package apperror

var (
	ErrPasswordEmpty              = CreateError(1001, "Password cant be empty")
	ErrEmailEmpty                 = CreateError(1002, "Email cant be empty")
	ErrUserEmpty                  = CreateError(1003, "User cant be empty")
	ErrEmailAlreadyExists         = CreateError(1004, "Email already exists")
	ErrPasswordHashError          = CreateError(1005, "Password hash error")
	ErrUserNotFound               = CreateError(1006, "User not found")
	ErrInvalidPassword            = CreateError(1007, "Invalid password")
	ErrInvalidToken               = CreateError(1008, "Invalid token")
	ErrUnexpectedSigningMethod    = CreateError(1009, "Unexpected signing method")
	ErrAuthorizationHeaderMissing = CreateError(1010, "Authorization header missing")
	ErrTokenEmpty                 = CreateError(1011, "Token cant be empty")
	ErrTokenNotFound              = CreateError(1012, "Token not found")
	ErrTokenExpired               = CreateError(1013, "Token expired")
)
