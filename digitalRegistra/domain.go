package digitalregistra

import (
	"fmt"
	"net/url"
)

func (a *API) DomainCheck(domain string) {
	if err := a.newRequest(a.endpointCheckDomain(url.Values{
		"domain": []string{domain},
	})).Error; err != nil {
		fmt.Println(err)
		return
	}

	body, err := a.do()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", string(body))
}
