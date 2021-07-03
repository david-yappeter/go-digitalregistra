package digitalregistra

import (
	"fmt"
	"net/http"
	"net/url"
)

func (a *API) endpointCheckDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}
