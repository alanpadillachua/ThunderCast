package main

import (
	"log"
	"net/http"

	"github.com/alanpadillachua/GoCast/goreceiver/gocastlisten"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Public")))
	http.HandleFunc("/listen", listen)
	log.Println("Receiver Server")
	log.Println("Listening on Port 3001")

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func listen(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: " + r.Host)
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	log.Println("Listening for file ...")
	gocastlisten.Receive("samplefile")
}
