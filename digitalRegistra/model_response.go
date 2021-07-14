package digitalregistra

import "encoding/xml"

// DOMAIN
// =============================================================

type DomainCheckResponse struct {
	XMLName xml.Name `xml:"epp"`
	Text    string   `xml:",chardata"`
	Result  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"result"`
	ResultData string `xml:"resultData"`
}

// CONTACT
// =============================================================

type CreateContactResponse struct {
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

type UpdateContactResponse struct {
	XMLName xml.Name `xml:"epp"`
	Text    string   `xml:",chardata"`
	Result  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"result"`
	ResultData string `xml:"resultData"`
}

type InfoContactResponse struct {
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

// USER
// =============================================================

type CreateUserResponse struct {
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
type InfoUserResponse struct {
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

type UpdateUserResponse struct {
	XMLName xml.Name `xml:"epp"`
	Text    string   `xml:",chardata"`
	Result  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"result"`
	ResultData string `xml:"resultData"`
}
