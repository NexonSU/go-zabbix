package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NexonSU/go-zabbix/types"
)

// ErrNotFound describes an empty result set for an API call.
var (
	ErrNotFound      = &NotFoundError{"No results were found matching the given search parameters"}
	zabbixVersion600 *types.ZBXVersion
	zabbixVersion640 *types.ZBXVersion
)

func init() {
	zabbixVersion600, _ = types.NewZBXVersion("6.0.0")
	zabbixVersion640, _ = types.NewZBXVersion("6.4.0")
}

// A Session is an authenticated Zabbix JSON-RPC API client. It must be
// initialized and connected with NewSession.
type Session struct {
	// URL of the Zabbix JSON-RPC API (ending in `/api_jsonrpc.php`).
	URL string `json:"url"`

	// Token is the cached authentication token returned by `user.login` and
	// used to authenticate all API calls in this Session.
	Token string `json:"token"`

	// ApiVersion is the software version string of the connected Zabbix API.
	APIVersion *types.ZBXVersion `json:"apiVersion"`

	client *http.Client
}

// NewSession returns a new Session given an API connection URL and an API
// username and password.
//
// An error is returned if there was an HTTP protocol error, the API credentials
// are incorrect or if the API version is indeterminable.
//
// The authentication token returned by the Zabbix API server is cached to
// authenticate all subsequent requests in this Session.
func NewSession(url string, username string, password string) (session *Session, err error) {
	// create session
	session = &Session{URL: url}
	err = session.login(username, password)
	return
}

func (c *Session) login(username, password string) error {
	// get Zabbix API version
	ver, err := c.GetVersion()
	if err != nil {
		return fmt.Errorf("Failed to retrieve Zabbix API version: %v", err)
	}

	// login to API
	params := map[string]string{
		"password": password,
	}

	// user param was renamed in 6.0 and removed in 6.4
	if ver.Compare(zabbixVersion600) < 0 {
		params["user"] = username
	} else {
		params["username"] = username
	}

	res, err := c.Do(NewRequest("user.login", params), true)
	if err != nil {
		return fmt.Errorf("Error logging in to Zabbix API: %v", err)
	}

	err = res.Bind(&c.Token)
	if err != nil {
		return fmt.Errorf("Error failed to decode Zabbix login response: %v", err)
	}

	return nil
}

// GetVersion returns the software version string of the connected Zabbix API.
func (c *Session) GetVersion() (*types.ZBXVersion, error) {
	if c.APIVersion == nil {
		// get Zabbix API version
		res, err := c.Do(NewRequest("apiinfo.version", nil), true)
		if err != nil {
			return nil, err
		}

		err = res.Bind(&c.APIVersion)
		if err != nil {
			return nil, err
		}
	}

	return c.APIVersion, nil
}

// AuthToken returns the authentication token used by this session to
// authentication all API calls.
func (c *Session) AuthToken() string {
	return c.Token
}

// Do sends a JSON-RPC request and returns an API Response, using connection
// configuration defined in the parent Session.
//
// An error is returned if there was an HTTP protocol error, a non-200 response
// is received, or if an error code is set is the JSON response body.
//
// When err is nil, resp always contains a non-nil resp.Body.
//
// Generally Get or a wrapper function will be used instead of Do.
func (c *Session) Do(req *Request, noAuthRequired bool) (resp *Response, err error) {
	if noAuthRequired == false {
		// get Zabbix API version
		ver, err := c.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve Zabbix API version: %v", err)
		}

		// Zabbix 6.4 uses `Authorization` header, therefore "auth" parameter
		// has been deprecated and was removed in 7.2
		// See: https://www.zabbix.com/documentation/7.2/en/manual/api/changes
		if ver.Compare(zabbixVersion640) < 0 {
			req.AuthToken = c.Token
		}
	}

	// encode request as json
	b, err := json.Marshal(req)
	if err != nil {
		return
	}

	dprintf("Call     [%s:%d]: %s\n", req.Method, req.RequestID, b)

	// create HTTP request
	r, err := http.NewRequest("POST", c.URL, bytes.NewReader(b))
	if err != nil {
		return
	}
	r.ContentLength = int64(len(b))
	r.Header.Add("Content-Type", "application/json-rpc")
	if noAuthRequired == false {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	// send request
	client := c.client
	if client == nil {
		client = http.DefaultClient
	}
	res, err := client.Do(r)
	if err != nil {
		return
	}

	defer res.Body.Close()

	// read response body
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	dprintf("Response [%s:%d]: %s\n", req.Method, req.RequestID, b)

	// map HTTP response to Response struct
	resp = &Response{
		StatusCode: res.StatusCode,
	}

	// unmarshal response body
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON response body: %v", err)
	}

	// check for API errors
	if err = resp.Err(); err != nil {
		return
	}

	return
}

// Get calls the given Zabbix API method with the given query parameters and
// unmarshals the JSON response body into the given interface.
//
// An error is return if a transport, marshalling or API error happened.
func (c *Session) Get(method string, params interface{}, v interface{}) error {
	req := NewRequest(method, params)
	resp, err := c.Do(req, false)
	if err != nil {
		return err
	}

	err = resp.Bind(v)
	if err != nil {
		return err
	}

	return nil
}
