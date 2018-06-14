package main

import (
	"fmt"
	"os/exec"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//udpsendcmd := exec.Command("udp-sender", "--async", "--file", "samplefiles/warpeace.txt")
	udpsendcmd := exec.Command("udp-sender", "--async", "--fec", "8x8", "--pipe", "\"gzip -fN\"", "--file", "samplefiles/warpeace.txt")

	sendout, err := udpsendcmd.Output()
	handleErr(err)
	fmt.Println(string(sendout))
}
