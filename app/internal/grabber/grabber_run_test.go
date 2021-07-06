package grabber

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getRealClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func getFakeClient() *http.Client {
	c := getRealClient()
	c.Transport = newMockTransport()
	return c
}

func TestRun(t *testing.T) {
	g := New(getRealClient())
	data, err := g.Grab("http://example.org")
	require.NoError(t, err)
	fmt.Println(data)
}

func TestFakeRun(t *testing.T) {
	g := New(getFakeClient())
	data, err := g.Grab("http://example.org")
	require.NoError(t, err)
	fmt.Println(data)
}
