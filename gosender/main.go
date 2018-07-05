package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

const receiverListenIP = "http://172.24.0.194:3001/listen/"

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
	/*if err := r.ParseMultipartForm(256 << 20); err != nil {
		log.Println(err.Error())
	}*/
	//reader, err := r.MultipartReader()

	// file, handle, err := r.FormFile("file")
	// log.Println("Reading file: " + handle.Filename)
	// if err != nil {
	// 	fmt.Fprintf(w, "%v", err)
	// 	return
	// }
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		log.Println(err.Error())
		return
	}
	//copy each part to destination.
	filename := ""
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}
		dst, err := os.Create("./files/" + part.FileName())
		filename = part.FileName()
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Println("Reading file: " + filename)
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
	startListening(filename)
	//defer file.Close()
	//saveFile(w, file, handle)

	r.Body.Close()
}

func startListening(file string) {
	resp, err := http.Get(receiverListenIP + file)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if resp.StatusCode == http.StatusOK {
		resp.Body.Close()
		log.Println("Transfering file: " + file)
		gocastsend.Send("./files/" + file) // send file through diod
	}

}
func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
