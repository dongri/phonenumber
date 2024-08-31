# phonenumber
![Coverage](https://img.shields.io/badge/Coverage-99.4%25-brightgreen)
[![Run Tests](https://github.com/dongri/phonenumber/actions/workflows/run-tests.yml/badge.svg)](https://github.com/dongri/phonenumber/actions?query=branch%3Amaster)
## Description

golang port of AfterShip/phone

phonenumber is used to normalize the mobile phone or landline number into a E.164 format.

The common problem is user normally input the phone number in this way:

```
09061354467 or
090-6135-4467 or
(090) 6135-4467 or
819061354467 or
+819061354467 or
81(90) 6135-4467 or
+81(90) 6135-4467 or ...
```

What we want is always:

```
819061354467
```

## Installation
Run the following command to install the package:
```
go get github.com/dongri/phonenumber
```

## Usage

### Clearing format
In this case landline numbers will be an invalid result:
```go
import "github.com/dongri/phonenumber"

number := phonenumber.Parse("090-6135-4467", "JP")
fmt.Println(number)
// Output: 819061354467
```

In this case you can format numbers included landline numbers:
```go
import "github.com/dongri/phonenumber"

number := phonenumber.ParseWithLandLine("+371 65 552-336", "LV")
fmt.Println(number)
// Output: 37165552336
```

### Get country for number
```go
import "github.com/dongri/phonenumber"

// Get country with mobile and landline numbers
// Let's try to get country for Latvian landline number
includeLandLine := true
country := phonenumber.GetISO3166ByNumber("37165552336", includeLandLine)
fmt.Println(country.CountryName)
// Output: Latvia

// Get country with mobile numbers only
includeLandLine = false
country := phonenumber.GetISO3166ByNumber("37165552336", includeLandLine)
fmt.Println(country.CountryName)
// Output:
```

## License
MIT
