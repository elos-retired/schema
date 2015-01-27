package schema

import (
	"errors"
)

var (
	ErrUndefinedKind      = errors.New("undefined kind")
	ErrUndefinedLink      = errors.New("undefined link")
	ErrUndefinedLinkKind  = errors.New("undefined link mind")
	ErrInvalidSchema      = errors.New("invalid schema")
	ErrIncompatibleModels = errors.New("incompatible models")
)
