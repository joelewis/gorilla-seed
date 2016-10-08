package utils

import (
    "log"
    "math/rand"
    "time"
    "net/http"
    "github.com/gorilla/sessions"
)
var Store *sessions.FilesystemStore

func CheckErr(e error, msg string) {
    if e != nil {
        log.Printf("%s %s", msg, e)
    }
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func InitSessionStore(SESSION_FILESTORE_URL, SESSION_SECRET_KEY string) *sessions.FilesystemStore {
    Store = sessions.NewFilesystemStore(SESSION_FILESTORE_URL, []byte(SESSION_SECRET_KEY))
    return Store
}

func GetStore() *sessions.FilesystemStore {
    return Store
}

func GetSession(name string, w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
    session, err := Store.Get(r, name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return nil, err
    }
    return session, nil
}
