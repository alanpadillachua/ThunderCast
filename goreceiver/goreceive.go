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
	udpsendcmd := exec.Command("udp-receiver", "--pipe \"gzip -d\"", "--file warpeace.txt")

	sendout, err := udpsendcmd.Output()
	handleErr(err)
	fmt.Println(string(sendout))
}
