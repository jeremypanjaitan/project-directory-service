package httperror

import "errors"

var ErrUnauthorized = errors.New("user is unauthorized")
var ErrFirebase = errors.New("errors in firebase")
