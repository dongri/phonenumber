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
}
