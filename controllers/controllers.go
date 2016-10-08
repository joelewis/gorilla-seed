package controllers

import (
	"fmt"
	models "github.com/joelewis/gorilla-seed/models"
	utils "github.com/joelewis/gorilla-seed/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := utils.GetSession("loginSession", w, r)
	if err != nil {
		// error writing in http is done in utils.getSession()
		return
	}

	user := models.GetSessionUser(session)
	if user != nil {
		fmt.Fprintf(w, "Welcome, %s", user.Username)
		return
	}
	fmt.Fprintf(w, "Welcome!")
}
