package controller

import (
	"net/http"

	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/business"
	"github.com/mberliner/gobase/09-webapp_modular/abm_mongodb/model"
)

const sessionCookie string = "session"

var dbSessions = map[string]string{}

func getUser(res http.ResponseWriter, req *http.Request) model.User {
	var u model.User

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	// Si existe lo tomo de la sesion
	if usu, ok := dbSessions[c.Value]; ok {
		u = business.UserB.BuscaPorUsuario(usu)
	}
	return u
}

func estaLogueado(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	usu := dbSessions[c.Value]

	u := business.UserB.BuscaPorUsuario(usu)
	if u.Error != nil {
		return false
	}
	return true
}
