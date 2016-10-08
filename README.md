# A quick seed for the golang-gorilla stack

## Getting Started

+ Clone the repo `git clone https://github.com/joelewis/gorilla-seed`
+ Update config variables in `main.go` (DB_URL, SERVER_PORT, etc.)
+ Rename the directory's name to suit yours. 
  `mv /path/to/gorilla-seed /path/to/my-app-name`
+ update import paths in go files.
  Eg. `import "path/to/gorilla-seed/models"` 
  should be `import "path/to/my-app-name/models"`
+ Compile: `go build`
+ Initialize tables: `./my-app-name inittables`
+ Run server : `./my-app-name`

## What's Inside?
Nothing much.
+ A developer friendly ORM - [gorm](https://github.com/jintzhu/gorm)
+ A User model with basic setup of [guerilla/sessions]() for authentication. There's *no* register, login pages or handlers burnt into the project. It's upto you to implement those.


## Where to Start?
Simple.
+ Define route in `routes.go`
+ Add corresponding handler in `controllers/controllers.go`
+ Add your model structs in `models/models.go`
