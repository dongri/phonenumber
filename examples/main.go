package main

import (
	"fmt"

	"github.com/dongri/phonenumber"
)

func main() {
	countries := phonenumber.GetISO3166ByMobileNumber("07933846223")
	fmt.Println(countries[0].CountryName)

	countries = phonenumber.GetISO3166ByMobileNumber("09061353368")
	fmt.Println(countries[0].CountryName)

	countries = phonenumber.GetISO3166ByMobileNumber("14855512329")
	fmt.Println(countries[0].CountryName)

	parsed := phonenumber.ParseWithLandLine("+1 289 2999", "US")
	fmt.Println(parsed)

	pn1 := "+44 07700900000"
	fmt.Println(phonenumber.Parse(pn1, "GB"))

	pn2 := "07700900000"
	fmt.Println(phonenumber.Parse(pn2, "GB"))

	pn3 := "447700900000"
	fmt.Println(phonenumber.Parse(pn3, "GB"))

	pn4 := "+8109012345678"
	fmt.Println(phonenumber.Parse(pn4, "JP"))

}
