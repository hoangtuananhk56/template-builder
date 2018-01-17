package user

import (
	"net/http"
	"template-builder/tbs/api/auth/session"
	"template-builder/tbs/o/org/user"
	"template-builder/tbs/x/web"
	"time"
)

type UserServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewUserServer() *UserServer {
	var s = &UserServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/find", s.HandleFindByName)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllUser)
	s.HandleFunc("/count", s.HandleCountUser)
	return s
}

func (s *UserServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u user.User
	s.MustDecodeBody(r, &u)
	web.AssertNil(u.Create())

	s.SendData(w, u)
	//cu.OnUserCreated(u.ID)
}

func (s *UserServer) mustGetUser(r *http.Request) *user.User {
	var id = r.URL.Query().Get("id")
	var u, err = user.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *UserServer) mustFindUser(r *http.Request) *user.User {
	var username = r.URL.Query().Get("username")
	var u, err = user.GetByUsername(username)
	web.AssertNil(err)
	return u
}

func (s *UserServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	s.MustDecodeBody(r, &newUser)
	var u = s.mustGetUser(r)

	web.AssertNil(u.Update(&newUser))

	s.SendData(w, nil)
}

func (s *UserServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	s.SendData(w, u)
}

func (s *UserServer) HandleFindByName(w http.ResponseWriter, r *http.Request) {
	var u = s.mustFindUser(r)
	s.SendData(w, u)
}

func (s *UserServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	web.AssertNil(user.MarkDelete(u.ID))
	s.Success(w)
}

func (s *UserServer) HandleAllUser(w http.ResponseWriter, r *http.Request) {
	var res, err = user.GetAll()
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}

type Date struct {
	Year  int `json: "year"`
	Month int `json: "month"`
	Day   int `json: "day"`
	Hour  int `json: "hour"`
	Min   int `json: "min"`
}

type DateFillter struct {
	Start Date `json: "start"`
	End   Date `json: "end"`
}

func (s *UserServer) HandleCountUser(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var x_user, err = user.GetByID(u.UserID)
	var df DateFillter
	s.MustDecodeBody(r, &df)

	var start = time.Date((df.Start.Year), time.Month((df.Start.Month)), (df.Start.Day), (df.Start.Hour), (df.Start.Min), 0, 0, time.Local)
	var end = time.Date((df.End.Year), time.Month((df.End.Month)), (df.End.Day), (df.End.Hour), (df.End.Min), 0, 0, time.Local)
	var cuser int
	if x_user.Role == "admin" {
		cuser = user.CountUser(start, end)
	}
	s.SendError(w, err)
	s.SendData(w, map[string]interface{}{"count": cuser})
}
