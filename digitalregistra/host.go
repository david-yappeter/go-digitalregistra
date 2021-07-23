package digitalregistra

import "encoding/xml"

func (a *API) CreateHost(input CreateHostParam) (*CreateHostResponse, error) {
	if err := a.newRequest(a.endpointCreateHost(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CreateHostResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}


func (a *API) UpdateHost(input UpdateHostParam) (*UpdateHostResponse, error) {
	if err := a.newRequest(a.endpointUpdateHost(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UpdateHostResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}


func (a *API) InfoHost(input InfoHostParam) (*InfoHostResponse, error) {
	if err := a.newRequest(a.endpointInfoHost(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoHostResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}


func (a *API) DeleteHost(input DeleteHostParam) (*DeleteHostResponse, error) {
	if err := a.newRequest(a.endpointDeleteHost(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel DeleteHostResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
