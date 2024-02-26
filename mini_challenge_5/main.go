package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	EmailAddress    string
	Password        string
	Address         string
	Job             string
	ReasonToBeHappy string
}

var persons = []Person{
	{EmailAddress: "jonathan1@gmail.com", Password: "sanjaya123", Address: "Puri Lavender 23", Job: "Software Engineer", ReasonToBeHappy: "My family"},
	{
		EmailAddress:    "danielle98@hotmail.com",
		Password:        "brightside456",
		Address:         "Maple Oak Drive 47",
		Job:             "Data Analyst",
		ReasonToBeHappy: "Exploring new places",
	},
	{
		EmailAddress:    "sarah.connor@techmail.com",
		Password:        "terminator2024",
		Address:         "Cedarwood Lane 12",
		Job:             "Cybersecurity Specialist",
		ReasonToBeHappy: "Helping others stay safe online",
	},
}

var PORT = ":9090"

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/login", loginData)
	http.HandleFunc("/404", errorPage)

	fmt.Println("Application is listening in port", PORT)
	http.ListenAndServe(PORT, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// using html/template package
	if r.Method == "GET" {
		filepath := "./mini_challenge_5/login.html"
		tpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, persons)
		return
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	// using html/template package
	if r.Method == "GET" {
		filepath := "./mini_challenge_5/error.html"
		tpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, nil)
		return
	}
}

func loginData(w http.ResponseWriter, r *http.Request) {
	var loginPerson Person

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		for _, person := range persons {
			if email == person.EmailAddress && password == person.Password {
				loginPerson = person
			}
		}

		if loginPerson != (Person{}) {
			filepath := "./mini_challenge_5/profile.html"
			tpl, err := template.ParseFiles(filepath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tpl.Execute(w, loginPerson)
			return
		} else {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
		}
	}
}
