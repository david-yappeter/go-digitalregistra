package digitalregistra

import "encoding/xml"

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
