package auth

import (
	"net/http"
	"template-builder/tbs/api/auth/session"
	"template-builder/tbs/o/org/user"
)

func (s *AuthServer) handleMySettings(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	me, _ := user.GetByID(u.UserID)
	s.SendData(w, map[string]interface{}{
		"me": me,
	})
}
