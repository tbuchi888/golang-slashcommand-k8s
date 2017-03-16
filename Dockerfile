FROM golang:1.8

RUN go get -u github.com/favclip/ucon

COPY ./main.go /go/src/main/
RUN go install main

ENTRYPOINT /go/bin/main
