package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// funcion para obtener tadas las tareas
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // envio de datos en formato json
	json.NewEncoder(w).Encode(users)
}

//modulo

type user struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	Status string `json:"Status"`
}
type alluser []user

var users = alluser{
	{
		Id:     1,
		Name:   "TEST one",
		Email:  "testgmailcom",
		Status: "on line",
	},
}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bienbenidos a mi api Go! ")
}

// crear usuarios
// incremento de usuarios
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Usuario insertado")
	}
	json.Unmarshal(reqBody, &newUser)
	newUser.Id = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// obteniendo usuario por id
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//obtener usuario por id
func getOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UsersId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, user := range users {
		if user.Id == UsersId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	}
}

//actualizar usuario por id
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId, err := strconv.Atoi(vars["id"])
	var updatedUser user

	if err != nil {
		fmt.Fprintf(w, "Id Invalido")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Verifique el id del usuario!")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, t := range users {
		if t.Id == UserId {
			users = append(users[:i], users[i+1:]...)

			updatedUser.Id = t.Id
			users = append(users, updatedUser)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "El usuario por ID %v se ha actualizado correctamente", UserId)
		}
	}

}

// eliminar usuario por id
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Usuario invalido por ID")
		return
	}

	for i, t := range users {
		if t.Id == UserId {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", UserId)
		}
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRouter)
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")

	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
