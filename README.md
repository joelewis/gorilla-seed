# A quick seed for the golang-gorilla stack

## Getting Started

+ Clone the repo: `git clone https://github.com/joelewis/gorilla-seed`
+ Update the config variables in `main.go` (DB_URL, SERVER_PORT, etc.)
+ Rename the directory's name to your app's name. 
  `mv /path/to/gorilla-seed /path/to/my-app-name`
+ Update import paths.
  Eg. `import "path/to/gorilla-seed/models"` 
  to `import "path/to/my-app-name/models"`
+ Compile: `go build`
+ Initialize tables: `./my-app-name inittables`
+ Run server : `./my-app-name`

## What is inside?
Nothing much.
+ A developer friendly ORM - [gorm](https://github.com/jintzhu/gorm)
+ A `user` model with a basic setup of [guerilla/sessions]() for authentication. There is *no* pages or handlers for user signups or login/logout. It is left for you to implement.

## Where to start?
+ Define the route in `routes.go`
+ Add corresponding handler in `controllers/controllers.go`
+ Add the model structs in `models/models.go`
