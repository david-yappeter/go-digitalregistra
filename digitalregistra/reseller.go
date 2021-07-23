package digitalregistra

import "encoding/xml"

func (a *API) RequestDepositReseller(input RequestDepositResellerParam) (*RequestDepositResellerResponse, error) {
	if err := a.newRequest(a.endpointRequestDepositReseller(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel RequestDepositResellerResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) CekResellerFund() (*CekResellerFundResponse, error) {
	if err := a.newRequest(a.endpointCekResellerFund(nil)).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CekResellerFundResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) GetResellerPrice() (*GetResellerPriceResponse, error) {
	if err := a.newRequest(a.endpointGetResellerPrice(nil)).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel GetResellerPriceResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) AddSubresellerDirectFund(input AddSubresellerDirectFundParam) (*AddSubresellerDirectFundResponse, error) {
	if err := a.newRequest(a.endpointAddSubresellerDirectFund(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel AddSubresellerDirectFundResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) ListSubreseller(input ListSubresellerParam) (*ListSubresellerResponse, error) {
	if err := a.newRequest(a.endpointListSubreseller(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel ListSubresellerResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) InfoSubreseller(input InfoSubresellerParam) (*InfoSubresellerResponse, error) {
	if err := a.newRequest(a.endpointInfoSubreseller(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoSubresellerResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
