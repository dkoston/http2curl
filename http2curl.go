package http2curl

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// CurlCommand contains exec.Command compatible command and args
type CurlCommand struct {
	command string
	args []string
}

// append appends a string to arguments of the CurlCommand
func (c *CurlCommand) append(newSlice ...string) {
	c.args = append(c.args, newSlice...)
}

// Args returns the arguments for the curl command as a slice
func (c *CurlCommand) Args() []string {
	return c.args
}

// Args returns the command (program) name as a string
func (c *CurlCommand) Command() string {
	return c.command
}

// String returns a ready to copy/paste command
func (c *CurlCommand) String() string {
	return c.command + " " + strings.Join(c.args, " ")
}

// nopCloser is used to create a new io.ReadCloser for req.Body
type nopCloser struct {
	io.Reader
}

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

func (nopCloser) Close() error { return nil }

// GetCurlCommand returns a CurlCommand corresponding to an http.Request
func GetCurlCommand(req *http.Request) (*CurlCommand, error) {
	c := CurlCommand{}
	c.command = "curl"


	c.append("-X", bashEscape(req.Method))

	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = nopCloser{bytes.NewBuffer(body)}
		bodyEscaped := bashEscape(string(body))
		c.append("-d", bodyEscaped)
	}

	var keys []string

	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		c.append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}

	c.append(bashEscape(req.URL.String()))

	return &c, nil
}
