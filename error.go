package limqa

import "errors"

var (
	// ErrNilBase :
	ErrNilBase = errors.New("Can not use <nil> in base")
	// ErrBaseNotConnected :
	ErrBaseNotConnected = errors.New("Base not connected")
)