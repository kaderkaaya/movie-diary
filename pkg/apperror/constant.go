package apperror

var (
	ErrPasswordEmpty      = New(1001, "Password cant be empty")
	ErrEmailEmpty         = New(1002, "Email cant be empty")
	ErrUserEmpty          = New(1003, "User cant be empty")
	ErrEmailAlreadyExists = New(1004, "Email already exists")
)
