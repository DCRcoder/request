package request

import "errors"

var (
	// OBJNotCanSet obj
	OBJNotCanSet = errors.New("json obj must be pointer")
)
