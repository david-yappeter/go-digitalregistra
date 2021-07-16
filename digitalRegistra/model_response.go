package digitalregistra

import "encoding/xml"

// DOMAIN
// =============================================================

type (
	RegisterDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text      string `xml:",chardata"`
			Productid string `xml:"productid"`
			Invoiceid string `xml:"invoiceid"`
		} `xml:"resultData"`
	}

	TransferDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	LockDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	UnlockDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	RenewDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text   string `xml:",chardata"`
			ExDate string `xml:"exDate"`
		} `xml:"resultData"`
	}

	InfoDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text              string `xml:",chardata"`
			Domainid          string `xml:"domainid"`
			ApiID             string `xml:"api_id"`
			Userid            string `xml:"userid"`
			Username          string `xml:"username"`
			Domain            string `xml:"domain"`
			Startdate         string `xml:"startdate"`
			Enddate           string `xml:"enddate"`
			Unixstartdate     string `xml:"unixstartdate"`
			Unixenddate       string `xml:"unixenddate"`
			Canceldate        string `xml:"canceldate"`
			ContactRegistrant string `xml:"contact_registrant"`
			ContactAdmin      string `xml:"contact_admin"`
			ContactBilling    string `xml:"contact_billing"`
			ContactTech       string `xml:"contact_tech"`
			Ns1               string `xml:"ns1"`
			Ns2               string `xml:"ns2"`
			Ns3               string `xml:"ns3"`
			Ns4               string `xml:"ns4"`
			Ns5               string `xml:"ns5"`
			Ns6               string `xml:"ns6"`
			Ns7               string `xml:"ns7"`
			Ns8               string `xml:"ns8"`
			Idprotection      string `xml:"idprotection"`
			Status            string `xml:"status"`
			Authcode          string `xml:"authcode"`
			Suspended         string `xml:"suspended"`
			Transferlock      string `xml:"transferlock"`
			Resellerlock      string `xml:"resellerlock"`
			Registrarlock     string `xml:"registrarlock"`
		} `xml:"resultData"`
	}

	UpdateNameserverResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	UploadDocumentLinkResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url"`
		} `xml:"resultData"`
	}

	UpdateContactDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	ExpressRegisterResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text   string `xml:",chardata"`
			Status string `xml:"status"`
			CrDate string `xml:"crDate"`
			ExDate string `xml:"exDate"`
		} `xml:"resultData"`
	}

	ExpressTransferResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	SuspendDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	UnsuspendDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}
	SetIDProtectionResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	CheckVerificationInfoResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text                     string `xml:",chardata"`
			RaaVerificationStatus    string `xml:"raaVerificationStatus"`
			RegistrantEmail          string `xml:"registrantEmail"`
			RaaVerificationStartTime string `xml:"raaVerificationStartTime"`
			RaaVerificationEndTime   string `xml:"raaVerificationEndTime"`
		} `xml:"resultData"`
	}

	ResendVerificationEmailResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	CancelDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	CheckDomainResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	PremiumExpressRegisterResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text   string `xml:",chardata"`
			Status string `xml:"status"`
			CrDate string `xml:"crDate"`
			ExDate string `xml:"exDate"`
		} `xml:"resultData"`
	}
)

// EPP
// =============================================================
type (
	GetEppResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text string `xml:",chardata"`
			Epp  string `xml:"epp"`
		} `xml:"resultData"`
	}

	SetEppResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}
)

// RESELLER
// =============================================================
type (
	RequestDepositResellerResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	CekResellerFundResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text    string `xml:",chardata"`
			Deposit string `xml:"deposit"`
		} `xml:"resultData"`
	}

	GetResellerPriceResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text string `xml:",chardata"`
			File string `xml:"file"`
		} `xml:"resultData"`
	}

	AddSubresellerDirectFundResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	ListSubresellerResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text        string `xml:",chardata"`
			ResellerNum string `xml:"resellerNum"`
			Reseller    []struct {
				Text       string `xml:",chardata"`
				Username   string `xml:"username"`
				Email      string `xml:"email"`
				Fname      string `xml:"fname"`
				Lname      string `xml:"lname"`
				Company    string `xml:"company"`
				Address    string `xml:"address"`
				Address2   string `xml:"address2"`
				Address3   string `xml:"address3"`
				City       string `xml:"city"`
				Province   string `xml:"province"`
				Phone      string `xml:"phone"`
				Mphone     string `xml:"mphone"`
				Fax        string `xml:"fax"`
				Country    string `xml:"country"`
				PostalCode string `xml:"postal_code"`
				TaxStatus  string `xml:"tax_status"`
				Ns1        string `xml:"ns1"`
				Ns2        string `xml:"ns2"`
				Ns3        string `xml:"ns3"`
				Ns4        string `xml:"ns4"`
				Ns5        string `xml:"ns5"`
				Ns6        string `xml:"ns6"`
				Ns7        string `xml:"ns7"`
				Ns8        string `xml:"ns8"`
			} `xml:"reseller"`
		} `xml:"resultData"`
	}

	InfoSubresellerResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text        string `xml:",chardata"`
			ResellerNum string `xml:"resellerNum"`
			Reseller    struct {
				Text       string   `xml:",chardata"`
				Username   []string `xml:"username"`
				Email      string   `xml:"email"`
				Fname      string   `xml:"fname"`
				Lname      string   `xml:"lname"`
				Company    string   `xml:"company"`
				Address    string   `xml:"address"`
				Address2   string   `xml:"address2"`
				Address3   string   `xml:"address3"`
				City       string   `xml:"city"`
				Province   string   `xml:"province"`
				Phone      string   `xml:"phone"`
				Mphone     string   `xml:"mphone"`
				Fax        string   `xml:"fax"`
				Country    string   `xml:"country"`
				PostalCode string   `xml:"postal_code"`
				TaxStatus  string   `xml:"tax_status"`
				Ns1        string   `xml:"ns1"`
				Ns2        string   `xml:"ns2"`
				Ns3        string   `xml:"ns3"`
				Ns4        string   `xml:"ns4"`
				Ns5        string   `xml:"ns5"`
				Ns6        string   `xml:"ns6"`
				Ns7        string   `xml:"ns7"`
				Ns8        string   `xml:"ns8"`
			} `xml:"reseller"`
		} `xml:"resultData"`
	}
)

// CONTACT
// =============================================================
type (
	CreateContactResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text       string `xml:",chardata"`
			Nickhandle string `xml:"nickhandle"`
			Epphandle  string `xml:"epphandle"`
			Contactid  string `xml:"contactid"`
		} `xml:"resultData"`
	}

	UpdateContactResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	InfoContactResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text        string `xml:",chardata"`
			Contactid   string `xml:"contactid"`
			Userid      string `xml:"userid"`
			Nickhandle  string `xml:"nickhandle"`
			Fname       string `xml:"fname"`
			Lname       string `xml:"lname"`
			Email       string `xml:"email"`
			Company     string `xml:"company"`
			Address1    string `xml:"address1"`
			Address2    string `xml:"address2"`
			Address3    string `xml:"address3"`
			City        string `xml:"city"`
			State       string `xml:"state"`
			Phonenumber string `xml:"phonenumber"`
			Fax         string `xml:"fax"`
			Country     string `xml:"country"`
			Postcode    string `xml:"postcode"`
			Crdate      string `xml:"crdate"`
			Update      string `xml:"update"`
			Authcode    string `xml:"authcode"`
		} `xml:"resultData"`
	}
)

// USER
// =============================================================
type (
	CreateHostResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	UpdateHostResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	InfoHostResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}

	DeleteHostResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}
)

// USER
// =============================================================
type (
	CreateUserResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text   string `xml:",chardata"`
			Userid string `xml:"userid"`
		} `xml:"resultData"`
	}
	InfoUserResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text               string `xml:",chardata"`
			Userid             string `xml:"userid"`
			Username           string `xml:"username"`
			Email              string `xml:"email"`
			Fname              string `xml:"fname"`
			Lname              string `xml:"lname"`
			Company            string `xml:"company"`
			Address            string `xml:"address"`
			Address2           string `xml:"address2"`
			Address3           string `xml:"address3"`
			City               string `xml:"city"`
			Province           string `xml:"province"`
			Phone              string `xml:"phone"`
			Fax                string `xml:"fax"`
			PostalCode         string `xml:"postal_code"`
			Country            string `xml:"country"`
			Postcode           string `xml:"postcode"`
			Faktur             string `xml:"faktur"`
			Npwp               string `xml:"npwp"`
			WajibPajak         string `xml:"wajib_pajak"`
			AlamatWajibPajak   string `xml:"alamat_wajib_pajak"`
			KategoriWajibPajak string `xml:"kategori_wajib_pajak"`
			UserStatus         string `xml:"user_status"`
		} `xml:"resultData"`
	}

	UpdateUserResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}
)

// MOVE SERVICE
// =============================================================
type (
	MoveServiceResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}
)

// DNS
// =============================================================
type (
	InfoDNSResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text string `xml:",chardata"`
			Dns0 struct {
				Text        string `xml:",chardata"`
				Dnsid       string `xml:"dnsid"`
				Line        string `xml:"line"`
				Domainid    string `xml:"domainid"`
				Domain      string `xml:"domain"`
				Class       string `xml:"class"`
				Type        string `xml:"type"`
				Record      string `xml:"record"`
				Destination string `xml:"destination"`
				Ttl         string `xml:"ttl"`
				Priority    string `xml:"priority"`
				Weight      string `xml:"weight"`
				Port        string `xml:"port"`
				Mname       string `xml:"mname"`
				Rname       string `xml:"rname"`
				Refresh     string `xml:"refresh"`
				Retry       string `xml:"retry"`
				Serial      string `xml:"serial"`
				Expire      string `xml:"expire"`
				Minimum     string `xml:"minimum"`
				Preference  string `xml:"preference"`
			} `xml:"dns0"`
			Dns1 struct {
				Text        string `xml:",chardata"`
				Dnsid       string `xml:"dnsid"`
				Line        string `xml:"line"`
				Domainid    string `xml:"domainid"`
				Domain      string `xml:"domain"`
				Class       string `xml:"class"`
				Type        string `xml:"type"`
				Record      string `xml:"record"`
				Destination string `xml:"destination"`
				Ttl         string `xml:"ttl"`
				Priority    string `xml:"priority"`
				Weight      string `xml:"weight"`
				Port        string `xml:"port"`
				Mname       string `xml:"mname"`
				Rname       string `xml:"rname"`
				Refresh     string `xml:"refresh"`
				Retry       string `xml:"retry"`
				Serial      string `xml:"serial"`
				Expire      string `xml:"expire"`
				Minimum     string `xml:"minimum"`
				Preference  string `xml:"preference"`
			} `xml:"dns1"`
			Dns2 struct {
				Text        string `xml:",chardata"`
				Dnsid       string `xml:"dnsid"`
				Line        string `xml:"line"`
				Domainid    string `xml:"domainid"`
				Domain      string `xml:"domain"`
				Class       string `xml:"class"`
				Type        string `xml:"type"`
				Record      string `xml:"record"`
				Destination string `xml:"destination"`
				Ttl         string `xml:"ttl"`
				Priority    string `xml:"priority"`
				Weight      string `xml:"weight"`
				Port        string `xml:"port"`
				Mname       string `xml:"mname"`
				Rname       string `xml:"rname"`
				Refresh     string `xml:"refresh"`
				Retry       string `xml:"retry"`
				Serial      string `xml:"serial"`
				Expire      string `xml:"expire"`
				Minimum     string `xml:"minimum"`
				Preference  string `xml:"preference"`
			} `xml:"dns2"`
			DomainNs1   string `xml:"domain_ns1"`
			DomainNs2   string `xml:"domain_ns2"`
			DomainNs3   string `xml:"domain_ns3"`
			DomainNs4   string `xml:"domain_ns4"`
			ResellerNs1 string `xml:"reseller_ns1"`
			ResellerNs2 string `xml:"reseller_ns2"`
			ResellerNs3 string `xml:"reseller_ns3"`
			ResellerNs4 string `xml:"reseller_ns4"`
		} `xml:"resultData"`
	}

	InisialisasiDNSResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	EditDNSRecordResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	CreateDNSRecordResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	DeleteDNSRecordResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
	}
)

// FORWARD
// =============================================================
type (
	InitialForwarderResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	StatusForwardResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text      string `xml:",chardata"`
			Target    string `xml:"target"`
			Type      string `xml:"type"`
			Header    string `xml:"header"`
			Noframe   string `xml:"noframe"`
			Subdomain string `xml:"subdomain"`
			Path      string `xml:"path"`
		} `xml:"resultData"`
	}

	UpdateForwardResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}
)

// DNSSEC
// =============================================================
type (
	ListDNSSECRecordsResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData struct {
			Text   string `xml:",chardata"`
			Dnssec struct {
				Text  string `xml:",chardata"`
				Value []struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id"`
					Keytag     string `xml:"keytag"`
					Algorithm  string `xml:"algorithm"`
					Digesttype string `xml:"digesttype"`
					Digest     string `xml:"digest"`
				} `xml:"value"`
			} `xml:"dnssec"`
			Domain string `xml:"domain"`
		} `xml:"resultData"`
	}

	AddDNSSECRecordsResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}

	DeleteDNSSECRecordsResponse struct {
		XMLName xml.Name `xml:"epp"`
		Text    string   `xml:",chardata"`
		Result  struct {
			Text       string `xml:",chardata"`
			ResultCode string `xml:"resultCode"`
			ResultMsg  string `xml:"resultMsg"`
		} `xml:"result"`
		ResultData string `xml:"resultData"`
	}
)
