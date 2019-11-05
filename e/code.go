package e

/*
	  0 Success
	1XX (Program)Error
	2XX Not found
	3XX Permission related
	4XX Format mismatch
	5XX Time issues
	6XX Duplicate
*/
const (
	Success            = 0
	InternalError      = 100
	RegistrationFailed = 101
	NotFound           = 200
	NoMoreLocker       = 201
	Unauthorized       = 300
	PermissionDenied   = 301
	InvalidParams      = 400
	JWTNotAToken       = 401
	JWTInvalid         = 402
	JWTOutOfTime       = 500
	RegisterDuplicated = 600
)
