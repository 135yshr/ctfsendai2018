FROM golang:1.10

WORKDIR /go/src/github.com/135yshr/ctfsendai2018
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp ./cmd

EXPOSE 8080
ENTRYPOINT ["./goapp"]

