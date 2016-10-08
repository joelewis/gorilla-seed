package models

import (
    "time"
    "github.com/satori/go.uuid"
)


type Model struct {
    ID        string `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

func (model *Model) BeforeCreate() {
    model.ID = uuid.NewV4().String()
}

type User struct {
    Model
    Username string
    Password string
    Phonenumber string
    Email string    
}

func (user *User) Get(id string) *User {
    Db.First(user, id)
    return user
}

func (user *User) Save() {
    Db.Create(user)
}

func NewUser(username, password, phonenumber, email string) *User {
    user := User{
        Username: username,
        Password: password,
        Phonenumber: phonenumber,
        Email: email,
    }
    return &user
}
