package digitalregistra

// DOMAIN
// ==============================================================

type CheckDomainParam struct {
	Domain string `json:"domain" digitalregistra:"domain"`
}

type RegisterDomainParam struct {
	Domain         string  `json:"domain" digitalregistra:"domain"`
	ApiID          int     `json:"api_id" digitalregistra:"api_id"`
	Periode        int     `json:"periode" digitalregistra:"periode"`
	NS1            string  `json:"ns1" digitalregistra:"ns1"`
	NS2            string  `json:"ns2" digitalregistra:"ns2"`
	NS3            *string `json:"ns3" digitalregistra:"ns3"`
	NS4            *string `json:"ns4" digitalregistra:"ns4"`
	Firstname      string  `json:"first_name" digitalregistra:"fname"`
	Lastname       string  `json:"last_name" digitalregistra:"lname"`
	Company        *string `json:"company" digitalregistra:"company"`
	Address1       string  `json:"address1" digitalregistra:"address1"`
	Address2       *string `json:"address2" digitalregistra:"address2"`
	City           string  `json:"city" digitalregistra:"city"`
	State          string  `json:"state" digitalregistra:"state"`
	Country        string  `json:"country" digitalregistra:"country"`
	PostCode       string  `json:"post_code" digitalregistra:"postcode"`
	PhoneNumber    string  `json:"phone_number" digitalregistra:"phonenumber"`
	Email          string  `json:"email" digitalregistra:"email"`
	UserUsername   string  `json:"user_username" digitalregistra:"user_username"`
	UserFirstname  string  `json:"user_firstname" digitalregistra:"user_fname"`
	UserLastname   *string `json:"user_lastname" digitalregistra:"user_lname"`
	UserCompany    string  `json:"user_company" digitalregistra:"user_company"`
	UserAddress    string  `json:"user_address" digitalregistra:"user_address"`
	UserAddress2   *string `json:"user_address2" digitalregistra:"user_address2"`
	UserCity       string  `json:"user_city" digitalregistra:"user_city"`
	UserProvince   string  `json:"user_province" digitalregistra:"user_province"`
	UserCountry    string  `json:"user_country" digitalregistra:"user_country"`
	UserPostalCode string  `json:"user_postal_code" digitalregistra:"user_postal_code"`
	AutoActive     *string `json:"auto_active" digitalregistra:"autoactive"`
	Category       string  `json:"category" digitalregistra:"category"`
}

// CONTACT
// ==============================================================

type CreateContactParam struct {
	UserID           string  `json:"user_id" digitalregistra:"userid"`
	Firstname        string  `json:"firstname" digitalregistra:"fname"`
	Lastname         *string `json:"lastname" digitalregistra:"lname"`
	Email            string  `json:"email" digitalregistra:"email"`
	Company          string  `json:"company" digitalregistra:"company"`
	Address          string  `json:"address" digitalregistra:"address"`
	Address2         *string `json:"address2" digitalregistra:"address2"`
	Address3         *string `json:"address3" digitalregistra:"address3"`
	City             string  `json:"city" digitalregistra:"city"`
	Province         string  `json:"provice" digitalregistra:"province"`
	PhoneCountryCode *string `json:"phone_country_code" digitalregistra:"phone-cc"`
	Phone            string  `json:"phone" digitalregistra:"phone"`
	FaxCountryCode   *string `json:"fax_country_code" digitalregistra:"fax-cc"`
	Fax              *string `json:"fax" digitalregistra:"fax"`
	Country          string  `json:"country" digitalregistra:"country"`
	PostalCode       string  `json:"postal_code" digitalregistra:"postal_code"`
}

type UpdateContactParam struct {
	ContactID        *string `json:"user_id" digitalregistra:"contactid"`
	NickHandle       *string `json:"nick_handle" digitalregistra:"nickhandle"`
	Firstname        *string `json:"firstname" digitalregistra:"fname"`
	Lastname         *string `json:"lastname" digitalregistra:"lname"`
	Email            *string `json:"email" digitalregistra:"email"`
	Company          *string `json:"company" digitalregistra:"company"`
	Address          *string `json:"address" digitalregistra:"address"`
	Address2         *string `json:"address2" digitalregistra:"address2"`
	Address3         *string `json:"address3" digitalregistra:"address3"`
	City             *string `json:"city" digitalregistra:"city"`
	Province         *string `json:"provice" digitalregistra:"province"`
	PhoneCountryCode *string `json:"phone_country_code" digitalregistra:"phone-cc"`
	Phone            *string `json:"phone" digitalregistra:"phone"`
	FaxCountryCode   *string `json:"fax_country_code" digitalregistra:"fax-cc"`
	Fax              *string `json:"fax" digitalregistra:"fax"`
	Country          *string `json:"country" digitalregistra:"country"`
	PostalCode       *string `json:"postal_code" digitalregistra:"postal_code"`
}

type InfoContactParam struct {
	ContactID  *string `json:"contact_id" digitalregistra:"contactid"`
	NickHandle *string `json:"nick_handle" digitalregistra:"nickhandle"`
}

type DeleteContactParam struct {
	ContactID string `json:"contact_id" digitalregistra:"contactid"`
}

// USER
// ==============================================================

type CreateUserParam struct {
	Username         string  `json:"user_name" digitalregistra:"user_username"`
	Password         string  `json:"password" digitalregistra:"user_password"`
	Email            *string `json:"email" digitalregistra:"email"`
	Firstname        string  `json:"firstname" digitalregistra:"fname"`
	Company          *string `json:"company" digitalregistra:"company"`
	Address          string  `json:"address" digitalregistra:"address"`
	Address2         *string `json:"address2" digitalregistra:"address2"`
	Address3         *string `json:"address3" digitalregistra:"address3"`
	City             string  `json:"city" digitalregistra:"city"`
	Province         string  `json:"province" digitalregistra:"province"`
	PostalCode       string  `json:"postal_code" digitalregistra:"postal_code"`
	Country          string  `json:"country" digitalregistra:"country"`
	PhoneCountryCode string  `json:"phone_country_code" digitalregistra:"p_kdnegara"`
	PhoneAreaCode    string  `json:"phone_area_code" digitalregistra:"p_kdarea"`
	Phone            string  `json:"phone" digitalregistra:"phone"`
	FaxCountryCode   *string `json:"fax_country_code" digitalregistra:"f_kdnegara"`
	FaxAreaCode      *string `json:"fax_area_code" digitalregistra:"f_kdarea"`
	Fax              *string `json:"fax" digitalregistra:"fax"`
}

type InfoUserParam struct {
	UserID   *int    `json:"user_id" digitalregistra:"userid"`
	Username *string `json:"username" digitalregistra:"user_username"`
}

type UpdateUserParam struct {
	UserID           string  `json:"user_id" digitalregistra:"userid"`
	Username         *string `json:"username" digitalregistra:"user_username"`
	Email            *string `json:"email" digitalregistra:"email"`
	Firstname        *string `json:"firstname" digitalregistra:"fname"`
	Company          *string `json:"company" digitalregistra:"company"`
	Address          *string `json:"address" digitalregistra:"address"`
	Address2         *string `json:"address2" digitalregistra:"address2"`
	Address3         *string `json:"address3" digitalregistra:"address3"`
	City             *string `json:"city" digitalregistra:"city"`
	Province         *string `json:"province" digitalregistra:"province"`
	PostalCode       *string `json:"postal_code" digitalregistra:"postal_code"`
	Country          *string `json:"country" digitalregistra:"country"`
	PhoneCountryCode *string `json:"phone_country_code" digitalregistra:"p_kdnegara"`
	PhoneAreaCode    *string `json:"phone_area_code" digitalregistra:"p_kdarea"`
	Phone            *string `json:"phone" digitalregistra:"phone"`
	FaxCountryCode   *string `json:"fax_country_code" digitalregistra:"f_kdnegara"`
	FaxAreaCode      *string `json:"fax_area_code" digitalregistra:"f_kdarea"`
	Fax              *string `json:"fax" digitalregistra:"fax"`
}
