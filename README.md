# http2curl
:triangular_ruler: Convert Golang's http.Request to CURL commands for either:

- Printing in logs or for reference
- Executing via os/exec's exec.Command()

[![Build Status](https://travis-ci.org/dkoston/http2curl.svg?branch=master)](https://travis-ci.org/dkoston/http2curl)
[![GoDoc](https://godoc.org/github.com/dkoston/http2curl?status.svg)](https://godoc.org/github.com/dkoston/http2curl)
[![Coverage Status](https://coveralls.io/repos/dkoston/http2curl/badge.svg)](https://coveralls.io/github/dkoston/http2curl)


To do the reverse, check out [mholt/curl-to-go](https://github.com/mholt/curl-to-go).

## Examples

### Printing the curl command

```go
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
```

### exec.Command() execution of curl command

```go
package main
import (
	"bytes"
	"fmt"
    "github.com/dkoston/http2curl"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
    data := bytes.NewBufferString(`{"hello":"world","answer":42}`)
    req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", data)
    req.Header.Set("Content-Type", "application/json")
    
    command, err := http2curl.GetCurlCommand(req)
    if err != nil {
        log.Fatal(err)
    }
    
    commandName := command.Command()
    args := command.Args()

    cmd := exec.Command(commandName, args...)
    readBuffer := bytes.Buffer{}
    cmd.Stdout = &readBuffer
    cmd.Stderr = os.Stderr

    err = cmd.Run()
    if err != nil {
    	log.Fatalf("Failed running curl command: %v. %v", err, os.Stderr)
    }
    fmt.Printf("Successfully ran: %s. %v", command, os.Stdout)
}

```

## Install

```bash
$ go get github.com/dkoston/http2curl
```

## License

MIT

## Attribution

Forked from [https://github.com/moul/http2curl](https://github.com/moul/http2curl)

motivation: moul's library was focused on printing the command only. I also 
wanted to be able to pass the command to exec.Command() easily
