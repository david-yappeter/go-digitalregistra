package digitalregistra

import (
	"encoding/xml"
	"fmt"
)

func (a *API) CheckDomain(input CheckDomainParam) (*DomainCheckResponse, error) {
	if err := a.newRequest(a.endpointCheckDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel DomainCheckResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) RegisterDomain(input RegisterDomainParam) {
	if err := a.newRequest(a.endpointDomainRegistration(ModelToURLValues(input))).Error; err != nil {
		return
	}

	body, err := a.do()
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", string(body))
	// var responseModel DomainCheckResponse
	// err = xml.Unmarshal(body, &responseModel)
	// if err != nil {
	// 	return nil, err
	// }

	// return &responseModel, nil
}
