package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Jordi Polla",
		Age: 29,
		Meta: UserMeta{
			Visits: 12,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}