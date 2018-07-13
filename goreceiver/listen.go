package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/alanpadillachua/GoCast/goreceiver/gocastlisten"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./Public/"))))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))

	router.HandleFunc("/listen/{vars}", listen).Methods("GET").Queries("filename", "{filename}", "hash", "{hash}")
	log.Println("Receiver Server")
	log.Println("Listening on Port 3001")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":3001", loggedRouter))

}

func listen(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from: " + r.Host)
	params := mux.Vars(r)
	filename := params["filename"]
	hashsum := params["hash"]
	log.Println("Listening for file:" + filename)
	log.Println("Hash of file: " + hashsum)
	gocastlisten.Receive("./files/" + filename)

	// Decompress file

	hashbuilt, err := hashFileMd5(filename)
	if err != nil {
		log.Println(err.Error())
	}
	if hashbuilt == hashsum {
		log.Println("File hash verified. File transfered successfully")
		log.Println("Hash Expected: " + hashsum)
		log.Println("Hash Recieved: " + hashbuilt)

	} else {
		log.Println("Error File hash integreity lost. Please retry transfer ")
		log.Println("Hash Expected: " + hashsum)
		log.Println("Hash Recieved: " + hashbuilt)
	}
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
