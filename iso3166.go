package phonenumber

// ISO3166 ...
type ISO3166 struct {
	Alpha2             string
	Alpha3             string
	CountryCode        string
	CountryName        string
	MobileBeginWith    []string
	PhoneNumberLengths []int
}

func init() {
	// Run once during package initialization in order to avoid data races
	// https://go.dev/doc/effective_go#init
	populateISO3166()
}

var iso3166Datas []ISO3166

// GetISO3166 returns the ISO3166 configuration for each country.
// Data are loaded during package initialization.
func GetISO3166() []ISO3166 {
	return iso3166Datas
}

// populateISO3166 contains the definitions of the per-country mobile number configuration.
// It operates on the iso3166Datas global variable and will return it after population.
func populateISO3166() {
	if iso3166Datas != nil {
		return
	}

	var i = ISO3166{}

	i.Alpha2 = "US"
	i.Alpha3 = "USA"
	i.CountryCode = "1"
	i.CountryName = "United States"
	i.MobileBeginWith = []string{
		"201", "202", "203", "205", "206", "207", "208", "209", "210", "212", "213", "214", "215",
		"216", "217", "218", "219", "224", "225", "227", "228", "229", "231", "234", "239", "240", "248", "251",
		"252", "253", "254", "256", "260", "262", "267", "269", "270", "272", "274", "276", "278", "281", "283",
		"301", "302", "303", "304", "305", "307", "308", "309", "310", "312", "313", "314", "315", "316", "317",
		"318", "319", "320", "321", "323", "325", "327", "330", "331", "334", "336", "337", "339", "341", "346",
		"347", "351", "352", "360", "361", "364", "369", "380", "385", "386", "401", "402", "404", "405", "406",
		"407", "408", "409", "410", "412", "413", "414", "415", "417", "419", "423", "424", "425", "430", "432",
		"434", "435", "440", "442", "443", "445", "447", "458", "464", "469", "470", "475", "478", "479", "480",
		"484", "501", "502", "503", "504", "505", "507", "508", "509", "510", "512", "513", "515", "516", "517",
		"518", "520", "530", "531", "534", "539", "540", "541", "551", "557", "559", "561", "562", "563", "564",
		"567", "570", "571", "573", "574", "575", "580", "582", "585", "586", "601", "602", "603", "605", "606",
		"607", "608", "609", "610", "612", "614", "615", "616", "617", "618", "619", "620", "623", "626", "627",
		"628", "630", "631", "636", "641", "646", "650", "651", "657", "659", "660", "661", "662", "667", "669",
		"678", "679", "681", "682", "689", "701", "702", "703", "704", "706", "707", "708", "712", "713", "714",
		"715", "716", "717", "718", "719", "720", "724", "725", "727", "730", "731", "732", "734", "737", "740",
		"747", "752", "754", "757", "760", "762", "763", "764", "765", "769", "770", "772", "773", "774", "775",
		"779", "781", "785", "786", "801", "802", "803", "804", "805", "806", "808", "810", "812", "813", "814",
		"815", "816", "817", "818", "828", "830", "831", "832", "835", "843", "845", "847", "848", "850", "856",
		"857", "858", "859", "860", "862", "863", "864", "865", "870", "872", "878", "901", "903", "904", "906",
		"907", "908", "909", "910", "912", "913", "914", "915", "916", "917", "918", "919", "920", "925", "927",
		"928", "929", "931", "935", "936", "937", "938", "940", "941", "947", "949", "951", "952", "954", "956",
		"957", "959", "970", "971", "972", "973", "975", "978", "979", "980", "984", "985", "989"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AW"
	i.Alpha3 = "ABW"
	i.CountryCode = "297"
	i.CountryName = "Aruba"
	i.MobileBeginWith = []string{"5", "6", "7", "9"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AF"
	i.Alpha3 = "AFG"
	i.CountryCode = "93"
	i.CountryName = "Afghanistan"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AO"
	i.Alpha3 = "AGO"
	i.CountryCode = "244"
	i.CountryName = "Angola"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AI"
	i.Alpha3 = "AIA"
	i.CountryCode = "1"
	i.CountryName = "Anguilla"
	i.MobileBeginWith = []string{"2645", "2647"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AX"
	i.Alpha3 = "ALA"
	i.CountryCode = "358"
	i.CountryName = "Åland Islands"
	i.MobileBeginWith = []string{"18"}
	i.PhoneNumberLengths = []int{6, 7, 8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AL"
	i.Alpha3 = "ALB"
	i.CountryCode = "355"
	i.CountryName = "Albania"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AD"
	i.Alpha3 = "AND"
	i.CountryCode = "376"
	i.CountryName = "Andorra"
	i.MobileBeginWith = []string{"3", "4", "6"}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AE"
	i.Alpha3 = "ARE"
	i.CountryCode = "971"
	i.CountryName = "United Arab Emirates"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AR"
	i.Alpha3 = "ARG"
	i.CountryCode = "54"
	i.CountryName = "Argentina"
	i.MobileBeginWith = []string{"6", "7", "8", "9"}
	i.PhoneNumberLengths = []int{10, 11, 12, 13, 14}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AM"
	i.Alpha3 = "ARM"
	i.CountryCode = "374"
	i.CountryName = "Armenia"
	i.MobileBeginWith = []string{"4", "5", "7", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AS"
	i.Alpha3 = "ASM"
	i.CountryCode = "1"
	i.CountryName = "American Samoa"
	i.MobileBeginWith = []string{"684733", "684258"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AG"
	i.Alpha3 = "ATG"
	i.CountryCode = "1"
	i.CountryName = "Antigua and Barbuda"
	i.MobileBeginWith = []string{"2687"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AU"
	i.Alpha3 = "AUS"
	i.CountryCode = "61"
	i.CountryName = "Australia"
	i.MobileBeginWith = []string{"4"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AT"
	i.Alpha3 = "AUT"
	i.CountryCode = "43"
	i.CountryName = "Austria"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{7, 8, 9, 10, 11, 12, 13}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "AZ"
	i.Alpha3 = "AZE"
	i.CountryCode = "994"
	i.CountryName = "Azerbaijan"
	i.MobileBeginWith = []string{"4", "5", "6", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BI"
	i.Alpha3 = "BDI"
	i.CountryCode = "257"
	i.CountryName = "Burundi"
	i.MobileBeginWith = []string{"7", "29"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BE"
	i.Alpha3 = "BEL"
	i.CountryCode = "32"
	i.CountryName = "Belgium"
	i.MobileBeginWith = []string{"4"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BJ"
	i.Alpha3 = "BEN"
	i.CountryCode = "229"
	i.CountryName = "Benin"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BF"
	i.Alpha3 = "BFA"
	i.CountryCode = "226"
	i.CountryName = "Burkina Faso"
	i.MobileBeginWith = []string{"6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BD"
	i.Alpha3 = "BGD"
	i.CountryCode = "880"
	i.CountryName = "Bangladesh"
	i.MobileBeginWith = []string{"1"}
	i.PhoneNumberLengths = []int{8, 9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BG"
	i.Alpha3 = "BGR"
	i.CountryCode = "359"
	i.CountryName = "Bulgaria"
	i.MobileBeginWith = []string{"87", "88", "89", "98", "99", "43"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BH"
	i.Alpha3 = "BHR"
	i.CountryCode = "973"
	i.CountryName = "Bahrain"
	i.MobileBeginWith = []string{"3", "6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BS"
	i.Alpha3 = "BHS"
	i.CountryCode = "1"
	i.CountryName = "Bahamas"
	i.MobileBeginWith = []string{"242"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BA"
	i.Alpha3 = "BIH"
	i.CountryCode = "387"
	i.CountryName = "Bosnia and Herzegovina"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BY"
	i.Alpha3 = "BLR"
	i.CountryCode = "375"
	i.CountryName = "Belarus"
	i.MobileBeginWith = []string{"25", "29", "33", "44"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BZ"
	i.Alpha3 = "BLZ"
	i.CountryCode = "501"
	i.CountryName = "Belize"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BM"
	i.Alpha3 = "BMU"
	i.CountryCode = "1"
	i.CountryName = "Bermuda"
	i.MobileBeginWith = []string{"4413", "4415", "4417"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BO"
	i.Alpha3 = "BOL"
	i.CountryCode = "591"
	i.CountryName = "Bolivia"
	i.MobileBeginWith = []string{"5", "6", "7"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BR"
	i.Alpha3 = "BRA"
	i.CountryCode = "55"
	i.CountryName = "Brazil"
	i.MobileBeginWith = []string{
		"119", "129", "139", "149", "159", "169", "179", "189", "199", "219", "229", "249", "279", "289", "31", "32",
		"34", "38", "41", "43", "44", "45", "47", "48", "51", "53", "54", "55", "61", "62", "65", "67", "68", "69",
		"71", "73", "74", "75", "77", "79", "81", "82", "83", "84", "85", "86", "91", "92", "95", "96", "98"}
	i.PhoneNumberLengths = []int{10, 11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BB"
	i.Alpha3 = "BRB"
	i.CountryCode = "1"
	i.CountryName = "Barbados"
	i.MobileBeginWith = []string{"246"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BN"
	i.Alpha3 = "BRN"
	i.CountryCode = "673"
	i.CountryName = "Brunei Darussalam"
	i.MobileBeginWith = []string{"7", "8"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BT"
	i.Alpha3 = "BTN"
	i.CountryCode = "975"
	i.CountryName = "Bhutan"
	i.MobileBeginWith = []string{"17"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "BW"
	i.Alpha3 = "BWA"
	i.CountryCode = "267"
	i.CountryName = "Botswana"
	i.MobileBeginWith = []string{"71", "72", "73", "74", "75", "76"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CF"
	i.Alpha3 = "CAF"
	i.CountryCode = "236"
	i.CountryName = "Central African Republic"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CA"
	i.Alpha3 = "CAN"
	i.CountryCode = "1"
	i.CountryName = "Canada"
	i.MobileBeginWith = []string{"204", "226", "236", "249", "250", "289", "306", "343", "365", "403", "416", "418", "431",
		"437", "438", "450", "506", "514", "519", "579", "581", "587", "600", "604", "613", "639", "647", "705",
		"709", "778", "780", "807", "819", "867", "873", "902", "905"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CH"
	i.Alpha3 = "CHE"
	i.CountryCode = "41"
	i.CountryName = "Switzerland"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CL"
	i.Alpha3 = "CHL"
	i.CountryCode = "56"
	i.CountryName = "Chile"
	i.MobileBeginWith = []string{"2", "3", "5", "6", "8", "9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CN"
	i.Alpha3 = "CHN"
	i.CountryCode = "86"
	i.CountryName = "China"
	i.MobileBeginWith = []string{"13", "14", "15", "16", "17", "18"}
	i.PhoneNumberLengths = []int{10, 11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CI"
	i.Alpha3 = "CIV"
	i.CountryCode = "225"
	i.CountryName = "Côte D'Ivoire"
	i.MobileBeginWith = []string{"0", "4", "5", "6", "7"}
	i.PhoneNumberLengths = []int{8, 9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CM"
	i.Alpha3 = "CMR"
	i.CountryCode = "237"
	i.CountryName = "Cameroon"
	i.MobileBeginWith = []string{"6", "7", "9"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CD"
	i.Alpha3 = "COD"
	i.CountryCode = "243"
	i.CountryName = "Congo, The Democratic Republic Of The"
	i.MobileBeginWith = []string{"8", "9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CG"
	i.Alpha3 = "COG"
	i.CountryCode = "242"
	i.CountryName = "Congo"
	i.MobileBeginWith = []string{"0"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CK"
	i.Alpha3 = "COK"
	i.CountryCode = "682"
	i.CountryName = "Cook Islands"
	i.MobileBeginWith = []string{"5", "7"}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CO"
	i.Alpha3 = "COL"
	i.CountryCode = "57"
	i.CountryName = "Colombia"
	i.MobileBeginWith = []string{"3"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KM"
	i.Alpha3 = "COM"
	i.CountryCode = "269"
	i.CountryName = "Comoros"
	i.MobileBeginWith = []string{"3", "76"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CV"
	i.Alpha3 = "CPV"
	i.CountryCode = "238"
	i.CountryName = "Cape Verde"
	i.MobileBeginWith = []string{"5", "9"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CR"
	i.Alpha3 = "CRI"
	i.CountryCode = "506"
	i.CountryName = "Costa Rica"
	i.MobileBeginWith = []string{"3", "5", "6", "7", "8"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CU"
	i.Alpha3 = "CUB"
	i.CountryCode = "53"
	i.CountryName = "Cuba"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KY"
	i.Alpha3 = "CYM"
	i.CountryCode = "1"
	i.CountryName = "Cayman Islands"
	i.MobileBeginWith = []string{"345"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CY"
	i.Alpha3 = "CYP"
	i.CountryCode = "357"
	i.CountryName = "Cyprus"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CZ"
	i.Alpha3 = "CZE"
	i.CountryCode = "420"
	i.CountryName = "Czech Republic"
	i.MobileBeginWith = []string{"6", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DE"
	i.Alpha3 = "DEU"
	i.CountryCode = "49"
	i.CountryName = "Germany"
	i.MobileBeginWith = []string{"15", "16", "17"}
	i.PhoneNumberLengths = []int{10, 11, 12, 13, 14}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DJ"
	i.Alpha3 = "DJI"
	i.CountryCode = "253"
	i.CountryName = "Djibouti"
	i.MobileBeginWith = []string{"77"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DM"
	i.Alpha3 = "DMA"
	i.CountryCode = "1"
	i.CountryName = "Dominica"
	i.MobileBeginWith = []string{"767"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DK"
	i.Alpha3 = "DNK"
	i.CountryCode = "45"
	i.CountryName = "Denmark"
	i.MobileBeginWith = []string{"2", "30", "31", "37", "40", "41", "42", "50", "51", "52", "53", "60", "61", "71", "81", "91", "92", "93"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DO"
	i.Alpha3 = "DOM"
	i.CountryCode = "1"
	i.CountryName = "Dominican Republic"
	i.MobileBeginWith = []string{"809", "829", "849"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "DZ"
	i.Alpha3 = "DZA"
	i.CountryCode = "213"
	i.CountryName = "Algeria"
	i.MobileBeginWith = []string{"5", "6", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "EC"
	i.Alpha3 = "ECU"
	i.CountryCode = "593"
	i.CountryName = "Ecuador"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "EG"
	i.Alpha3 = "EGY"
	i.CountryCode = "20"
	i.CountryName = "Egypt"
	i.MobileBeginWith = []string{"1"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ER"
	i.Alpha3 = "ERI"
	i.CountryCode = "291"
	i.CountryName = "Eritrea"
	i.MobileBeginWith = []string{"1", "7", "8"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ES"
	i.Alpha3 = "ESP"
	i.CountryCode = "34"
	i.CountryName = "Spain"
	i.MobileBeginWith = []string{"6", "7", "59"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "EE"
	i.Alpha3 = "EST"
	i.CountryCode = "372"
	i.CountryName = "Estonia"
	i.MobileBeginWith = []string{"5", "81", "82", "83", "84", "85", "86", "87", "89"}
	i.PhoneNumberLengths = []int{7, 8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ET"
	i.Alpha3 = "ETH"
	i.CountryCode = "251"
	i.CountryName = "Ethiopia"
	i.MobileBeginWith = []string{"9", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FI"
	i.Alpha3 = "FIN"
	i.CountryCode = "358"
	i.CountryName = "Finland"
	i.MobileBeginWith = []string{"4", "5"}
	i.PhoneNumberLengths = []int{9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FJ"
	i.Alpha3 = "FJI"
	i.CountryCode = "679"
	i.CountryName = "Fiji"
	i.MobileBeginWith = []string{"7", "9"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FK"
	i.Alpha3 = "FLK"
	i.CountryCode = "500"
	i.CountryName = "Falkland Islands (Malvinas)"
	i.MobileBeginWith = []string{"5", "6"}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FR"
	i.Alpha3 = "FRA"
	i.CountryCode = "33"
	i.CountryName = "France"
	i.MobileBeginWith = []string{"6", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FO"
	i.Alpha3 = "FRO"
	i.CountryCode = "298"
	i.CountryName = "Faroe Islands"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "FM"
	i.Alpha3 = "FSM"
	i.CountryCode = "691"
	i.CountryName = "Micronesia, Federated States Of"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GA"
	i.Alpha3 = "GAB"
	i.CountryCode = "241"
	i.CountryName = "Gabon"
	i.MobileBeginWith = []string{"05", "06", "07"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GB"
	i.Alpha3 = "GBR"
	i.CountryCode = "44"
	i.CountryName = "United Kingdom"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GE"
	i.Alpha3 = "GEO"
	i.CountryCode = "995"
	i.CountryName = "Georgia"
	i.MobileBeginWith = []string{"5", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GH"
	i.Alpha3 = "GHA"
	i.CountryCode = "233"
	i.CountryName = "Ghana"
	i.MobileBeginWith = []string{"2", "5"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GI"
	i.Alpha3 = "GIB"
	i.CountryCode = "350"
	i.CountryName = "Gibraltar"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GN"
	i.Alpha3 = "GIN"
	i.CountryCode = "224"
	i.CountryName = "Guinea"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GP"
	i.Alpha3 = "GLP"
	i.CountryCode = "590"
	i.CountryName = "Guadeloupe"
	i.MobileBeginWith = []string{"690"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GM"
	i.Alpha3 = "GMB"
	i.CountryCode = "220"
	i.CountryName = "Gambia"
	i.MobileBeginWith = []string{"7", "9"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GW"
	i.Alpha3 = "GNB"
	i.CountryCode = "245"
	i.CountryName = "Guinea-Bissau"
	i.MobileBeginWith = []string{"5", "6", "7"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GQ"
	i.Alpha3 = "GNQ"
	i.CountryCode = "240"
	i.CountryName = "Equatorial Guinea"
	i.MobileBeginWith = []string{"222", "551"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GR"
	i.Alpha3 = "GRC"
	i.CountryCode = "30"
	i.CountryName = "Greece"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GD"
	i.Alpha3 = "GRD"
	i.CountryCode = "1"
	i.CountryName = "Grenada"
	i.MobileBeginWith = []string{"473"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GL"
	i.Alpha3 = "GRL"
	i.CountryCode = "299"
	i.CountryName = "Greenland"
	i.MobileBeginWith = []string{"4", "5"}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GT"
	i.Alpha3 = "GTM"
	i.CountryCode = "502"
	i.CountryName = "Guatemala"
	i.MobileBeginWith = []string{"3", "4", "5"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GF"
	i.Alpha3 = "GUF"
	i.CountryCode = "594"
	i.CountryName = "French Guiana"
	i.MobileBeginWith = []string{"694"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GU"
	i.Alpha3 = "GUM"
	i.CountryCode = "1"
	i.CountryName = "Guam"
	i.MobileBeginWith = []string{"671"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "GY"
	i.Alpha3 = "GUY"
	i.CountryCode = "592"
	i.CountryName = "Guyana"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "HK"
	i.Alpha3 = "HKG"
	i.CountryCode = "852"
	i.CountryName = "Hong Kong"
	i.MobileBeginWith = []string{"5", "6", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "HN"
	i.Alpha3 = "HND"
	i.CountryCode = "504"
	i.CountryName = "Honduras"
	i.MobileBeginWith = []string{"3", "7", "8", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "HR"
	i.Alpha3 = "HRV"
	i.CountryCode = "385"
	i.CountryName = "Croatia"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "HT"
	i.Alpha3 = "HTI"
	i.CountryCode = "509"
	i.CountryName = "Haiti"
	i.MobileBeginWith = []string{"3", "4"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "HU"
	i.Alpha3 = "HUN"
	i.CountryCode = "36"
	i.CountryName = "Hungary"
	i.MobileBeginWith = []string{"20", "30", "50", "31", "71", "70"}
	i.PhoneNumberLengths = []int{9, 12}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ID"
	i.Alpha3 = "IDN"
	i.CountryCode = "62"
	i.CountryName = "Indonesia"
	i.MobileBeginWith = []string{"8"}
	i.PhoneNumberLengths = []int{9, 10, 11, 12}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IN"
	i.Alpha3 = "IND"
	i.CountryCode = "91"
	i.CountryName = "India"
	i.MobileBeginWith = []string{"0", "6", "7", "8", "9"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IE"
	i.Alpha3 = "IRL"
	i.CountryCode = "353"
	i.CountryName = "Ireland"
	i.MobileBeginWith = []string{"82", "83", "84", "85", "86", "87", "88", "89"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IR"
	i.Alpha3 = "IRN"
	i.CountryCode = "98"
	i.CountryName = "Iran, Islamic Republic Of"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IQ"
	i.Alpha3 = "IRQ"
	i.CountryCode = "964"
	i.CountryName = "Iraq"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IS"
	i.Alpha3 = "ISL"
	i.CountryCode = "354"
	i.CountryName = "Iceland"
	i.MobileBeginWith = []string{"6", "7", "8"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IL"
	i.Alpha3 = "ISR"
	i.CountryCode = "972"
	i.CountryName = "Israel"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "IT"
	i.Alpha3 = "ITA"
	i.CountryCode = "39"
	i.CountryName = "Italy"
	i.MobileBeginWith = []string{"3"}
	i.PhoneNumberLengths = []int{9, 10, 11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "JM"
	i.Alpha3 = "JAM"
	i.CountryCode = "1"
	i.CountryName = "Jamaica"
	i.MobileBeginWith = []string{"876"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "JO"
	i.Alpha3 = "JOR"
	i.CountryCode = "962"
	i.CountryName = "Jordan"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "JP"
	i.Alpha3 = "JPN"
	i.CountryCode = "81"
	i.CountryName = "Japan"
	i.MobileBeginWith = []string{"70", "80", "90"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KZ"
	i.Alpha3 = "KAZ"
	i.CountryCode = "7"
	i.CountryName = "Kazakhstan"
	i.MobileBeginWith = []string{"70", "77", "747"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KE"
	i.Alpha3 = "KEN"
	i.CountryCode = "254"
	i.CountryName = "Kenya"
	i.MobileBeginWith = []string{"7", "1"} // https://ca.go.ke/new-series-of-mobile-number-introduced-in-kenya/
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KG"
	i.Alpha3 = "KGZ"
	i.CountryCode = "996"
	i.CountryName = "Kyrgyzstan"
	i.MobileBeginWith = []string{"5", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KH"
	i.Alpha3 = "KHM"
	i.CountryCode = "855"
	i.CountryName = "Cambodia"
	i.MobileBeginWith = []string{"1", "6", "7", "8", "9"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KI"
	i.Alpha3 = "KIR"
	i.CountryCode = "686"
	i.CountryName = "Kiribati"
	i.MobileBeginWith = []string{"9", "30"}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KN"
	i.Alpha3 = "KNA"
	i.CountryCode = "1"
	i.CountryName = "Saint Kitts And Nevis"
	i.MobileBeginWith = []string{"869"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KR"
	i.Alpha3 = "KOR"
	i.CountryCode = "82"
	i.CountryName = "Korea, Republic of"
	i.MobileBeginWith = []string{"1"}
	i.PhoneNumberLengths = []int{9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "KW"
	i.Alpha3 = "KWT"
	i.CountryCode = "965"
	i.CountryName = "Kuwait"
	i.MobileBeginWith = []string{"5", "6", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LA"
	i.Alpha3 = "LAO"
	i.CountryCode = "856"
	i.CountryName = "Lao People's Democratic Republic"
	i.MobileBeginWith = []string{"20"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LB"
	i.Alpha3 = "LBN"
	i.CountryCode = "961"
	i.CountryName = "Lebanon"
	i.MobileBeginWith = []string{"3", "7"}
	i.PhoneNumberLengths = []int{7, 8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LR"
	i.Alpha3 = "LBR"
	i.CountryCode = "231"
	i.CountryName = "Liberia"
	i.MobileBeginWith = []string{"4", "5", "6", "7"}
	i.PhoneNumberLengths = []int{7, 8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LY"
	i.Alpha3 = "LBY"
	i.CountryCode = "218"
	i.CountryName = "Libyan Arab Jamahiriya"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LC"
	i.Alpha3 = "LCA"
	i.CountryCode = "1"
	i.CountryName = "Saint Lucia"
	i.MobileBeginWith = []string{"758"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LI"
	i.Alpha3 = "LIE"
	i.CountryCode = "423"
	i.CountryName = "Liechtenstein"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LK"
	i.Alpha3 = "LKA"
	i.CountryCode = "94"
	i.CountryName = "Sri Lanka"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LS"
	i.Alpha3 = "LSO"
	i.CountryCode = "266"
	i.CountryName = "Lesotho"
	i.MobileBeginWith = []string{"5", "6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LT"
	i.Alpha3 = "LTU"
	i.CountryCode = "370"
	i.CountryName = "Lithuania"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LU"
	i.Alpha3 = "LUX"
	i.CountryCode = "352"
	i.CountryName = "Luxembourg"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "LV"
	i.Alpha3 = "LVA"
	i.CountryCode = "371"
	i.CountryName = "Latvia"
	i.MobileBeginWith = []string{"2"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MO"
	i.Alpha3 = "MAC"
	i.CountryCode = "853"
	i.CountryName = "Macao"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MA"
	i.Alpha3 = "MAR"
	i.CountryCode = "212"
	i.CountryName = "Morocco"
	i.MobileBeginWith = []string{"526", "527", "533", "534", "54", "55", "6", "7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MC"
	i.Alpha3 = "MCO"
	i.CountryCode = "377"
	i.CountryName = "Monaco"
	i.MobileBeginWith = []string{"4", "6"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MD"
	i.Alpha3 = "MDA"
	i.CountryCode = "373"
	i.CountryName = "Moldova, Republic of"
	i.MobileBeginWith = []string{"6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MG"
	i.Alpha3 = "MDG"
	i.CountryCode = "261"
	i.CountryName = "Madagascar"
	i.MobileBeginWith = []string{"3"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MV"
	i.Alpha3 = "MDV"
	i.CountryCode = "960"
	i.CountryName = "Maldives"
	i.MobileBeginWith = []string{"7", "9"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MX"
	i.Alpha3 = "MEX"
	i.CountryCode = "52"
	i.CountryName = "Mexico"
	i.MobileBeginWith = []string{""}
	i.PhoneNumberLengths = []int{10, 11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MH"
	i.Alpha3 = "MHL"
	i.CountryCode = "692"
	i.CountryName = "Marshall Islands"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MK"
	i.Alpha3 = "MKD"
	i.CountryCode = "389"
	i.CountryName = "Macedonia, the Former Yugoslav Republic Of"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ML"
	i.Alpha3 = "MLI"
	i.CountryCode = "223"
	i.CountryName = "Mali"
	i.MobileBeginWith = []string{"6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MT"
	i.Alpha3 = "MLT"
	i.CountryCode = "356"
	i.CountryName = "Malta"
	i.MobileBeginWith = []string{"79", "99", "98", "72", "92", "77", "96"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MM"
	i.Alpha3 = "MMR"
	i.CountryCode = "95"
	i.CountryName = "Myanmar"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ME"
	i.Alpha3 = "MNE"
	i.CountryCode = "382"
	i.CountryName = "Montenegro"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MN"
	i.Alpha3 = "MNG"
	i.CountryCode = "976"
	i.CountryName = "Mongolia"
	i.MobileBeginWith = []string{"5", "8", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MP"
	i.Alpha3 = "MNP"
	i.CountryCode = "1"
	i.CountryName = "Northern Mariana Islands"
	i.MobileBeginWith = []string{"670"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MZ"
	i.Alpha3 = "MOZ"
	i.CountryCode = "258"
	i.CountryName = "Mozambique"
	i.MobileBeginWith = []string{"8"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MR"
	i.Alpha3 = "MRT"
	i.CountryCode = "222"
	i.CountryName = "Mauritania"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MS"
	i.Alpha3 = "MSR"
	i.CountryCode = "1"
	i.CountryName = "Montserrat"
	i.MobileBeginWith = []string{"664"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MQ"
	i.Alpha3 = "MTQ"
	i.CountryCode = "596"
	i.CountryName = "Martinique"
	i.MobileBeginWith = []string{"696"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MU"
	i.Alpha3 = "MUS"
	i.CountryCode = "230"
	i.CountryName = "Mauritius"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{7, 8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MW"
	i.Alpha3 = "MWI"
	i.CountryCode = "265"
	i.CountryName = "Malawi"
	i.MobileBeginWith = []string{"77", "88", "99"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "MY"
	i.Alpha3 = "MYS"
	i.CountryCode = "60"
	i.CountryName = "Malaysia"
	i.MobileBeginWith = []string{"1"}
	i.PhoneNumberLengths = []int{9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "YT"
	i.Alpha3 = "MYT"
	i.CountryCode = "269"
	i.CountryName = "Mayotte"
	i.MobileBeginWith = []string{"639"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NA"
	i.Alpha3 = "NAM"
	i.CountryCode = "264"
	i.CountryName = "Namibia"
	i.MobileBeginWith = []string{"60", "81", "82", "85"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NC"
	i.Alpha3 = "NCL"
	i.CountryCode = "687"
	i.CountryName = "New Caledonia"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NE"
	i.Alpha3 = "NER"
	i.CountryCode = "227"
	i.CountryName = "Niger"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NF"
	i.Alpha3 = "NFK"
	i.CountryCode = "672"
	i.CountryName = "Norfolk Island"
	i.MobileBeginWith = []string{"5", "8"}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NG"
	i.Alpha3 = "NGA"
	i.CountryCode = "234"
	i.CountryName = "Nigeria"
	i.MobileBeginWith = []string{"70", "80", "81", "9"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NI"
	i.Alpha3 = "NIC"
	i.CountryCode = "505"
	i.CountryName = "Nicaragua"
	i.MobileBeginWith = []string{"8"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NU"
	i.Alpha3 = "NIU"
	i.CountryCode = "683"
	i.CountryName = "Niue"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{4}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NL"
	i.Alpha3 = "NLD"
	i.CountryCode = "31"
	i.CountryName = "Netherlands"
	i.MobileBeginWith = []string{"6", "97"}
	i.PhoneNumberLengths = []int{9, 11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NO"
	i.Alpha3 = "NOR"
	i.CountryCode = "47"
	i.CountryName = "Norway"
	i.MobileBeginWith = []string{"4", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NP"
	i.Alpha3 = "NPL"
	i.CountryCode = "977"
	i.CountryName = "Nepal"
	i.MobileBeginWith = []string{"97", "98"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NR"
	i.Alpha3 = "NRU"
	i.CountryCode = "674"
	i.CountryName = "Nauru"
	i.MobileBeginWith = []string{"555"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "NZ"
	i.Alpha3 = "NZL"
	i.CountryCode = "64"
	i.CountryName = "New Zealand"
	i.MobileBeginWith = []string{"2"}
	i.PhoneNumberLengths = []int{8, 9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "OM"
	i.Alpha3 = "OMN"
	i.CountryCode = "968"
	i.CountryName = "Oman"
	i.MobileBeginWith = []string{"7", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PK"
	i.Alpha3 = "PAK"
	i.CountryCode = "92"
	i.CountryName = "Pakistan"
	i.MobileBeginWith = []string{"3"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PA"
	i.Alpha3 = "PAN"
	i.CountryCode = "507"
	i.CountryName = "Panama"
	i.MobileBeginWith = []string{"5", "6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PE"
	i.Alpha3 = "PER"
	i.CountryCode = "51"
	i.CountryName = "Peru"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PH"
	i.Alpha3 = "PHL"
	i.CountryCode = "63"
	i.CountryName = "Philippines"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PW"
	i.Alpha3 = "PLW"
	i.CountryCode = "680"
	i.CountryName = "Palau"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PG"
	i.Alpha3 = "PNG"
	i.CountryCode = "675"
	i.CountryName = "Papua New Guinea"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PL"
	i.Alpha3 = "POL"
	i.CountryCode = "48"
	i.CountryName = "Poland"
	i.MobileBeginWith = []string{"45", "5", "6", "7", "8"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PR"
	i.Alpha3 = "PRI"
	i.CountryCode = "1"
	i.CountryName = "Puerto Rico"
	i.MobileBeginWith = []string{"787", "939"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PT"
	i.Alpha3 = "PRT"
	i.CountryCode = "351"
	i.CountryName = "Portugal"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PY"
	i.Alpha3 = "PRY"
	i.CountryCode = "595"
	i.CountryName = "Paraguay"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PS"
	i.Alpha3 = "PSE"
	i.CountryCode = "970"
	i.CountryName = "Palestinian Territory, Occupied"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PF"
	i.Alpha3 = "PYF"
	i.CountryCode = "689"
	i.CountryName = "French Polynesia"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "QA"
	i.Alpha3 = "QAT"
	i.CountryCode = "974"
	i.CountryName = "Qatar"
	i.MobileBeginWith = []string{"3", "5", "6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "RE"
	i.Alpha3 = "REU"
	i.CountryCode = "262"
	i.CountryName = "Réunion"
	i.MobileBeginWith = []string{"692", "693"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "RO"
	i.Alpha3 = "ROU"
	i.CountryCode = "40"
	i.CountryName = "Romania"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "RU"
	i.Alpha3 = "RUS"
	i.CountryCode = "7"
	i.CountryName = "Russian Federation"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "RW"
	i.Alpha3 = "RWA"
	i.CountryCode = "250"
	i.CountryName = "Rwanda"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SA"
	i.Alpha3 = "SAU"
	i.CountryCode = "966"
	i.CountryName = "Saudi Arabia"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SD"
	i.Alpha3 = "SDN"
	i.CountryCode = "249"
	i.CountryName = "Sudan"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SN"
	i.Alpha3 = "SEN"
	i.CountryCode = "221"
	i.CountryName = "Senegal"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SG"
	i.Alpha3 = "SGP"
	i.CountryCode = "65"
	i.CountryName = "Singapore"
	i.MobileBeginWith = []string{"8", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SH"
	i.Alpha3 = "SHN"
	i.CountryCode = "290"
	i.CountryName = "Saint Helena"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{4}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SJ"
	i.Alpha3 = "SJM"
	i.CountryCode = "47"
	i.CountryName = "Svalbard And Jan Mayen"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SB"
	i.Alpha3 = "SLB"
	i.CountryCode = "677"
	i.CountryName = "Solomon Islands"
	i.MobileBeginWith = []string{"7", "8"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SL"
	i.Alpha3 = "SLE"
	i.CountryCode = "232"
	i.CountryName = "Sierra Leone"
	i.MobileBeginWith = []string{"21", "25", "30", "33", "34", "40", "44", "50", "55", "76", "77", "78", "79", "88"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SV"
	i.Alpha3 = "SLV"
	i.CountryCode = "503"
	i.CountryName = "El Salvador"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SM"
	i.Alpha3 = "SMR"
	i.CountryCode = "378"
	i.CountryName = "San Marino"
	i.MobileBeginWith = []string{"3", "6"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SO"
	i.Alpha3 = "SOM"
	i.CountryCode = "252"
	i.CountryName = "Somalia"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SX"
	i.Alpha3 = "SXM"
	i.CountryCode = "1"
	i.CountryName = "Sint Maarten"
	i.MobileBeginWith = []string{"721"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "PM"
	i.Alpha3 = "SPM"
	i.CountryCode = "508"
	i.CountryName = "Saint Pierre And Miquelon"
	i.MobileBeginWith = []string{"55"}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "RS"
	i.Alpha3 = "SRB"
	i.CountryCode = "381"
	i.CountryName = "Serbia"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8, 9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ST"
	i.Alpha3 = "STP"
	i.CountryCode = "239"
	i.CountryName = "Sao Tome and Principe"
	i.MobileBeginWith = []string{"98", "99"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SR"
	i.Alpha3 = "SUR"
	i.CountryCode = "597"
	i.CountryName = "Suriname"
	i.MobileBeginWith = []string{"6", "7", "8"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SK"
	i.Alpha3 = "SVK"
	i.CountryCode = "421"
	i.CountryName = "Slovakia"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SI"
	i.Alpha3 = "SVN"
	i.CountryCode = "386"
	i.CountryName = "Slovenia"
	i.MobileBeginWith = []string{"3", "4", "5", "6", "7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SE"
	i.Alpha3 = "SWE"
	i.CountryCode = "46"
	i.CountryName = "Sweden"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SC"
	i.Alpha3 = "SYC"
	i.CountryCode = "248"
	i.CountryName = "Seychelles"
	i.MobileBeginWith = []string{"2"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SZ"
	i.Alpha3 = "SWZ"
	i.CountryCode = "268"
	i.CountryName = "Swaziland"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "SY"
	i.Alpha3 = "SYR"
	i.CountryCode = "963"
	i.CountryName = "Syrian Arab Republic"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TC"
	i.Alpha3 = "TCA"
	i.CountryCode = "1"
	i.CountryName = "Turks and Caicos Islands"
	i.MobileBeginWith = []string{"6492", "6493", "6494"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TD"
	i.Alpha3 = "TCD"
	i.CountryCode = "235"
	i.CountryName = "Chad"
	i.MobileBeginWith = []string{"6", "7", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TG"
	i.Alpha3 = "TGO"
	i.CountryCode = "228"
	i.CountryName = "Togo"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TH"
	i.Alpha3 = "THA"
	i.CountryCode = "66"
	i.CountryName = "Thailand"
	i.MobileBeginWith = []string{"6", "8", "9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TJ"
	i.Alpha3 = "TJK"
	i.CountryCode = "992"
	i.CountryName = "Tajikistan"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TK"
	i.Alpha3 = "TKL"
	i.CountryCode = "690"
	i.CountryName = "Tokelau"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{4}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TM"
	i.Alpha3 = "TKM"
	i.CountryCode = "993"
	i.CountryName = "Turkmenistan"
	i.MobileBeginWith = []string{"6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TL"
	i.Alpha3 = "TLS"
	i.CountryCode = "670"
	i.CountryName = "Timor-Leste"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TO"
	i.Alpha3 = "TON"
	i.CountryCode = "676"
	i.CountryName = "Tonga"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TT"
	i.Alpha3 = "TTO"
	i.CountryCode = "1"
	i.CountryName = "Trinidad and Tobago"
	i.MobileBeginWith = []string{"868"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TN"
	i.Alpha3 = "TUN"
	i.CountryCode = "216"
	i.CountryName = "Tunisia"
	i.MobileBeginWith = []string{"2", "4", "5", "9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TR"
	i.Alpha3 = "TUR"
	i.CountryCode = "90"
	i.CountryName = "Turkey"
	i.MobileBeginWith = []string{"5"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TV"
	i.Alpha3 = "TUV"
	i.CountryCode = "688"
	i.CountryName = "Tuvalu"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{5}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TW"
	i.Alpha3 = "TWN"
	i.CountryCode = "886"
	i.CountryName = "Taiwan, Province Of China"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TZ"
	i.Alpha3 = "TZA"
	i.CountryCode = "255"
	i.CountryName = "Tanzania, United Republic of"
	i.MobileBeginWith = []string{"7", "6"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "UG"
	i.Alpha3 = "UGA"
	i.CountryCode = "256"
	i.CountryName = "Uganda"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "UA"
	i.Alpha3 = "UKR"
	i.CountryCode = "380"
	i.CountryName = "Ukraine"
	i.MobileBeginWith = []string{"39", "50", "63", "66", "67", "68", "71", "72", "73", "9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "UY"
	i.Alpha3 = "URY"
	i.CountryCode = "598"
	i.CountryName = "Uruguay"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "UZ"
	i.Alpha3 = "UZB"
	i.CountryCode = "998"
	i.CountryName = "Uzbekistan"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VC"
	i.Alpha3 = "VCT"
	i.CountryCode = "1"
	i.CountryName = "Saint Vincent And The Grenedines"
	i.MobileBeginWith = []string{"784"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VE"
	i.Alpha3 = "VEN"
	i.CountryCode = "58"
	i.CountryName = "Venezuela, Bolivarian Republic of"
	i.MobileBeginWith = []string{"4"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VG"
	i.Alpha3 = "VGB"
	i.CountryCode = "1"
	i.CountryName = "Virgin Islands, British"
	i.MobileBeginWith = []string{"284"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VI"
	i.Alpha3 = "VIR"
	i.CountryCode = "1"
	i.CountryName = "Virgin Islands, U.S."
	i.MobileBeginWith = []string{"340"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VN"
	i.Alpha3 = "VNM"
	i.CountryCode = "84"
	i.CountryName = "Viet Nam"
	i.MobileBeginWith = []string{"9", "1"}
	i.PhoneNumberLengths = []int{9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "VU"
	i.Alpha3 = "VUT"
	i.CountryCode = "678"
	i.CountryName = "Vanuatu"
	i.MobileBeginWith = []string{"5", "7"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "WF"
	i.Alpha3 = "WLF"
	i.CountryCode = "681"
	i.CountryName = "Wallis and Futuna"
	i.MobileBeginWith = []string{}
	i.PhoneNumberLengths = []int{6}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "WS"
	i.Alpha3 = "WSM"
	i.CountryCode = "685"
	i.CountryName = "Samoa"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{7}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "XK"
	i.Alpha3 = "XKX"
	i.CountryCode = "383"
	i.CountryName = "Kosovo"
	i.MobileBeginWith = []string{"3", "4", "6"}
	i.PhoneNumberLengths = []int{8}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "YE"
	i.Alpha3 = "YEM"
	i.CountryCode = "967"
	i.CountryName = "Yemen"
	i.MobileBeginWith = []string{"7"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ZA"
	i.Alpha3 = "ZAF"
	i.CountryCode = "27"
	i.CountryName = "South Africa"
	i.MobileBeginWith = []string{"6", "7", "8"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ZM"
	i.Alpha3 = "ZMB"
	i.CountryCode = "260"
	i.CountryName = "Zambia"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "ZW"
	i.Alpha3 = "ZWE"
	i.CountryCode = "263"
	i.CountryName = "Zimbabwe"
	i.MobileBeginWith = []string{"71", "73", "77"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)
}
