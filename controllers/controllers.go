package controllers

import (
	"fmt"
	models "github.com/joelewis/guerilla-seed/models"
	utils "github.com/joelewis/guerilla-seed/utils"
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
