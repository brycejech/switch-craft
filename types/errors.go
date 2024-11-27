package types

import "errors"

var ErrNotFound = errors.New("Item not found")
var ErrItemExists = errors.New("Item already exists")
var ErrOperationNotPermitted = errors.New("Operation not permitted")
