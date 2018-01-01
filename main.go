package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Person struct {
  ID        string   `json: "id, omitempty"`
  Firstname string   `json: "firstname, omitempty"`
  Lastname  string   `json:"lastname, omitempty"`
  Address   *Address `json: "address, omitempty"`
}

type Address struct {
  City  string `json:"city, omitempty"`
  State string `json:"state, omitempty"`
}

var people []Person

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("Get")
  router.HandleFunc("/people/{id}", GetPerson).Methods("Get")
  router.HandleFunc("/people/{id}", Createperson).Methods("Post")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("Delete")
  log.Fatal(http.ListenAndServe(":8080", router))

  people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
  people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
  people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(people)
  fmt.Println(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func Createperson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range people {
    if item.ID == params["id"] {
      people = append(people[:index], people[index+1:]...)
      break
    }
    json.NewEncoder(w).Encode(people)
  }
}
