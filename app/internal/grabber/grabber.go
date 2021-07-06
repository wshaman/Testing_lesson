package grabber

import (
	"io/ioutil"
	"net/http"
)

type Grabber struct {
	T *http.Client
}

func New(h *http.Client) *Grabber {
	return &Grabber{T: h}
}

func (g *Grabber) Grab(url string) (string, error) {

	resp, err := g.T.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
