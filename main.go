package main

import (
	digitalregistra "myapp/digitalRegistra"
	"os"
)

func main() {
	api := digitalregistra.NewAPI(os.Getenv("DIGITAL_REGISTRA_RESELLER_URL"))

	api.DomainCheck("asjidjasds.com")
}
