package grabber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type mockTransport struct{}

func newMockTransport() http.RoundTripper {
	return &mockTransport{}
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "text/html")

	responseBody := fmt.Sprintf(
		`HTTP/1.1 200 OK
Server: FakeServer
Keep-Alive: timeout=2, max=200
Content-Type: text/html

MOCK for %s %s/%s
`, req.Method, req.Host, req.RequestURI)
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}
