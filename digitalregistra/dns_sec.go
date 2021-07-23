package digitalregistra

import "encoding/xml"

func (a *API) ListDNSSECRecords(input ListDNSSECRecordsParam) (*ListDNSSECRecordsResponse, error) {
	if err := a.newRequest(a.endpointListDNSSECRecords(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel ListDNSSECRecordsResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}


func (a *API) AddDNSSECRecords(input AddDNSSECRecordsParam) (*AddDNSSECRecordsResponse, error) {
	if err := a.newRequest(a.endpointAddDNSSECRecords(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel AddDNSSECRecordsResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}


func (a *API) DeleteDNSSECRecords(input DeleteDNSSECRecordsParam) (*DeleteDNSSECRecordsResponse, error) {
	if err := a.newRequest(a.endpointDeleteDNSSECRecords(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel DeleteDNSSECRecordsResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
