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
	{"090 6135 3368", "JP", "819061353368"},
	{"+8615948692360", "CN", "8615948692360"},
	{"(817) 569-8900", "USA", "18175698900"},
	{"+371 25 641 580", "LV", "37125641580"},
	{"+62 895349866066", "ID", "62895349866066"},
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
	{"+371 (67) 881-727", "LV"},
	{"3726347343", "EE"},
	{"7499 709 88 33", "RU"},
	{"+48 22 (483) 53-34", "PL"},
	{"4970523743", "DE"},
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
	{"+371 (67) 881-727", "LV", "37167881727"},
	{"00371 (67) 881-727", "LV", "37167881727"},
	// {"3726347343", "EE", "3726347343"}, FIXME: expected `372634734312`, actual `372372634734312`
	{"7499 709 88 33", "RU", "74997098833"},
	//{"8499 709 88 33", "RU", "74997098833"}, fixme
	{"22 (483) 53-34", "PL", "48224835334"},
	// Mobile numbers
	{"090 6135 3368", "JP", "819061353368"},
	{"+8615948692360", "CN", "8615948692360"},
	{"008615948692360", "CN", "8615948692360"},
	{"(817) 569-8900", "USA", "18175698900"},
	{"+371 25 641 580", "LV", "37125641580"},
	{"00371 25 641 580", "LV", "37125641580"},
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
	{"3726347343", "EE"},
	{"74997098833", "RU"},
	{"37167881727", "LV"},
	// Mobile numbers
	{"39339638066", "IT"},
	{"37125641580", "LV"},
	{"43663242739", "AT"},
	{"21655886170", "TN"},
	{"3197010280754", "NL"},
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
	input string
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
	number          string
	expectedCountry string
	expectedIsMobile bool
}{
	{"3726347343", "EE", false},
	{"74997098833", "RU", false},
	{"37167881727", "LV", false},
	{"39339638066", "IT", true},
	{"3933431022608", "IT", true},
	{"37125641580", "LV", true},
	{"43663242739", "AT", true},
	{"21655886170", "TN", true},
	{"2349091500528", "NG", true},
	{"5491138697327", "AR", true},
	{"96871983009", "OM", true},
	{"23059402290", "MU", true},
	{"387644523518", "BA", true},
	{"380721753127", "UA", true},
	{"380713707383", "UA", true},
	{"48486565565", "PL", false},
	{"48453234651", "PL", true},
}

func TestGetCountryForMobileNumber(t *testing.T) {
	for _, tt := range mobCountryTests {

		// Check number has correct type
		countryWithLandline := GetISO3166ByNumber(tt.number, true)
		countryMobileOnly := GetISO3166ByNumber(tt.number, false)

		if tt.expectedIsMobile == true {
			// Mobile type is included in the subset of landline numbers
			if countryMobileOnly.CountryName == ""  {
				t.Errorf("number type is incorrect [number=%s, actual=landline, expected=mobile]", tt.number)
			} else {
				if countryMobileOnly.Alpha2 != tt.expectedCountry {
					t.Errorf("country for number is invalid [number=%s, actual=%s, expected=%s]", tt.number, countryMobileOnly.Alpha2, tt.expectedCountry)
				}
			}

		}
		if tt.expectedIsMobile == false {
			// Landline type is excluded from the subset of mobile number rules
			if countryWithLandline.CountryName == "" || countryMobileOnly.CountryName != "" {
				t.Errorf("number type is incorrect [number=%s, actual=mobile, expected=landline]", tt.number)
			} else {
				if countryWithLandline.Alpha2 != tt.expectedCountry {
					t.Errorf("country for number is invalid [number=%s, actual=%s, expected=%s]", tt.number, countryWithLandline.Alpha2, tt.expectedCountry)
				}
			}
		}

	}
}
