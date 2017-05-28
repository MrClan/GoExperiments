package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Nic", LastName: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "2", FirstName: "Maria", LastName: "Raboy"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}

// go practice : tutorial at https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/

// _______________________________________________________________________
// _______________________________________________________________________
// _______________________________________________________________________
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }

// _______________________________________________________________________
// _______________________________________________________________________
// _______________________________________________________________________
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	http.HandleFunc("/", hello)
// 	fmt.Println("listening...")
// 	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func hello(res http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintln(res, "hello, world from GO API")
// }
