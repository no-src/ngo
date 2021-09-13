package _http

import (
	"errors"
	"github.com/no-src/ngo/res/lang"
)

var (
	MethodNotSupported = errors.New(lang.HTTP_Error_MethodNotSupported)
)
