package main

import (
	"log"
	"net/http"

	"github.com/alanpadillachua/GoCast/goreceiver/gocastlisten"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Public")))
	http.HandleFunc("/files", listFiles)
	log.Println("Listening on Port 3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	gocastlisten.Receive("samplefile")
	log.Println("Listening for file")
}
func listFiles(w http.ResponseWriter, r *http.Request) {
	gocastlisten.Receive("samplefile")
	log.Println("Listening for file")

	log.Println("Connection from: " + r.Host)
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
