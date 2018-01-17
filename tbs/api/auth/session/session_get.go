package session

import (
	"template-builder/tbs/o/auth/session"
	"template-builder/tbs/x/web"
)

const (
	errReadSessonFailed   = web.InternalServerError("read session failed")
	errSessionNotFound    = web.Unauthorized("session not found")
	errUnauthorizedAccess = web.Unauthorized("unauthorized access")
)

func Get(sessionID string) (*session.Session, error) {
	var s, err = session.GetByID(sessionID)
	if err != nil {
		sessionLog.Error(err)
		return nil, errReadSessonFailed
	}
	if s == nil {
		return nil, errSessionNotFound
	}
	return s, nil
}
