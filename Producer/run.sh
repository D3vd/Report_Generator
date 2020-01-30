#!/bin/sh

# Start Beanstalk
/usr/bin/beanstalkd &

# Start app
go run .
