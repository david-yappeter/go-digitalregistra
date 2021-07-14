package digitalregistra

import (
	"encoding/xml"
	"fmt"
)

func (a *API) CreateContact(input CreateContactParam) (*CreateContactResponse, error) {
	if err := a.newRequest(a.endpointCreateContact(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CreateContactResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UpdateContact(input UpdateContactParam) (*UpdateContactResponse, error) {
	if err := a.newRequest(a.endpointUpdateContact(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UpdateContactResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
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

func (a *API) DeleteContact(input DeleteContactParam) (*InfoContactResponse, error) {
	if err := a.newRequest(a.endpointDeleteContact(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	// var responseModel InfoContactResponse
	// err = xml.Unmarshal(body, &responseModel)
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Printf("%+v\n", string(body))

	// return &responseModel, nil
	return nil, nil
}
