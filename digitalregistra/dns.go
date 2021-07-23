package digitalregistra

import "encoding/xml"

func (a *API) InfoDNS(input InfoDNSParam) (*InfoDNSResponse, error) {
	if err := a.newRequest(a.endpointInfoDNS(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoDNSResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) InisialisasiDNS(input InisialisasiDNSParam) (*InisialisasiDNSResponse, error) {
	if err := a.newRequest(a.endpointInisialisasiDNS(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InisialisasiDNSResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) EditDNSRecord(input EditDNSRecordParam) (*EditDNSRecordResponse, error) {
	if err := a.newRequest(a.endpointEditDNSRecord(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel EditDNSRecordResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) CreateDNSRecord(input CreateDNSRecordParam) (*CreateDNSRecordResponse, error) {
	if err := a.newRequest(a.endpointCreateDNSRecord(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CreateDNSRecordResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) DeleteDNSRecord(input DeleteDNSRecordParam) (*DeleteDNSRecordResponse, error) {
	if err := a.newRequest(a.endpointDeleteDNSRecord(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel DeleteDNSRecordResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
