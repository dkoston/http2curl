package http2curl

import (
	"bytes"
	"fmt"
	"github.com/go-test/deep"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleGetCurlCommand() {
	form := url.Values{}
	form.Add("age", "10")
	form.Add("name", "Hudson")
	body := form.Encode()

	req, _ := http.NewRequest(http.MethodPost, "http://foo.com/cats", ioutil.NopCloser(bytes.NewBufferString(body)))
	req.Header.Set("API_KEY", "123")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'POST' -d 'age=10&name=Hudson' -H 'Api_key: 123' 'http://foo.com/cats'
}

func ExampleGetCurlCommand_json() {
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(`{"hello":"world","answer":42}`))
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_noBody() {
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", nil)
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'PUT' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_emptyStringBody() {
	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(""))
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'PUT' -d '' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_newlineInBody() {
	req, _ := http.NewRequest("POST", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString("hello\nworld"))
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'POST' -d 'hello
	// world' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_specialCharsInBody() {
	req, _ := http.NewRequest("POST", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu", bytes.NewBufferString(`Hello $123 o'neill -"-`))
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCurlCommand(req)
	fmt.Println(command)

	// Output:
	// curl -X 'POST' -d 'Hello $123 o'\''neill -"-' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func TestGetCurlCommand(t *testing.T) {
	Convey("Testing http2curl", t, func() {
		uri := "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu"
		payload := new(bytes.Buffer)
		payload.Write([]byte(`{"hello":"world","answer":42}`))
		req, err := http.NewRequest("PUT", uri, payload)
		So(err, ShouldBeNil)
		req.Header.Set("X-Auth-Token", "private-token")
		req.Header.Set("Content-Type", "application/json")

		command, err := GetCurlCommand(req)
		So(err, ShouldBeNil)
		expected := `curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' -H 'X-Auth-Token: private-token' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'`
		So(command.String(), ShouldEqual, expected)
	})
}

func TestCurlCommand_CommandSlice(t *testing.T) {
	headers := make(map[string]string,0)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:67.0) Gecko/20100101 Firefox/67.0"

	data := []byte("type=new&param=1&param2=banana")

	host := "https://somedomain.tldthatdoesnotexist"

	method := "POST"

	req, err := http.NewRequest(method, host, bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	expectedCommand := "curl"
	expectedArgs := []string{"-X", "'POST'", "-d","'type=new&param=1&param2=banana'","-H","'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:67.0) Gecko/20100101 Firefox/67.0'", "'https://somedomain.tldthatdoesnotexist'"}

	command, err := GetCurlCommand(req)
	if err != nil {
		t.Error(err)
	}

	cmd := command.Command()

	if cmd != expectedCommand {
		t.Errorf("Got unexpected command name. Expected: %s. Got: %s", expectedCommand, cmd)
	}

	args := command.Args()


	if diff := deep.Equal(args, expectedArgs); diff != nil {
		t.Errorf("Got unexpected command args. Diff: %v", diff)
	}
}