package main

import (
	// "fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	models "github.com/joelewis/guerilla-seed/models"
	utils "github.com/joelewis/guerilla-seed/utils"
	"log"
	"net/http"
	"os"
	"time"
)

// config vars
var SERVER_PORT = "5000"
var SERVER_HOST = "127.0.0.1"

// refer http://jinzhu.me/gorm/database.html#connecting-to-a-database for other db clients
var DB_URL = "./seed.sqlite3" // use absolute path
var DB_DIALECT = "sqlite3"

var SESSION_FILESTORE_URL = "/Users/username/home/sessions" // use absolute path
var SESSION_SECRET_KEY = "Some secret phrase, that shouldn't be in version control"

// absolute path to the directory that contains static files.
var STATIC_DIR = "./public"

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	actionMap := make(map[string]interface{})
	actionMap["inittables"] = models.InitTables

	if len(os.Args) > 1 {
		action := os.Args[1]
		log.Printf("Executing %s", action)
		actionMap[action].(func(string, string))(DB_URL, DB_DIALECT)
		return
	} else {
		log.Printf("starting server @ " + SERVER_HOST + ":" + SERVER_PORT)
	}

	// init db
	db := models.InitDb(DB_URL, DB_DIALECT)
	defer db.Close()

	utils.InitSessionStore(SESSION_FILESTORE_URL, SESSION_SECRET_KEY)

	// register routes
	router := RegisterRoutes()
	srv := &http.Server{
		Handler:      Log(router),
		Addr:         SERVER_HOST + ":" + SERVER_PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
