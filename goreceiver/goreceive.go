package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	udpcastcmd := "udp-receiver"
	udpcastargs := []string{"--pipe", "\"gzip -d\"", "--file", "warpeace.txt"}

	cmd := exec.Command(udpcastcmd, udpcastargs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf(" udpcast output | %s\n", scanner.Text())
		}
	}()

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

}
