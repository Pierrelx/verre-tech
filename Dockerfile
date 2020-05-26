FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./store .

RUN go build -o store/cmd .
RUN chmod +x ./store/cmd 
EXPOSE 8081

CMD ["./store/cmd"]
#https://www.callicoder.com/docker-compose-multi-container-orchestration-golang/

#https://medium.com/tourradar/lean-golang-docker-images-using-multi-stage-builds-1015a6b4d1d1

#https://opensource.com/article/19/5/source-image-golang-part-2

#https://dev.to/ivan/go-build-a-minimal-docker-image-in-just-three-steps-514i

#https://stackoverflow.com/questions/50865880/building-golang-project-in-docker-cannot-find-package-in-any-of-gopath-or-go/50884748