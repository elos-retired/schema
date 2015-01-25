package schema

import (
	"errors"
)

var UndefinedKindError = errors.New("Error: undefined kind")
var UndefinedLinkError = errors.New("Error: undefined link")
var UndefinedLinkKindError = errors.New("Error: undefined link kind")
var InvalidSchemaError = errors.New("Error: invalid schema")
