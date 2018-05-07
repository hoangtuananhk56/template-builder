package person

import (
	"net/http"
	"template-builder/tbs/o/org/person"
	"template-builder/tbs/x/web"
)

type PersonServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewPersonServer() *PersonServer {
	var s = &PersonServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/search", s.HandleAllTemp)
	return s
}

func (s *PersonServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var temp person.Person
	s.MustDecodeBody(r, &temp)
	// var u = session.MustAuthScope(r)
	// temp.UserID = u.ID
	web.AssertNil(temp.Create())

	s.SendData(w, temp)
}
func (s *PersonServer) mustGetTemp(r *http.Request) *person.Person {
	var id = r.URL.Query().Get("id")
	var u, err = person.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *PersonServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newTemp person.Person
	s.MustDecodeBody(r, &newTemp)
	var u = s.mustGetTemp(r)

	web.AssertNil(u.Update(&newTemp))

	s.SendData(w, nil)
}

func (s *PersonServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetTemp(r)
	s.SendData(w, u)
}

func (s *PersonServer) HandleAllTemp(w http.ResponseWriter, r *http.Request) {
	var res, err = person.GetAll(map[string]interface{}{})
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}
