FROM golang:1.9

ENV DIRPATH /go/src/github.com/Khigashiguchi/profile-site
WORKDIR $DIRPATH
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

WORKDIR $DIRPATH
EXPOSE 8080
CMD ["./main"]