FROM golang:1.18-alpine
COPY . /usr/src/client/
WORKDIR /usr/src/client/
RUN go install ./main.go
CMD ["/go/bin/main"]