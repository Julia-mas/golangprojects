package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadFile("/home/lin/golang/golangprojects/users.json")
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("Json decode error!")
		return
	}

	fmt.Println(data)

	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8005", router))
}
