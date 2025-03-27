package e

// Application wide codes
const (
	ErrCodeAuto            = 0
	ErrCodeInternalService = 666
)

// 400 errors
const (
	// ErrInvalidRequest : when post body, query param, or path
	// param is invalid, or any post body validation error is encountered
	ErrInvalidRequest int = 400000 + iota

	// ErrDecodeRequestBody : error when decode the request body
	ErrDecodeRequestBody

	// ErrValidateRequest : error when validating the request
	ErrValidateRequest

	// ErrSaveUserDetails : error when saving user details
	ErrSaveUserDetails

	// ErrSaveUserDetails : error when saving user details
	ErrInvaliPassword

	// ErrTokenNotGenerated : error when getting author by id
	ErrTokenNotGenerated
)

// 404 errors
const (
	// ErrResourceNotFound : when no record corresponding to the requested id is found in the DB
	ErrResourceNotFound int = 404000 + iota
)

// 500 errors
const (
	// ErrInternalServer : the default error, which is unexpected from the developers
	ErrInternalServer int = 500000 + iota

	// ErrExecuteSQL : when execute the sql, meet unexpected error
	ErrExecuteSQL
)
