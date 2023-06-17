package lib

import (
	"fmt"
	"net/http"
)

const (
	API_ROOT = "https://api-extern.systembolaget.se/site/V2"
	API_KEY  = "process.env.SYSTEMBOLAGET_API_KEY"
)

func SystembolagetApi(path string) {
	fullUrl := API_ROOT + path
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println(res.Body)
	// TODO: parse json
}
