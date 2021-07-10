FROM golang:1.16-stretch

WORKDIR /go/src

CMD [ "go", "run", "main.go" ]
