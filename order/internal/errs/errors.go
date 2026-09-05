package errs

import "errors"

var ErrorOrderNotFound = errors.New("order not found")

var ErrorConflict = errors.New("conflict: order paid")

var ErrorUnprocessableEntity = errors.New("unprocessable entity: the parts not found in order")
