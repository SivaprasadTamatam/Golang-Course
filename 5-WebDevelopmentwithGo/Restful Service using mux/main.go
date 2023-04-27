package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var users []User

func main() {
	router := mux.NewRouter()

	// Add sample data
	users = append(users, User{ID: "1", Name: "Sivaprasad", Email: "Sivaprasade@example.com", Password: "password"})
	users = append(users, User{ID: "2", Name: "Tamatam", Email: "Tamatam@example.com", Password: "password"})

	// Define routes
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, user := range users {
		if user.ID == params["id"] {
			users[i] = User{
				ID:       params["id"],
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
			}
			_ = json.NewDecoder(r.Body).Decode(&users[i])
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, user := range users {
		if user.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

/*
curl -X GET http://localhost:8000/users
curl -X GET http://localhost:8000/users/1
curl -d '{"id":"3","name":"Core","email":"core@example.com","password":"password"}' -H "Content-Type: application/json" -X POST http://localhost:8000/users
curl -d '{"id":"1","name":"Sivaprasad Tamatam","email":"Sivaprasad Tamatam@example.com","password":"newpassword"}' -H "Content-Type: application/json" -X PUT http://localhost:8000/users/1
curl -X DELETE http://localhost:8000/users/1
*/
