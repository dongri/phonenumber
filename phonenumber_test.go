package phonenumber

import (
	"testing"
)

// Tests format mobile
var mobFormatTests = []struct {
	input    string
	country  string
	expected string
}{
	{"090 6135 3368", "jp", "819061353368"},
	{"+8615948692360", "cn", "8615948692360"},
	{"(817) 569-8900", "usa", "18175698900"},
	{"+371 25 641 580", "lv", "37125641580"},
}

func TestFormatMobile(t *testing.T) {
	for _, tt := range mobFormatTests {
		number := Parse(tt.input, tt.country)
		if number != tt.expected {
			t.Errorf("Parse(number=`%s`, country=`%s`): expected `%s`, actual `%s`", tt.input, tt.country, tt.expected, number)
		}
	}
}

// Negative tests for mobile format (landline numbers is not valid)
var mobFormatTestsNegative = []struct {
	input   string
	country string
}{
	// Land line numbers
	{"+371 (67) 881-727", "lv"},
	{"3726347343", "ee"},
	{"7499 709 88 33", "ru"},
	{"+48 22 (483) 53-34", "pl"},
	{"4970523743", "de"},
}

func TestFormatForLandLineIsEmpty(t *testing.T) {
	for _, tt := range mobFormatTestsNegative {
		number := Parse(tt.input, tt.country)
		if number != "" {
			t.Errorf("Parse(number=`%s`, country=`%s`) for land line number miust be empty, actual `%s`", tt.input, tt.country, number)
		}
	}
}

// Land line numbers format
var mobWithLLFormatTests = []struct {
	input    string
	country  string
	expected string
}{
	// Land line numbers
	{"+371 (67) 881-727", "lv", "37167881727"},
	{"00371 (67) 881-727", "lv", "37167881727"},
	{"3726347343", "ee", "3726347343"},
	{"7499 709 88 33", "ru", "74997098833"},
	//{"8499 709 88 33", "ru", "74997098833"}, fixme
	{"22 (483) 53-34", "pl", "48224835334"},
	// Mobile numbers
	{"090 6135 3368", "jp", "819061353368"},
	{"+8615948692360", "cn", "8615948692360"},
	{"008615948692360", "cn", "8615948692360"},
	{"(817) 569-8900", "usa", "18175698900"},
	{"+371 25 641 580", "lv", "37125641580"},
	{"00371 25 641 580", "lv", "37125641580"},
}

func TestFormatWithLandLine(t *testing.T) {
	for _, tt := range mobWithLLFormatTests {
		number := ParseWithLandLine(tt.input, tt.country)
		if number != tt.expected {
			t.Errorf("Parse(number=`%s`, country=`%s`): expected `%s`, actual `%s`", tt.input, tt.country, tt.expected, number)
		}
	}
}

// Get country by mobile number only
var mobWithLLCountryTests = []struct {
	input    string
	expected string
}{
	// Land line numbers
	{"3726347343", "ee"},
	{"74997098833", "ru"},
	{"37167881727", "lv"},
	// Mobile numbers
	{"39339638066", "it"},
	{"37125641580", "lv"},
	{"43663242739", "at"},
	{"21655886170", "tn"},
}

func TestGetCountryForMobileNumberWithLandLine(t *testing.T) {
	for _, tt := range mobWithLLCountryTests {
		country := GetISO3166ByNumber(tt.input, true)
		if tt.expected == "" {
			if country.CountryName != "" {
				t.Errorf("GetISO3166ByNumber(number=`%s`, withLandline=false): must be empty, actual `%s`", tt.input, country.CountryName)
			}
		} else {
			expected := getISO3166ByCountry(tt.expected)
			if country.CountryName != expected.CountryName {
				t.Errorf("GetISO3166ByNumber(number=`%s`, withLandline=false): expected `%s`, actual `%s`", tt.input, expected.CountryName, country.CountryName)
			}
		}
	}
}

// Get country by mobile number only
var mobCountryTests = []struct {
	input    string
	expected string
}{
	// Land line numbers
	{"3726347343", ""},
	{"74997098833", ""},
	{"37167881727", ""},
	// Mobile numbers
	{"39339638066", "it"},
	{"3933431022608", "it"},
	{"37125641580", "lv"},
	{"43663242739", "at"},
	{"21655886170", "tn"},
}

func TestGetCountryForMobileNumber(t *testing.T) {
	for _, tt := range mobCountryTests {
		country := GetISO3166ByNumber(tt.input, false)
		if tt.expected == "" {
			if country.CountryName != "" {
				t.Errorf("GetISO3166ByNumber(number=`%s`, withLandline=false): must be empty, actual `%s`", tt.input, country.CountryName)
			}
		} else {
			expected := getISO3166ByCountry(tt.expected)
			if country.CountryName != expected.CountryName {
				t.Errorf("GetISO3166ByNumber(number=`%s`, withLandline=false): expected `%s`, actual `%s`", tt.input, expected.CountryName, country.CountryName)
			}
		}
	}
}
