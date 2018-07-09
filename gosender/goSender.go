package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

const receiverListenIP = "http://172.24.0.194:3001/listen/v1?"

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
			jsonResponse(w, http.StatusCreated, "Error: File Upload")
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			jsonResponse(w, http.StatusCreated, "Error: File Upload")
			return
		}
	}
	log.Println("Reading file: " + filename)
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
	filehash, err := hashFileMd5(filename)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		log.Println(err.Error())
	}

	log.Println("File hash: " + filehash)

	go startListening(filename, filehash)
	time.Sleep(2 * time.Second)
	// compress file
	gocastsend.Send("./files/" + filename) // send file through diod

	r.Body.Close()
	time.Sleep(time.Second)
	log.Println("Deleting File: " + filename)
	deleteFile(filename) // delete file locally

}

func startListening(file string, hash string) {
	log.Println("Making call request to listen @:" + receiverListenIP)
	request := receiverListenIP + "filename=" + file + "&hash=" + hash
	http.Get(request)
}
func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
func hashFileMd5(filename string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open("./files/" + filename)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func deleteFile(filename string) {
	err := os.Remove("./files/" + filename)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
