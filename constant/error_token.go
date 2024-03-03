package constant

import (
	"fmt"
)

// Error JWT
var (
	ErrInvalidOrEmptyToken       = fmt.Errorf("unauthorized")
	ErrorHttpInvalidServiceToken = fmt.Errorf("invalid service token")
	ErrTokenIsExpired            = fmt.Errorf("token is expired")
)
