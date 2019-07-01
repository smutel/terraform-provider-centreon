package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	pkgerrors "github.com/pkg/errors"
)

// Client struct is used as generic client
type Client struct {
	HTTPClient *http.Client
	RequestURL *url.URL
}

// New returns a new generic client created with specified parameters
func New(requestURL string, insecure bool) (*Client, error) {
	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "URL not correctly formatted")
	}

	http := &http.Client{
		Transport:     httpTransport(insecure),
		CheckRedirect: doNotFollowRedirects,
	}

	c := &Client{
		HTTPClient: http,
		RequestURL: u,
	}

	return c, nil
}

// InsecureSkipTLSVerify turns off TLS verification for the client connection
func (c *Client) InsecureSkipTLSVerify() {
	if c.HTTPClient == nil {
		return
	}

	c.HTTPClient.Transport = httpTransport(true)
}

func httpTransport(insecureSkipTLSVerify bool) *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConns:        10,
		IdleConnTimeout:     15 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipTLSVerify,
		},
	}
}

func doNotFollowRedirects(*http.Request, []*http.Request) error {
	return http.ErrUseLastResponse
}

// ClearBody removes double quotes from the body
func ClearBody(body []byte) []byte {
	if len(body) > 0 && body[0] == '"' {
		body = body[1:]
	}

	if len(body) > 0 && body[len(body)-1] == '"' {
		body = body[:len(body)-1]
	}

	return body
}

// ----------------------------------------------------------------------------

// RequestInput stores the parameters of the centreon API call
type RequestInput struct {
	Method string
	Path   string
	Query  *url.Values
	Header *http.Header
	Body   interface{}
}

// ExecuteRequest will execute the request sent in parameter
func (c *Client) ExecuteRequest(inputs RequestInput) (io.ReadCloser, error) {
	method := inputs.Method
	path := inputs.Path
	query := inputs.Query
	header := inputs.Header
	body := inputs.Body

	var requestBody io.Reader
	if body != nil {
		switch body.(type) {
		default:
			marshaled, err := json.MarshalIndent(body, "", "    ")
			if err != nil {
				return nil, pkgerrors.Wrap(err, "Unable to create JSON with this body")
			}
			requestBody = bytes.NewReader(marshaled)
		case *url.Values:
			requestBody = bytes.NewBufferString(body.(*url.Values).Encode())
		}
	}

	endpoint := c.RequestURL
	endpoint.Path = path
	if query != nil {
		endpoint.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, endpoint.String(), requestBody)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Unable to create new http request")
	}

	if header != nil {
		for k := range *header {
			req.Header.Set(k, header.Get(k))
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Unable to execute http request")
	}

	if resp.StatusCode >= http.StatusOK &&
		resp.StatusCode < http.StatusMultipleChoices {
		return resp.Body, nil
	}

	return nil, pkgerrors.New("API error: " + resp.Status)
}
