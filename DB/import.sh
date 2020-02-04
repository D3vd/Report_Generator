#!/bin/sh

/usr/share/elasticsearch/bin/elasticsearch -Des.insecure.allow.root=true > /dev/null &
sleep 10

python3 import.py

/usr/share/elasticsearch/bin/elasticsearch -Des.insecure.allow.root=true
