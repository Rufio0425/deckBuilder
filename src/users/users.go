package users

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type user struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := uuid.Parse(params["user_id"])
	if err != nil {
		fmt.Printf("Error parsing uuid: %s", err.Error())
	}

	user := user{
		ID:       userId,
		UserName: "test_user",
		Email:    "test@gmail.com",
	}

	if err != json.NewEncoder(w).Encode(user) {
		fmt.Printf("Error encoding JSON: %s", err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		formattedErr := fmt.Sprintf("Error Decoding user: %v", err.Error())
		http.Error(w, formattedErr, http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(user) {
		fmt.Printf("Error encoding JSON: %s", err)
	}
}
