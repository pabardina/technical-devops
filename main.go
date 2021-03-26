package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

var Todos []Todo

type Todo struct {
    Id      string `json:"Id"`
    Description   string `json:"Description"`
}

func getAllTodos(w http.ResponseWriter, r *http.Request){
    log.Println("Get All todos")
    json.NewEncoder(w).Encode(Todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
    log.Println("Add a todo")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var todo Todo
    json.Unmarshal(reqBody, &todo)
    Todos = append(Todos, todo)
    json.NewEncoder(w).Encode(todo)
    log.Println("Todo added")
}

func getTodo(w http.ResponseWriter, r *http.Request){
    log.Println("Get a todo")
    vars := mux.Vars(r)
    key := vars["id"]
    for _, todo := range Todos {
        if todo.Id == key {
            json.NewEncoder(w).Encode(todo)
        }
    }
}

func deleteTodo(w http.ResponseWriter, r *http.Request){
    log.Println("Delete a todo")
    vars := mux.Vars(r)
    id := vars["id"]

    for index, todo := range Todos {
        if todo.Id == id {
            Todos = append(Todos[:index], Todos[index+1:]...)
        }
    }
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
    log.Println("Update a todo")
    vars := mux.Vars(r)
    id := vars["id"]

    reqBody, _ := ioutil.ReadAll(r.Body)
    var updateTodo Todo
    json.Unmarshal(reqBody, &updateTodo)

    for index, todo := range Todos {
        if todo.Id == id {
            todo := &Todos[index]
            todo.Description = updateTodo.Description
        }
    }

}

func main() {
	fmt.Println("Starting this awesome api..")

    Todos = []Todo{
        Todo{Id:"1", Description: "Meditate"},
        Todo{Id:"2", Description: "Yoga"},
    }

    router := mux.NewRouter()
    router.HandleFunc("/", getAllTodos).Methods("GET")
    router.HandleFunc("/", addTodo).Methods("POST")
    router.HandleFunc("/{id}", deleteTodo).Methods("DELETE")
    router.HandleFunc("/{id}", updateTodo).Methods("PUT")
    router.HandleFunc("/{id}", getTodo).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}