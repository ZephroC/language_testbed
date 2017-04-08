package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/ZephroC/test_go_app/status"
	"github.com/ZephroC/test_go_app/config"
	"github.com/ZephroC/test_go_app/books"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("config.yml")
	check(err)
	serverConfig := config.Config{}
	err2 := yaml.Unmarshal(data, &serverConfig)
	check(err2)
	log.Println("Working with config: \n", string(data))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", status.StatusHandler(serverConfig)).Methods("GET")
	router.HandleFunc("/books",books.ListBooks(serverConfig)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

