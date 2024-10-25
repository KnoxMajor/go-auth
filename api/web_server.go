package api

import (
	"encoding/json"
	"github.com/knoxmajor/go-auth/internal/user"
	"log"
	"net/http"
)

func fourOhFourHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("I'm sorry but this page doesn't exist"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid route", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	err = user.Signup(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func RequestController(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		next.ServeHTTP(w, r)
	}
}

func InitRoutes() {
	http.HandleFunc("/", RequestController(fourOhFourHandler))
	http.HandleFunc("/signup", RequestController(signupHandler))
}

func StartServer() {
	log.Print("Starting Server at :8080", "\n")
	InitRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Server start failed.", "\n")
		return
	}
}
