package template

import (
	"net/http"
	"template-builder/tbs/o/org/template"
	"template-builder/tbs/x/web"
)

type TemplateServer struct {
	web.JsonServer
	*http.ServeMux
}

func  NewTemplateServer()*TemplateServer{
	var s =&TemplateServer{
		ServeMux:http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/search", s.HandleAllTemp)
	return s
}

func (s *TemplateServer) HandleCreate(w http.ResponseWriter, r *http.Request){
	var u template.Template
	s.MustDecodeBody(r, &u)
	web.AssertNil(u.Create())

	s.SendData(w, u)
	//cu.OnUserCreated(u.ID)
}
func (s *TemplateServer) mustGetTemp(r *http.Request) *template.Template {
	var id = r.URL.Query().Get("id")
	var u, err = template.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *TemplateServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newTemp template.Template
	s.MustDecodeBody(r, &newTemp)
	var u = s.mustGetTemp(r)

	web.AssertNil(u.Update(&newTemp))

	s.SendData(w, nil)
}

func (s *TemplateServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetTemp(r)
	s.SendData(w, u)
}

func (s *TemplateServer) HandleAllTemp(w http.ResponseWriter, r *http.Request) {
	var res, err = template.GetAll()
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}