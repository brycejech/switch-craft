package types

import "errors"

var ErrNotFound = errors.New("Item not found")
var ErrOperationNotPermitted = errors.New("Operation not permitted")
