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
	//router.Handle("/", http.FileServer(http.Dir("./Public")))
	router.Handle("/files/{fn}", http.StripPrefix("/files", http.FileServer(http.Dir("./files"))))

	router.HandleFunc("/listen/{fn}", listen).Methods("GET")
	//router.HandleFunc("/listfiles", listfiles).Methods("GET")
	log.Println("Receiver Server")
	log.Println("Listening on Port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))
	// if err := http.ListenAndServe(":3001", nil); err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

}
func listen(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: " + r.Host)
	params := mux.Vars(r)
	filename := params["fn"]
	log.Println("Listening for file ... " + filename)
	gocastlisten.Receive(filename)
}

// func listfiles(w http.ResponseWriter, r *http.Request) {
// }
