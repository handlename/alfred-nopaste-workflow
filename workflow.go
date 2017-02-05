package wf

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// Client for nopaste
type Client struct {
	Endpoint string
	User     string
	Password string
}

// Post text to nopaste.
func (c *Client) Post(content io.Reader) (resultURL *url.URL, err error) {
	if c.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}

	text, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read content")
	}

	values := make(url.Values)
	values.Set("text", string(text))

	req, err := http.NewRequest("POST", c.Endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// for basic auth
	if c.User != "" {
		req.SetBasicAuth(c.User, c.Password)
	}

	client := new(http.Client)
	client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		resultURL = r.URL
		return errors.New("")
	}

	res, err := client.Do(req)
	if err != nil && res.StatusCode != http.StatusFound {
		return nil, errors.Wrap(err, "failed to post")
	}

	defer res.Body.Close()

	if resultURL == nil {
		return nil, fmt.Errorf("failed to post")
	}

	return resultURL, nil
}
