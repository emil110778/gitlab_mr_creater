package cmd

import "errors"

var (
	ErrConfigure = errors.New("please configure")
	InternalErr  = errors.New("internal error")
)
