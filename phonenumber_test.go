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
	{"3197010280754", "nl"},
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

// Test the real and validated mobile number for India country
// We added "910" prefix that does not match a specification, but the numbers are really exists
var indiaMobileTests = []struct {
	input    string
}{
	// Land line numbers
	{"916361673045"},
	{"916354492726"},
	{"910024980635"},
	{"916392460280"},
	{"916354492726"},
	{"916350369833"},
	{"916367189828"},
	{"916350369833"},
	{"916201870408"},
	{"916201870408"},
	{"916377284213"},
	{"916397162236"},
	{"916200914995"},
	{"916381596784"},
	{"916377284213"},
	{"916283965522"},
	{"916354231595"},
	{"916260568572"},
	{"916395549904"},
	{"916354231595"},
	{"916268561529"},
	{"916268561529"},
	{"916268561529"},
	{"916268561529"},
	{"916382170172"},
	{"916370486858"},
	{"916304932224"},
	{"916206639955"},
	{"916370486858"},
	{"916397449357"},
	{"916263905286"},
	{"916382336575"},
	{"910000000024"},
	{"916264430744"},
	{"910000000022"},
	{"910022528998"},
	{"910020261021"},
	{"916376956621"},
	{"916376956621"},
	{"910022528998"},
	{"910022528998"},
	{"910020782681"},
	{"910022408610"},
	{"910022528998"},
	{"910027049906"},
	{"916297393807"},
	{"916297393807"},
	{"910020782681"},
	{"910020923964"},
	{"910024940077"},
	{"916201093536"},
	{"916390717591"},
	{"916265805658"},
	{"916201093536"},
	{"916265805658"},
	{"916371726938"},
	{"916283596832"},
	{"916371726938"},
	{"916283596832"},
	{"916371726938"},
	{"916387118360"},
	{"916387118360"},
	{"916377594419"},
	{"916283596832"},
	{"916387118360"},
	{"916377594419"},
	{"916295828064"},
	{"916295828064"},
	{"916377594419"},
	{"916377594419"},
	{"910027080756"},
	{"916295828064"},
	{"910027080756"},
	{"910027080756"},
	{"910383474849"},
	{"910023579095"},
	{"910024682947"},
	{"910027080756"},
	{"910029581805"},
	{"910024317903"},
	{"910024755680"},
	{"910023579095"},
	{"910021695580"},
	{"910025192442"},
	{"910026385462"},
	{"910029581805"},
	{"916355957699"},
	{"910022984851"},
	{"910029581805"},
	{"910024308531"},
	{"910027522826"},
	{"910024672083"},
	{"910029323709"},
	{"910028421589"},
	{"910029323709"},
	{"916264458980"},
	{"910029414657"},
	{"910029566885"},
	{"910024308531"},
}

func TestIndiaMobileNumber(t *testing.T) {
	for _, tt := range indiaMobileTests {
		country := GetISO3166ByNumber(tt.input, false)
		if country.CountryName != "India" {
			t.Errorf("GetISO3166ByNumber(number=`%s`, withLandline=false): expected `%s`, actual `%s`", tt.input, "India", country.CountryName)
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
	{"2349091500528", "ng"},
	{"5491138697327", "ar"},
	{"96871983009", "om"},
	{"23059402290", "mu"},
	{"387644523518", "ba"},
	{"380721753127", "ua"},
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

