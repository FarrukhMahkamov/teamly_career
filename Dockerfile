FROM golang:1.2-buster

RUN go version

ENV GOPATH=/

RUN go mod download
RUN go build teamly_career ./app/cmd/main.go

CMD ["./teamly_career"]