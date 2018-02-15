# This Dockerfile builds on golang:alpine by building Terraform from source
# using the current working directory.

FROM golang:alpine

RUN apk add --update git bash openssh

#WORKDIR $GOPATH/src/github.com/fatdamyr/village/services
COPY village-services /village-services
ENTRYPOINT ["/village-services"]