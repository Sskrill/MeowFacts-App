## MeowFacts App 

In this application I use the meowfacts API here is its [link](https://github.com/wh-iterabb-it/meowfacts)

# *Example*
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
cl.GetAssetOnLang("rus")  The language must be used as an argument in this method.Here are the supported languages. 
![langs](https://cdn.discordapp.com/attachments/592741750393536522/1189586373510967336/image.png?ex=659eb3a4&is=658c3ea4&hm=bbc81df124dede1bde61ff716bfd1fa74bcf3f902a341c1c21ba499d52591fbc&)
