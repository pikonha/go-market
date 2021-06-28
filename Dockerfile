FROM golang:1.16-stretch

WORKDIR /go/src

CMD ["tail", "-f","/dev/null"]
