FROM ubuntu:latest

RUN apt-get update

# Install Java and Python
RUN apt-get install -y wget
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk
RUN apt-get install -y python3
RUN apt-get install -y python3-pip

#  Download and run elasticsearch
RUN wget https://download.elastic.co/elasticsearch/release/org/elasticsearch/distribution/tar/elasticsearch/2.4.0/elasticsearch-2.4.0.tar.gz
RUN tar -xvf elasticsearch-2.4.0.tar.gz
RUN cd elasticsearch-2.4.0
RUN bin/elasticsearch -Des.insecure.allow.root=true > /dev/null 2>&1 &
RUN sleep 20

# Expose ports for elasticsearch
EXPOSE 9200
EXPOSE 9300

# Copy the csv and import file to the root directort
WORKDIR /
COPY . .

# Install elasticsearch and run import.py
RUN pip3 install elasticsearch
RUN python3 import.py

CMD [ "/elasticsearch-2.4.0/bin/elasticsearch", "-Des.insecure.allow.root=true" ]
