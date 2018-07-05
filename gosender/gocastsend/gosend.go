package gocastsend

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Send file starts udp-sender and sends file
func Send(file string) {
	//udpsendcmd := exec.Command("udp-sender", "--async", "--file", "samplefiles/warpeace.txt")
	//udpsendcmd := exec.Command("udp-sender", "--async", "--fec", "8x8", "--pipe", "\"gzip -fN\"", "--file", "samplefiles/warpeace.txt")

	udpcastcmd := "udp-sender"
	//"--pipe", "\"gzip -f\"", "--min-receivers", "1",
	udpcastargs := []string{"--nokbd", "--async", "--max-bitrate", "40m", "--fec", "8x8", "--file", file}
	cmd := exec.Command(udpcastcmd, udpcastargs...)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s\n", out)
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

func compress(file string) {
	f, _ := os.Open("./files/" + file)

	// Create a Reader and use ReadAll to get all the bytes from the file.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Replace txt extension with gz extension.

	// Open file for writing.
	f, _ = os.Create("./files/" + file)

	// Write compressed data.
	w := gzip.NewWriter(f)
	w.Write(content)
	w.Close()

	// Done.
	log.Println("File: " + file + " Compressed in ./files")
}
