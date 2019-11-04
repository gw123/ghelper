package main

import (
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"log"
)

func main()  {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	data ,_:= ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}

