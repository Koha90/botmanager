package domain

import "errors"

var ErrInvalidState error = errors.New("confirmed order cannot be canceled")
