FROM golang:1.13.5

RUN apt-get update

# CD into app directory
WORKDIR /go/src/app

# Copy contents of the app to the directory
COPY . .

# Install all the dependcies needed for the app
RUN go get -d -v ./...
RUN go install -v ./...

# Expose ports for go server
EXPOSE 8080

CMD ["go", "run", "."]
