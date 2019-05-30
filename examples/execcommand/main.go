package main

import (
	"bytes"
	"fmt"
	"github.com/dkoston/http2curl"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	data := bytes.NewBufferString(``)
	req, _ := http.NewRequest("GET", "https://davekoston.com/http/test.php", data)

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		log.Fatal(err)
	}

	commandName := command.Command()
	args := command.Args()

	cmd := exec.Command(commandName, args...)
	var readBuffer, writeBuffer bytes.Buffer
	cmd.Stdout = &readBuffer
	cmd.Stderr = &writeBuffer

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed running curl command: %v. %v", err, writeBuffer.String())
	}
	fmt.Printf("Successfully ran: %s.\n %v", command.String(), readBuffer.String())
}
