package api

import "myapp/digitalregistra"

func NewClient(defaultUrl string, username string, password string, client ...digitalregistra.Option) *digitalregistra.API {
	return digitalregistra.NewAPI(defaultUrl, username, password, client...)
}
