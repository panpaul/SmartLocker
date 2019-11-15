package e

/*
	  0 Success
	1XX (Program/Http)Error
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
	UploadFailed       = 102
	RedisError         = 103
	JsonUnmarshalError = 104
	NotFound           = 200
	NoMoreLocker       = 201
	CacheNotFound      = 202
	Unauthorized       = 300
	PermissionDenied   = 301
	InvalidParams      = 400
	JWTNotAToken       = 401
	JWTInvalid         = 402
	FileTypeMismatch   = 403
	UsernameInvalid    = 404
	JWTOutOfTime       = 500
	RegisterDuplicated = 600
)
