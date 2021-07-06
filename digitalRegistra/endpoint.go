package digitalregistra

import (
	"fmt"
	"net/http"
	"net/url"
)

func (a *API) endpointCheckDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}

func (a *API) endpointDomainRegistration(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}

func (a *API) endpointCreateContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/create"), data
}

func (a *API) endpointInfoContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/info"), data
}
