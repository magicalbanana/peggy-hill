# builder
FROM golang:1.15.0-alpine3.12 as builder

ARG cgo_enabled=0
ARG goos=linux
ARG goarch=amd64

RUN mkdir -p /go/src/github.com/magicalbanana/peggy-hill/

WORKDIR /go/src/github.com/magicalbanana/peggy-hill/

COPY . .

RUN apk add --update --no-cache alpine-sdk git

RUN go mod vendor
RUN CGO_ENABLED=$cgo_enabled GOOS=$goos GOARCH=$goarch \
    go build -mod vendor -v -a -installsuffix cgo -o peggo \
    main.go

# actual container
FROM alpine:3.12

RUN apk add --update --no-cache bash ca-certificates

RUN mkdir -p /app/testdata
RUN mkdir -p /app/bin

WORKDIR /app

COPY --from=builder /go/src/github.com/magicalbanana/peggy-hill/peggo /app/bin
ENV PATH="${PATH}:/app/bin"
