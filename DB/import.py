from elasticsearch import helpers, Elasticsearch
import csv

es = Elasticsearch()

if not es.indices.exists(index='flights'):
    print("Importing sample data into ES")
    with open('./flights.csv') as f:
        reader = csv.DictReader(f)
        helpers.bulk(es, reader, index='flights', doc_type='doc')
