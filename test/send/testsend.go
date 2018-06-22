package main

import (
	"log"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

func main() {
	log.Println("Sending file...")
	gocastsend.Send("../../gosender/samplefiles/" + "warpeace.txt")
}
