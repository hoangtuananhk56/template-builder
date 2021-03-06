package api

import (
	"net/http"
	"template-builder/tbs/api/auth"
	"template-builder/tbs/api/person"
	"template-builder/tbs/api/template"
	"template-builder/tbs/api/user"
	"template-builder/tbs/x/web"
)

type ApiServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewApiServer() *ApiServer {
	var s = &ApiServer{
		ServeMux: http.NewServeMux(),
	}

	s.Handle("/user/", http.StripPrefix("/user", user.NewUserServer()))
	s.Handle("/auth/", http.StripPrefix("/auth", auth.NewAuthServer()))
	s.Handle("/template/", http.StripPrefix("/template", template.NewTemplateServer()))
	s.Handle("/per/", http.StripPrefix("/per", person.NewPersonServer()))
	return s
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}
