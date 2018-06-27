package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alanpadillachua/GoCast/goreceiver/gocastlisten"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./Public"))))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))

	router.HandleFunc("/listen/{fn}", listen).Methods("GET")
	router.HandleFunc("/files", files).Methods("GET")

	log.Println("Receiver Server")
	log.Println("Listening on Port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))

}

func listen(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: " + r.Host)
	params := mux.Vars(r)
	filename := params["fn"]
	log.Println("Listening for file ... " + filename)
	gocastlisten.Receive(filename)
}
