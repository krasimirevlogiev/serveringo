
package main

import (
    "fmt"
    "net/http"
    "html/template"
    "github.com/pesho65/serveringo/user"  // Import the user package
)

var templates = template.Must(template.ParseFiles("static/register.html", "static/login.html"))

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        templates.ExecuteTemplate(w, "register.html", nil)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    if user.RegisterUser(username, password) {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    } else {
        http.Error(w, "User already exists", http.StatusConflict)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        templates.ExecuteTemplate(w, "login.html", nil)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    if user.AuthenticateUser(username, password) {
        fmt.Fprintf(w, "Welcome, %s!", username)
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}

func main() {
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)

    fmt.Println("Starting server at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}

