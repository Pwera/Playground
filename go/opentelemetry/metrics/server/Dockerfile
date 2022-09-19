FROM golang:1.18-alpine
COPY . /usr/src/server/
WORKDIR /usr/src/server/
RUN go install ./main.go
CMD ["/go/bin/main"]