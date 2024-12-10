package types

import "errors"

var ErrNotFound = errors.New("item not found")
var ErrItemExists = errors.New("item already exists")
var ErrOperationNotPermitted = errors.New("operation not permitted")
