package goudpcast

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func receive(file string) {
	udpcastcmd := "udp-receiver"
	//"--pipe", "\"gzip -fd\"",
	udpcastargs := []string{"--file", file}

	cmd := exec.Command(udpcastcmd, udpcastargs...)
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, err := cmd.StdoutPipe()
	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err = cmd.Start()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, outErr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	fmt.Println(outStr)
	fmt.Println(outErr)

}
