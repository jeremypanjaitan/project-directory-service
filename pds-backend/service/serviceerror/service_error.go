package serviceerror

import "errors"

var ErrInvalidAccess = errors.New("invalid access")
var ErrTokenNotFound = errors.New("token not found")
var ErrSigningMethod = errors.New("signing method invalid")
var ErrInvalidNumberOfSegments = errors.New("token contains an invalid number of segments")
