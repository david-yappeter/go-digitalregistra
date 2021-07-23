package digitalregistra

import "encoding/xml"

func (a *API) InitialForwarder(input InitialForwarderParam) (*InitialForwarderResponse, error) {
	if err := a.newRequest(a.endpointInitialForwarder(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InitialForwarderResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) StatusForward(input StatusForwardParam) (*StatusForwardResponse, error) {
	if err := a.newRequest(a.endpointStatusForward(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel StatusForwardResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UpdateForward(input UpdateForwardParam) (*UpdateForwardResponse, error) {
	if err := a.newRequest(a.endpointUpdateForward(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UpdateForwardResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
