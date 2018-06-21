package gosender

import (
	"bufio"
	"compress/gzip"
	"fmt"
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
}

func compress(file string) {
	f, _ := os.Open("../files/" + file)

	// Create a Reader and use ReadAll to get all the bytes from the file.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Replace txt extension with gz extension.

	// Open file for writing.
	f, _ = os.Create("../files/" + file)

	// Write compressed data.
	w := gzip.NewWriter(f)
	w.Write(content)
	w.Close()

	// Done.
	log.Println("File: " + file + " Compressed in ./files")
}
