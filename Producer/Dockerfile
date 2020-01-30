FROM golang:1.13.5

# Install Beanstalk
RUN apt-get update && apt-get install -y curl build-essential

WORKDIR /
RUN curl -sL https://github.com/kr/beanstalkd/archive/v1.10.tar.gz | tar xvz -C /tmp

WORKDIR /tmp/beanstalkd-1.10
RUN make
RUN cp beanstalkd /usr/bin

# CD into app directory
WORKDIR /go/src/app

# Copy contents of the app to the directory
COPY . .

# Install all the dependcies needed for the app
RUN go get -d -v ./...
RUN go install -v ./...

# Expose ports for beanstalk and go server
EXPOSE 8080
EXPOSE 11300

# Cleanup package manager
RUN apt-get remove --purge -y curl build-essential && apt-get autoclean && apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN chmod +x /go/src/app/run.sh

CMD /go/src/app/run.sh
