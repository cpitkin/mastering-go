FROM golang:alpine

RUN mkdir /files

COPY stdout.go /files

WORKDIR /files

RUN go build -o /files/stdout stdout.go

ENTRYPOINT [ "/files/stdout" ]
