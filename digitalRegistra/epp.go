package digitalregistra

import "encoding/xml"

func (a *API) GetEpp(input GetEppParam) (*GetEppResponse, error) {
	if err := a.newRequest(a.endpointGetEpp(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel GetEppResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) SetEpp(input SetEppParam) (*SetEppResponse, error) {
	if err := a.newRequest(a.endpointSetEpp(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel SetEppResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
