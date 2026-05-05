package domain

import "errors"

// 要件 5.7: Validation と NotFound を呼び出し側が区別できるようにする。
var (
	ErrNotFound   = errors.New("task not found")
	ErrValidation = errors.New("validation error")
)
