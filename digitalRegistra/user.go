package digitalregistra

import (
	"encoding/xml"
)

func (a *API) CreateUser(input CreateUserParam) (*CreateUserResponse, error) {
	if err := a.newRequest(a.endpointCreateUser(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CreateUserResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) InfoUser(input InfoUserParam) (*InfoUserResponse, error) {
	if err := a.newRequest(a.endpointInfoUser(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoUserResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UpdateUser(input UpdateUserParam) (*UpdateUserResponse, error) {
	if err := a.newRequest(a.endpointUpdateUser(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UpdateUserResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
