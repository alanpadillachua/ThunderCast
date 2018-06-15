package goudpcast

import (
	"fmt"
	"os"
	"os/exec"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func send(file string) {
	//udpsendcmd := exec.Command("udp-sender", "--async", "--file", "samplefiles/warpeace.txt")
	//udpsendcmd := exec.Command("udp-sender", "--async", "--fec", "8x8", "--pipe", "\"gzip -fN\"", "--file", "samplefiles/warpeace.txt")

	udpcastcmd := "udp-sender"
	//"--pipe", "\"gzip -f\"",
	udpcastargs := []string{"--async", "--fec", "8x8", "--file", file}
	cmd := exec.Command(udpcastcmd, udpcastargs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", out)
	/*
		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
			os.Exit(1)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
			os.Exit(1)
		}
	*/
}
