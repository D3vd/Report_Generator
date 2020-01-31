package main

import (
	"fmt"
	"time"

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
		return false
	}

	e.client = client
	return true
}

// GetDocumentsWithCarrierAndTimeFrame : Query only with carrier name and timeframe
func (e *ES) GetDocumentsWithCarrierAndTimeFrame(carrierName string, start time.Time, end time.Time) {

	query := elastic.NewBoolQuery()

	query.Filter(elastic.NewMatchQuery("Carrier", carrierName))
	query.Filter(elastic.NewRangeQuery("timestamp").Format("strict_date_optional_time").From(start).To(end))

	res, err := e.client.Search().Index("flights").Type("doc").Query(query).Size(10000).Do()

	if err != nil {
		fmt.Println(err)
	}

	println(res.TotalHits())
}
