package main

import (
	digitalregistra "myapp/digitalregistra"
	"os"
)

func main() {
	api := digitalregistra.NewAPI(os.Getenv("DIGITAL_REGISTRA_RESELLER_URL"))

	api.UpdateNameserver(digitalregistra.UpdateNameserverParam{
		Domain: "testerdomainkuuuuuu.co.id",
		ApiID:  123,
	})

	// test, err := api.CancelDomain(digitalregistra.CancelDomainParam{
	// 	Domain: "testerdomainkuuuuuu.co.id",
	// 	ApiID:  123,
	// })

	// test, err := api.InfoDomain(digitalregistra.InfoDomainParam{
	// 	Domain: "testerdomainkuuuuuu.co.id",
	// 	ApiID:  123,
	// })

	// fmt.Printf("%+v\n", test)
	// fmt.Printf("%+v\n", err)

	// // resppp, errr := api.DomainRegistration(digitalregistra.RegisterDomainParam{
	// // 	Domain: "testerrrrrrr.co.id",
	// //     ApiID: 123123512,
	// //     Periode: ,
	// // })
	// // fmt.Printf("%+v\n", resppp)
	// // fmt.Printf("%+v\n", errr)

	// a := "011102407z32"
	// resppp, errr := api.InfoContact(digitalregistra.InfoContactParam{
	// 	// ContactID: &a,
	// 	NickHandle: &a,
	// })

	// fmt.Printf("%+v\n", resppp)
	// fmt.Printf("%+v\n", errr)

	// api.DeleteContact(digitalregistra.DeleteContactParam{
	// 	ContactID: "110240",
	// })

	// a := "qoweqiwoeiq.c"
	// tempp := digitalregistra.RegisterDomainParam{
	// 	Domain:         "testerdomainkuuuuuu.co.id",
	// 	ApiID:          123,
	// 	Periode:        2,
	// 	NS1:            "asdasdaasdasdsa",
	// 	NS2:            "bbdfjbidfjbofd",
	// 	NS3:            new(string),
	// 	NS4:            &a,
	// 	Firstname:      "asda",
	// 	Lastname:       "asgh",
	// 	Address1:       "asdas",
	// 	City:           "Med",
	// 	State:          "Su",
	// 	Country:        "ID",
	// 	PostCode:       "20000",
	// 	PhoneNumber:    "6285377778888",
	// 	Email:          "daviddummy2002@gmail.com",
	// 	UserUsername:   "daviddummy2002@gmail.com",
	// 	UserFirstname:  "dav",
	// 	UserCompany:    "personal",
	// 	UserAddress:    a,
	// 	UserCity:       "ASDASDAS",
	// 	UserProvince:   "EGSERGESR",
	// 	UserCountry:    "ID",
	// 	UserPostalCode: "20020",
	// 	Category:       "",
	// }
	// api.RegisterDomain(tempp)
}
