FROM golang:alpine as builder

LABEL maintainer="Pierre Leroux Gatien Montreuil Luigi Croni"

ADD . /go/src/github.com/Pierrelx/verre-tech

RUN go install github.com/Pierrelx/verre-tech/store/cmd

ENTRYPOINT /go/bin/verre-tech/store/cmd

EXPOSE 8080

CMD ["./main"]