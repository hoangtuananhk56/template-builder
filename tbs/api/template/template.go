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

func NewTemplateServer() *TemplateServer {
	var s = &TemplateServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/search", s.HandleAllTemp)
	return s
}

func (s *TemplateServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var temp template.Template
	s.MustDecodeBody(r, &temp)
	// var u = session.MustAuthScope(r)
	// temp.UserID = u.ID
	web.AssertNil(temp.Create())

	s.SendData(w, temp)
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

type SetType struct {
	Mode string `json:"mode"`
}

func (s *TemplateServer) HandleAllTemp(w http.ResponseWriter, r *http.Request) {
	var St SetType
	s.MustDecodeBody(r, &St)
	var res, err = template.GetAll(map[string]interface{}{
		"mode": St.Mode,
	})
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}
