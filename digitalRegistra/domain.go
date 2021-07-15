package digitalregistra

import (
	"encoding/xml"
)

func (a *API) RegisterDomain(input RegisterDomainParam) (*RegisterDomainResponse, error) {
	if err := a.newRequest(a.endpointDomainRegistration(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel RegisterDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) TransferDomain(input TransferDomainParam) (*TransferDomainResponse, error) {
	if err := a.newRequest(a.endpointTransferDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel TransferDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
func (a *API) LockDomain(input LockDomainParam) (*LockDomainResponse, error) {
	if err := a.newRequest(a.endpointLockDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel LockDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
func (a *API) UnlockDomain(input UnlockDomainParam) (*UnlockDomainResponse, error) {
	if err := a.newRequest(a.endpointUnlockDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UnlockDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) RenewDomain(input RenewDomainParam) (*RenewDomainParam, error) {
	if err := a.newRequest(a.endpointRenewDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel RenewDomainParam
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) InfoDomain(input InfoDomainParam) (*InfoDomainResponse, error) {
	if err := a.newRequest(a.endpointInfoDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel InfoDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UpdateNameserver(input UpdateNameserverParam) (*UpdateNameserverResponse, error) {
	if err := a.newRequest(a.endpointUpdateNameserver(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UpdateNameserverResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UploadDocumentLink(input UploadDocumentLinkParam) (*UploadDocumentLinkResponse, error) {
	if err := a.newRequest(a.endpointUploadDocumentLink(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UploadDocumentLinkResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UpdateContactDomain(input UpdateContactDomainParam) (*UploadDocumentLinkResponse, error) {
	if err := a.newRequest(a.endpointUpdateContactDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UploadDocumentLinkResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) ExpressRegister(input ExpressRegisterParam) (*ExpressRegisterResponse, error) {
	if err := a.newRequest(a.endpointExpressRegister(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel ExpressRegisterResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) ExpressTransfer(input ExpressTransferParam) (*ExpressTransferResponse, error) {
	if err := a.newRequest(a.endpointExpressTransfer(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel ExpressTransferResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) SuspendDomain(input SuspendDomainParam) (*SuspendDomainResponse, error) {
	if err := a.newRequest(a.endpointSuspendDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel SuspendDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) UnsuspendDomain(input UnsuspendDomainParam) (*UnsuspendDomainResponse, error) {
	if err := a.newRequest(a.endpointUnsuspendDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel UnsuspendDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) SetIDProtection(input SetIDProtectionParam) (*SetIDProtectionResponse, error) {
	if err := a.newRequest(a.endpointSetIDProtection(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel SetIDProtectionResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) CheckVerificationInfo(input CheckVerificationInfoParam) (*CheckVerificationInfoResponse, error) {
	if err := a.newRequest(a.endpointCheckVerificationInfo(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CheckVerificationInfoResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) ResendVerificationEmail(input ResendVerificationEmailParam) (*ResendVerificationEmailResponse, error) {
	if err := a.newRequest(a.endpointResendVerificationEmail(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel ResendVerificationEmailResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) CancelDomain(input CancelDomainParam) (*CancelDomainResponse, error) {
	if err := a.newRequest(a.endpointCancelDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CancelDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) CheckDomain(input CheckDomainParam) (*CheckDomainResponse, error) {
	if err := a.newRequest(a.endpointCheckDomain(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel CheckDomainResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) PremiumExpressRegister(input PremiumExpressRegisterParam) (*PremiumExpressRegisterResponse, error) {
	if err := a.newRequest(a.endpointPremiumExpressRegister(ModelToURLValues(input))).Error; err != nil {
		return nil, err
	}

	body, err := a.do()
	if err != nil {
		return nil, err
	}

	var responseModel PremiumExpressRegisterResponse
	err = xml.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}
