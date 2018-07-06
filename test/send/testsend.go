package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

const receiverListenIP = "http://172.24.0.194:3001/listen/"
const filename = "warpeace.txt"

func main() {
	// log.Println("Sending file...")
	// log.Println("Making call request to listen @:" + receiverListenIP)
	// go http.Get(receiverListenIP + filename)
	// go gocastsend.Send("../../gosender/samplefiles/" + filename)
	// log.Panicln("Files sent")
	//readyToSend := make(chan bool, 1)

	http.Get(receiverListenIP + filename)

	time.Sleep(1 * time.Second)
	log.Println("Transfering file: " + filename)
	gocastsend.Send("./files/" + filename) // send file through diod

}

func startListening(file string) {
	log.Println("Making call request to listen @:" + receiverListenIP)
	resp, err := http.Get(receiverListenIP + file)
	if err != nil {
		log.Println(err.Error())
		return
	}

	resp.Body.Close()
	log.Println("Transfering file: " + file)
	gocastsend.Send("./files/" + file) // send file through diod

}
