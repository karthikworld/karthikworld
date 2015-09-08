package main

import (
	"crypto/tls"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	url := "http://www.imdb.com/find?s=all&q=thanioruvan"
	client := &http.Client{Transport: tr}

	res, err := client.Get(url)

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	reg, _ := regexp.Compile(`tt[0-9]{7}`)

	//fmt.Printf("%s", string(b))

	data := reg.FindAllString(string(b), 1)

	//fmt.Printf("%s", data[0])

	url1 := ("http://www.imdb.com/title/" + data[0])

	res1, err := client.Get(url1)

	//fmt.Println(url)

	c, err := ioutil.ReadAll(res1.Body)

	defer res1.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(c))
}
