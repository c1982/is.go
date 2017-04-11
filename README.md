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

Checks
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