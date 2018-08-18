FROM golang:1.10-alpine3.7

ENV APPDIR /go/src/github.com/Khigashiguchi/portfolio
WORKDIR $APPDIR
COPY ../../app $APPDIR

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh \
    bash

RUN go get github.com/tockins/realize
RUN go get github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go build -v -o "app"

EXPOSE 8080
CMD ["./app"]
