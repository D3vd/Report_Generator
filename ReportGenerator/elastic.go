package main

import (
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
func (e *ES) GetDocumentsWithCarrierAndTimeFrame(carrierName string, start time.Time, end time.Time) (hits []*elastic.SearchHit, ok bool) {

	// Create query
	query := elastic.NewBoolQuery()

	query.Filter(elastic.NewMatchQuery("Carrier", carrierName))
	query.Filter(elastic.NewRangeQuery("timestamp").
		Format("strict_date_optional_time").
		From(start).
		To(end))

	// Perform the query
	res, err := e.client.Search().
		Index("flights").
		Type("doc").
		Query(query).
		Size(10000).
		Do()

	if err != nil {
		return res.Hits.Hits, false
	}

	return res.Hits.Hits, true
}
