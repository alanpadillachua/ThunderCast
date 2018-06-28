package main

import (
	"log"

	"github.com/alanpadillachua/GoCast/gosender/gocastsend"
)

func main() {
	log.Println("Sending file...")
	gocastsend.Send("../../gosender/samplefiles/" + "ubuntu-18.04-desktop-amd64.iso")
}
