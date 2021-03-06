package digitalregistra

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type API struct {
	DefaultUrl string        `json:"default_url"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Client     *http.Client  `json:"client"`
	Request    *http.Request `json:"request"`
	Error      error         `json:"error"`
}

type Option struct {
	HttpClient *http.Client `json:"http_client"`
}

func NewAPI(defaultUrl string, username string, password string, client ...Option) *API {
	api := API{
		DefaultUrl: defaultUrl,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	if len(client) > 0 {
		api.Client = client[0].HttpClient
	}

    api.Username = username
    api.Password = password
	return &api
}

func (a *API) newRequest(method string, link string, data url.Values) *API {
	var body io.Reader = nil
	if data != nil {
		data.Add("username", a.Username)
		data.Add("password", a.Password)
	} else {
		data = url.Values{}
		data.Add("username", a.Username)
		data.Add("password", a.Password)
	}
	body = bytes.NewBufferString(data.Encode())
	request, err := http.NewRequest(method, link, body)
	if err != nil {
		a.Error = err
		return a
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	a.Request = request
	return a
}

func (a *API) do() ([]byte, error) {
	resp, err := a.Client.Do(a.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
