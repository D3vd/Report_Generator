package main

import (
	"fmt"

	"gopkg.in/olivere/elastic.v3"
)

// ES : Elasticsearch structure
type ES struct {
	client *elastic.Client
}

// Init connection to Elasticsearch
func (e *ES) Init(port string) (ok bool) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + port))

	if err != nil {
		fmt.Println(err)
		return false
	}

	e.client = client
	return true
}
