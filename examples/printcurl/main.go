package main

import (
	"bytes"
	"fmt"
	"github.com/dkoston/http2curl"
	"log"
	"net/http"
)

func main() {
	data := bytes.NewBufferString(`{"hello":"world","answer":42}`)
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", data)
	req.Header.Set("Content-Type", "application/json")

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(command)
	// Output: curl -X PUT -d "{\"hello\":\"world\",\"answer\":42}" -H "Content-Type: application/json" http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu
}