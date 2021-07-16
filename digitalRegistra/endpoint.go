package digitalregistra

import (
	"fmt"
	"net/http"
	"net/url"
)

// DOMAIN
// ====================================================================
func (a *API) endpointDomainRegistration(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/register"), data
}

func (a *API) endpointTransferDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/transfer"), data
}

func (a *API) endpointLockDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/lock"), data
}

func (a *API) endpointUnlockDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/unlock"), data
}

func (a *API) endpointRenewDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/renew"), data
}

func (a *API) endpointInfoDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/info"), data
}

func (a *API) endpointUpdateNameserver(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/updatens"), data
}

func (a *API) endpointUploadDocumentLink(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/documenturl"), data
}

func (a *API) endpointUpdateContactDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/editcontact"), data
}
func (a *API) endpointExpressRegister(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/expressregister"), data
}
func (a *API) endpointExpressTransfer(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/expresstransfer"), data
}
func (a *API) endpointSuspendDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/suspend"), data
}

func (a *API) endpointUnsuspendDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/unsuspend"), data
}

func (a *API) endpointCheckVerificationInfo(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/veri_info"), data
}

func (a *API) endpointSetIDProtection(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/set_idprotection"), data
}

func (a *API) endpointResendVerificationEmail(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/resend_raa"), data
}

func (a *API) endpointCancelDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/cancel"), data
}

func (a *API) endpointCheckDomain(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/check"), data
}

func (a *API) endpointPremiumExpressRegister(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/premium_expressregister"), data
}

// EPP
// ====================================================================
func (a *API) endpointGetEpp(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/getepp"), data
}

func (a *API) endpointSetEpp(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/setepp"), data
}

// Contact
// ====================================================================
func (a *API) endpointCreateContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/create"), data
}

func (a *API) endpointUpdateContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/update"), data
}

func (a *API) endpointInfoContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/info"), data
}

func (a *API) endpointDeleteContact(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "contact/delete"), data
}

// HOST
// ====================================================================
func (a *API) endpointCreateHost(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "host/addchild"), data
}

func (a *API) endpointUpdateHost(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "host/updatechild"), data
}

func (a *API) endpointInfoHost(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "host/info"), data
}

func (a *API) endpointDeleteHost(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "host/delchild"), data
}


// RESELLER
// ====================================================================
func (a *API) endpointRequestDepositReseller(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "reseller/requestdeposit"), data
}

func (a *API) endpointCekResellerFund(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "reseller/cekfund"), data
}


func (a *API) endpointGetResellerPrice(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "reseller/getprice"), data
}

func (a *API) endpointAddSubresellerDirectFund(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "billing/addsubresellerfund"), data
}

func (a *API) endpointListSubreseller(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "v2/reseller/subresellers"), data
}

func (a *API) endpointInfoSubreseller(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "v2/reseller/subreseller"), data
}

// User
// ====================================================================

func (a *API) endpointCreateUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/create"), data
}

func (a *API) endpointInfoUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/info"), data
}

func (a *API) endpointUpdateUser(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "user/update"), data
}

// MOVE SERVICE
// ====================================================================

func (a *API) endpointMoveService(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/moveservice"), data
}

// DNS
// ====================================================================

func (a *API) endpointInfoDNS(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "dns/info"), data
}

func (a *API) endpointInisialisasiDNS(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "dns/start"), data
}

func (a *API) endpointEditDNSRecord(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "dns/edit"), data
}

func (a *API) endpointCreateDNSRecord(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "dns/create"), data
}

func (a *API) endpointDeleteDNSRecord(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "dns/delete"), data
}

// FORWARD
// ====================================================================

func (a *API) endpointInitialForwarder(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "forward/configure"), data
}

func (a *API) endpointStatusForward(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "forward/status"), data
}

func (a *API) endpointUpdateForward(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "forward/update"), data
}

// DNSSEC
// ====================================================================

func (a *API) endpointListDNSSECRecords(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/listds"), data
}

func (a *API) endpointAddDNSSECRecords(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/addds"), data
}

func (a *API) endpointDeleteDNSSECRecords(data url.Values) (string, string, url.Values) {
	return http.MethodPost, fmt.Sprintf("%s/api/%s", a.DefaultUrl, "domain/delds"), data
}
