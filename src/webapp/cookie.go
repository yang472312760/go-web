package main

import (
	"net/http"
	"fmt"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:60,
	}

	http.SetCookie(w, &cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("Cookie")
	fmt.Print("得到Cookie:", cookie)
}

func main() {

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.ListenAndServe(":8080", nil)
}
