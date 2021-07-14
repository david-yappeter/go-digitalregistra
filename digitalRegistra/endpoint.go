package digitalregistra

import (
	"fmt"
	"net/http"
	"net/url"
)

// DOMAIN
// ====================================================================

func (a *API) endpointCheckDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}

func (a *API) endpointDomainRegistration(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}

// Contact
// ====================================================================

func (a *API) endpointCreateContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/create"), data
}

func (a *API) endpointUpdateContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/update"), data
}

func (a *API) endpointInfoContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/info"), data
}

func (a *API) endpointDeleteContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/delete"), data
}

// User
// ====================================================================

func (a *API) endpointCreateUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/create"), data
}

func (a *API) endpointInfoUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/info"), data
}

func (a *API) endpointUpdateUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/update"), data
}
