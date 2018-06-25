package main

import (
	"log"

	"github.com/alanpadillachua/GoCast/goreceiver/gocastlisten"
)

func main() {
	log.Println("Listening for file ...")
	gocastlisten.Receive("samplefile")

}
