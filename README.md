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

	err = cl.GetAssetsOnLang(3, "rus")
	if err != nil {

		log.Fatal(err)
	}
}
```
**cl.GetAssetOnLang("rus")**  The language must be used as an argument in this method.Here are the supported languages. 
![langs](https://cdn.discordapp.com/attachments/592741750393536522/1189586373510967336/image.png?ex=659eb3a4&is=658c3ea4&hm=bbc81df124dede1bde61ff716bfd1fa74bcf3f902a341c1c21ba499d52591fbc&)

**cl.GetAsset(2)**  Here we have to insert how many facts we want to know

![resp](https://cdn.discordapp.com/attachments/592741750393536522/1189588069754617856/image.png?ex=659eb539&is=658c4039&hm=095039a90f08415fa9601939af6048e40472f33ffc24ffb5a95d020ae973419e&)

**cl.GetAssetsOnLang(3, "rus")** This method makes multiple requests in the specified language.(*But it sends a log before each request, because I have a RoundTripper with logging, maybe in the future I will change it and try to make the log once, but I will have to use DefaultClient instead of my Client structure*)

![resps](https://cdn.discordapp.com/attachments/592741750393536522/1189802856606335126/image.png?ex=659f7d42&is=658d0842&hm=3df2174696d8133559d85982f2264eee01b3f8351f6a96f19056ad730f784b9e&)
