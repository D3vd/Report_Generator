FROM elasticsearch:2.4.0

WORKDIR /
RUN echo "deb [check-valid-until=no] http://cdn-fastly.deb.debian.org/debian jessie main" > /etc/apt/sources.list.d/jessie.list
RUN echo "deb [check-valid-until=no] http://archive.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/jessie-backports.list
RUN sed -i '/deb http:\/\/deb.debian.org\/debian jessie-updates main/d' /etc/apt/sources.list
RUN apt-get -o Acquire::Check-Valid-Until=false update
# RUN apt-get update
RUN apt-get install -y python3
RUN apt-get install -y python3-pip

RUN pip3 install elasticsearch

WORKDIR /tmp/
COPY . .

RUN chmod +x ./import.sh

CMD [ "./import.sh"]
