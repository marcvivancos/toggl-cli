package toggl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURI    = "https://api.track.toggl.com/api/v9"
	retryCount = 3
)

var (
	errRetryCountExceeded = errors.New("API request retry count exceeded")
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(client *http.Client, token string) *Client {
	return &Client{
		client: client,
		token:  token,
	}
}

func NewDefaultClient(token string) *Client {
	return &Client{
		client: http.DefaultClient,
		token:  token,
	}
}

func (cl *Client) do(method string, endpoint string, queryParams map[string]string, bodyParams interface{}) (res *http.Response, err error) {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	// Add query parameters to the URL
	if queryParams != nil {
		q := uri.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		uri.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return
	}

	basic := base64.StdEncoding.EncodeToString([]byte(cl.token + ":api_token"))
	req.Header.Add("Authorization", "Basic "+basic)

	var buff []byte
	if bodyParams != nil {
		buff, err = json.Marshal(bodyParams)
		if err != nil {
			return
		}
		req.Body = io.NopCloser(bytes.NewReader(buff))
		req.Header.Set("Content-Type", "application/json")
	}

	count := 0
	for count < retryCount {
		res, err := cl.client.Do(req)
		if err == nil {
			return res, err
		}
		count++
	}

	return nil, errRetryCountExceeded
}
