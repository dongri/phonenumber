package phonenumber

import (
	"regexp"
	"strings"
)

// Parse ...
func Parse(number string, country string) string {
	number = strings.Replace(number, " ", "", -1)
	country = strings.Replace(country, " ", "", -1)
	plusSign := false
	if strings.HasPrefix(number, "+") {
		plusSign = true
	}

	// remove any non-digit character, included the +
	number = regexp.MustCompile(`\D`).ReplaceAllString(number, "")

	iso3166 := getISO3166ByCountry(country)

	if indexOfString(iso3166.Alpha3, []string{"GAB", "CIV", "COG"}) == -1 {
		r := regexp.MustCompile(`^0+`)
		number = r.ReplaceAllString(number, "")
	}
	r := regexp.MustCompile(`^89`)
	if iso3166.Alpha3 == "RUS" && len(number) == 11 && r.MatchString(number) == true {
		r := regexp.MustCompile(`^8+`)
		number = r.ReplaceAllString(number, "")
	}
	if plusSign {
		iso3166 = getISO3166ByNumber(number)
	} else {
		if indexOfInt(len(number), iso3166.PhoneNumberLengths) != -1 {
			number = iso3166.CountryCode + number
		}
	}
	if validatePhoneISO3166(number, iso3166) {
		return number
	}
	return ""
}

func getISO3166ByCountry(country string) ISO3166 {
	iso3166 := ISO3166{}
	uppperCaseCountry := strings.ToUpper(country)
	switch len(country) {
	case 0:
		iso3166 = GetISO3166()[0]
		break
	case 2:
		for _, i := range GetISO3166() {
			if i.Alpha2 == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
		break
	case 3:
		for _, i := range GetISO3166() {
			if i.Alpha3 == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
		break
	default:
		for _, i := range GetISO3166() {
			if strings.ToUpper(i.CountryName) == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
		break
	}
	return iso3166
}

func getISO3166ByNumber(number string) ISO3166 {
	iso3166 := ISO3166{}
	for _, i := range GetISO3166() {
		r := regexp.MustCompile(`^` + i.CountryCode)
		for _, l := range i.PhoneNumberLengths {
			if r.MatchString(number) && len(number) == len(i.CountryCode)+l {
				for _, w := range i.MobileBeginWith {
					r := regexp.MustCompile(`^` + i.CountryCode + w)
					if r.MatchString(number) {
						iso3166 = i
						break
					}
				}
			}
		}
	}
	return iso3166
}

func validatePhoneISO3166(number string, iso3166 ISO3166) bool {
	r := regexp.MustCompile(`^` + iso3166.CountryCode)
	number = r.ReplaceAllString(number, "")
	if len(iso3166.PhoneNumberLengths) == 0 {
		return false
	}
	for _, l := range iso3166.PhoneNumberLengths {
		if l == len(number) {
			for _, w := range iso3166.MobileBeginWith {
				r := regexp.MustCompile(`^` + w)
				if r.MatchString(number) == true {
					return true
				}
			}
		}
	}
	return false
}

func indexOfString(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func indexOfInt(word int, data []int) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
