# is.go
Micro check library for Go inspired by [Is.js](https://github.com/arasatasaygin/is.js)


## Usage

```
go get github.com/c1982/is.go
```


String checks
===========

```Go
package main


import "github.com/c1982/is.go"
import "fmt"

func main()  {
	
	isLower := Is.Lowercase("Test")
	fmt.Println(isLower)
}
```