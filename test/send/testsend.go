package main

import (
	"log"
	"net/http"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

const receiverListenIP = "http://172.24.0.194:3001/listen/"
const filename = "warpeace.txt"

func main() {
	log.Println("Sending file...")
	log.Println("Making call request to listen @:" + receiverListenIP)
	resp, err := http.Get(receiverListenIP + filename)
	if err != nil {
		log.Println(err.Error())
		return
	}
	resp.Body.Close()
	gocastsend.Send("../../gosender/samplefiles/" + filename)
}
