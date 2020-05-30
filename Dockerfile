FROM golang:1.12-alpine AS builder

# ENV GO111MODULE=ON \
#     CGO_ENABLED=1 \
#     GOOS=LINUX

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o github.com/Pierrelx/verre-tech/store/cmd .
# RUN chmod +x /store/cmd 

EXPOSE 8081

CMD ["./store/cmd"]

#https://www.callicoder.com/docker-compose-multi-container-orchestration-golang/

#https://medium.com/tourradar/lean-golang-docker-images-using-multi-stage-builds-1015a6b4d1d1

#https://opensource.com/article/19/5/source-image-golang-part-2

#https://dev.to/ivan/go-build-a-minimal-docker-image-in-just-three-steps-514i

#https://stackoverflow.com/questions/50865880/building-golang-project-in-docker-cannot-find-package-in-any-of-gopath-or-go/50884748