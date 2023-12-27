## MeowFacts App 

In this application I use the meowfacts API here is its [link](https://github.com/wh-iterabb-it/meowfacts)
#bold Example
```go
package main

import (
	"log"
	"time"

	"github.com/Sskrill/meowfacts_Sskrill/meowfacts"
)

func main() {
	cl, err := meowfacts.NewClient(time.Second * 5)
	if err != nil {
		log.Fatal(err)

	}
	err = cl.GetAssetOnLang("rus")
	if err != nil {

		log.Fatal(err)
	}

	err = cl.GetAsset(2)
	if err != nil {

		log.Fatal(err)
	}
}
```
