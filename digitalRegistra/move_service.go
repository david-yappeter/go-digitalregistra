package digitalregistra

import "encoding/xml"

func (a *API) MoveService(input MoveServiceParam) (*MoveServiceResponse, error) {
	if err := a.newRequest(a.endpointMoveService(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel MoveServiceResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
