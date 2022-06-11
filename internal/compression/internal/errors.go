package internal

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrNotAValidRequest = errors.BadRequest(
		"Not A Valid Request",
		"Compression was enabled ")
)
