# phonenumber

golang port of node-phone

phonenumber is used to normalize the mobile phone number into a E.164 format.

The common problem is user normally input the phone number in this way:

```
09061354467 or
090-6135-4467 or
(090) 6135-4467 or
8109061354467 or
+8109061354467 or
81(090) 6135-4467 or
+81(090) 6135-4467 or ...
```

What we want is always:

```
819061354467
```

# Install
```
go get github.com/dongri/phonenumber
```

# Usage

```
import "github.com/dongri/phonenumber"
number := phonenumber.Parse("090-6135-4467", "JP")
fmt.Println(number)
```
