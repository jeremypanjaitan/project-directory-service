package apperror

import "errors"

var ErrCreateUser = errors.New("error to create user")
var ErrRequiredField = errors.New("error required field")
var ErrDuplicateValue = errors.New("ERROR: duplicate key value violates unique constraint")
var ErrCredential = errors.New("wrong credential")
var ErrRoleNotFound = errors.New("role not found")
var ErrDivisionNotFound = errors.New("division not found")
var ErrUnableToLoadServiceAccount = errors.New("unable to load service account file")
var ErrDuplicateEmail = errors.New("email already exist")
var ErrRefreshToken = errors.New("invalid refresh token")
var ErrEmailAddress = errors.New("invalid email address")
var ErrFailedUpdateProject = errors.New("failed to update the project")
var ErrConstraintStringLength = errors.New("violate string length constraint")
