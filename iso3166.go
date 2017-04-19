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

// GetISO3166 ...
func GetISO3166() []ISO3166 {
	iso3166Datas := []ISO3166{}
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
	i.MobileBeginWith = []string{""}
	i.PhoneNumberLengths = []int{6, 7, 8, 9}
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
	i.PhoneNumberLengths = []int{10, 11, 12, 13, 14}
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

	i.Alpha2 = "KR"
	i.Alpha3 = "KOR"
	i.CountryCode = "82"
	i.CountryName = "Korea, Republic of"
	i.MobileBeginWith = []string{"1"}
	i.PhoneNumberLengths = []int{9, 10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "JP"
	i.Alpha3 = "JPN"
	i.CountryCode = "81"
	i.CountryName = "Japan"
	i.MobileBeginWith = []string{"70", "80", "90"}
	i.PhoneNumberLengths = []int{10}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "CN"
	i.Alpha3 = "CHN"
	i.CountryCode = "86"
	i.CountryName = "China"
	i.MobileBeginWith = []string{"13", "14", "15", "17", "18"}
	i.PhoneNumberLengths = []int{11}
	iso3166Datas = append(iso3166Datas, i)

	i.Alpha2 = "TW"
	i.Alpha3 = "TWN"
	i.CountryCode = "886"
	i.CountryName = "Taiwan, Province Of China"
	i.MobileBeginWith = []string{"9"}
	i.PhoneNumberLengths = []int{9}
	iso3166Datas = append(iso3166Datas, i)

	return iso3166Datas
}
