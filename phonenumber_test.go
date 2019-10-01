package phonenumber

import (
	"testing"
)

func TestJP(t *testing.T) {
	number := Parse("090 6135 3368", "jp")
	expected := "819061353368"
	if number != expected {
		t.Errorf("JP: got %v\nwant %v", number, expected)
	}
}

func TestCNFormat(t *testing.T) {
	number := Parse("+8615948692360", "cn")
	expected := "8615948692360"
	if number != expected {
		t.Errorf("CN: got %v\nwant %v", number, expected)
	}
}

func TestUSAFormat(t *testing.T) {
	number := Parse("(817) 569-8900", "usa")
	expected := "18175698900"
	if number != expected {
		t.Errorf("CN: got %v\nwant %v", number, expected)
	}
}

func TestLVMobileFormat(t *testing.T) {
	number := Parse("+371 25 641 580", "lv")
	expected := "37125641580"
	if number != expected {
		t.Errorf("LV mobile: got %v\nwant %v", number, expected)
	}
}

func TestLVLandLineIsNotValid(t *testing.T) {
	number := Parse("+371 (67) 881-727", "lv")
	if number != "" {
		t.Errorf("Backward compatibility fail: Parse() for land line number must be empty")
	}
}

func TestLVLandLineFormat(t *testing.T) {
	number := ParseWithLandLine("+371 (67) 881-727", "lv")
	expected := "37167881727"
	if number != expected {
		t.Errorf("LV land line parse: got %v\nwant %v", number, expected)
	}
}

func TestLVLandLineFormatWithMobileNumberIsEmpty(t *testing.T) {
	number := ParseWithLandLine("+371(25)641580", "lv")
	expected := "37125641580"
	if number != expected {
		t.Errorf("LV land line parse: got %v\nwant %v", number, expected)
	}
}

func TestGetCountryForLVMobile(t *testing.T) {
	tv := "37125641580" // mobile number
	country := GetISO3166ByNumber(tv, true)
	expected := getISO3166ByCountry("lv")
	if country.CountryName == "" {
		t.Errorf("Country is empty for Latvian mobile number %s", tv)
	}
	if country.CountryName != expected.CountryName {
		t.Errorf("For Latvian mobile number %s got country %s\nwant %s", tv, country.CountryName, expected.CountryName)
	}
}

func TestGetCountryIsEmptyForLandLineNumberInMobileOnlySearch(t *testing.T) {
	tv := "37167881727" // land line number
	country := GetISO3166ByNumber(tv, false)
	if country.CountryName != "" {
		t.Errorf("Got country %s for landline number %s for lookup through mobile only numbers, expected: not found", country.CountryName, tv)
	}
}

func TestGetCountryForLVLandLine(t *testing.T) {
	tv := "37167881727" // land line number
	country := GetISO3166ByNumber(tv, true)
	expected := getISO3166ByCountry("lv")
	if country.CountryName == "" {
		t.Errorf("Country is empty for Latvian land line number %s", tv)
	}
	if country.CountryName != expected.CountryName {
		t.Errorf("For Latvian landline number %s got country %s\nwant %s", tv, country.CountryName, expected.CountryName)
	}
}

func TestGetCountryForITMobile(t *testing.T) {
	tv := "39335370971" // mobile number
	expected := getISO3166ByCountry("it") // country

	country := GetISO3166ByNumber(tv, true)
	if country.CountryName == "" {
		t.Errorf("Country is empty for %s mobile number `%s`", expected.CountryName, tv)
	}
	if country.CountryName != expected.CountryName {
		t.Errorf("For `%s` mobile number `%s` got country `%s`\nwant `%s`", expected.CountryName, tv, country.CountryName, expected.CountryName)
	}
}
