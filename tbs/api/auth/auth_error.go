package auth

import (
	"template-builder/tbs/x/web"
)

const (
	errUserNotFound = web.Unauthorized("USER_NOT_FOUND")
)
