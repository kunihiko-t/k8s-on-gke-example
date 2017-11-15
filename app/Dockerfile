FROM golang:1.8-alpine

RUN apk update && \
  apk --no-cache add git mercurial curl make gcc g++ bash && \
  go get bitbucket.org/liamstask/goose/cmd/goose && go get -u github.com/golang/dep/cmd/dep


RUN mkdir -p /go/src/github.com/kunihiko-t/k8s-example/app
WORKDIR /go/src/github.com/kunihiko-t/k8s-example/app
ADD . /go/src/github.com/kunihiko-t/k8s-example/app
RUN dep ensure

EXPOSE 9090