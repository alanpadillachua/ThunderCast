package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Public")))
	http.HandleFunc("/upload", UploadFile)
	log.Println("Sender Server")
	log.Println("Listening on Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//
}

// UploadFile uploads a file to the server
func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: " + r.Host)
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseMultipartForm(256 << 20); err != nil {
		log.Println(err.Error())
	}
	file, handle, err := r.FormFile("file")
	log.Println("Reading file: " + handle.Filename)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()
	saveFile(w, file, handle)
	gocastsend.Send("./files/" + handle.Filename) // send file through diod
	r.Body.Close()                                // close
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	log.Println("Uploading File: " + handle.Filename)
	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
