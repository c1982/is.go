# is.go
Micro check library for Go inspired by [Is.js](https://github.com/arasatasaygin/is.js)


## Usage

```
go get github.com/c1982/is.go
```

```Go
package main


import "github.com/c1982/is.go"
import "fmt"

func main()  {
	
	isCap := Is.Capitalized("This Is Test Input")
	fmt.Println(isCap)
}
```

String Checks
===========

```Go
Is.Lowercase("lower case text")
-> true
```

```Go
Is.Uppercase("UPPER TEXT")
-> true
```

```Go
Is.EndWith("bloket.pro",".pro")
-> true
```

```Go
Is.StartWith("bloket.pro","bloket")
-> true
```

```Go
Is.Palindrome("makam")
-> true
```

```Go
Is.Include("Welcome to matrix","matrix")
-> true
```

Regex Checks
===========

```Go
Is.URL("http://www.bloket.pro")
-> true
```

```Go
Is.Email("info@spacex.com")
-> true
```

```Go
Is.IPv4("4.2.2.1")
-> true
```

```Go
Is.IPv6("2001:db8::1")
-> true
```

```Go
Is.CreditCard("378282246310005")
-> true
```

```Go
Is.UsZipCode("02201-1020")
-> true
```

```Go
Is.UkPostCode("B184BJ")
-> true
```