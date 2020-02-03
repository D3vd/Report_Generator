FROM ubuntu:latest

RUN apt-get update

# Install Java and wget
RUN apt-get install -y wget
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk

#  Download and run elasticsearch
RUN wget https://download.elastic.co/elasticsearch/release/org/elasticsearch/distribution/tar/elasticsearch/2.4.0/elasticsearch-2.4.0.tar.gz
RUN tar -xvf elasticsearch-2.4.0.tar.gz
RUN cd elasticsearch-2.4.0

CMD [ "/elasticsearch-2.4.0/bin/elasticsearch"]
