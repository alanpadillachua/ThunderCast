package main

import (
	"log"
	"net/http"
)

func main() {
	const prefix = "/GoUPDCast/v1"
	http.HandleFunc(prefix+"/upload/", uploadHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method)
	if r.Method != http.MethodPost {
		w.Header().Add("accept", http.MethodPost)
		http.Error(w, "method must be POST", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseMultipartForm(64); err != nil {
		log.Println(err.Error())
	}

	file, handler, fileErr := r.FormFile("file")

	if fileErr != nil {
		http.Error(w, fileErr.Error(), http.StatusInternalServerError)
		log.Println(fileErr)
		return
	}
	log.Println(handler.Filename)
	log.Println(handler.Header)

	file.Close()
}
