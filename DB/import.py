from elasticsearch import helpers, Elasticsearch
import csv

es = Elasticsearch()

with open('./flights.csv') as f:
    reader = csv.DictReader(f)
    helpers.bulk(es, reader, index='flights', doc_type='_doc')
