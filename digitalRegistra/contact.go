package digitalregistra

import (
	"encoding/xml"
	"fmt"
)

func (a *API) CreateContact(input CreateContactParam) {
	if err := a.newRequest(a.endpointCreateContact(ModelToURLValues(input))).Error; err != nil {
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
	// 	return
	// }

	// return &responseModel, nil
}

func (a *API) InfoContact(input InfoContactParam) (*InfoContactResponse, error) {
	if err := a.newRequest(a.endpointInfoContact(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoContactResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
