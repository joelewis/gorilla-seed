package models

import (
    "log"
    "github.com/gorilla/sessions" 
)


func GetSessionUser(session *sessions.Session) *User {
    log.Printf("fetching session user")

    userId, ok := session.Values["user"]
    
    if ok {
        var user User
        return user.Get(userId.(string))
    }
    
    log.Printf("User id not found in session.")
    return nil
}
