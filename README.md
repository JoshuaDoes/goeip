# goeip
Single file library for [Kealper](https://github.com/Kealper)'s geoip utility written in Golang without extra dependencies

# Installing
`go get github.com/JoshuaDoes/goeip`

# Example
```go
package main

import (
	"fmt"

	"github.com/JoshuaDoes/goeip"
)

func main() {
	result, err := goeip.Lookup("joshuadoes.com")
	if err != nil {
		fmt.Println("Error: " + fmt.Sprintf("%v", err))
		return
	}
	if result.Error > 0 {
		fmt.Println("Lookup error: " + result.Details)
	}

	fmt.Println("City: " + result.City)
	fmt.Println("State: " + result.State)
}
```
### Output

```
> go run main.go
City: Dallas
State: Texas
```