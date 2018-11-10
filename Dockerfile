FROM golang:1.10

WORKDIR /go/src/github.com/135yshr/ctfsendai2018
COPY . .
WORKDIR /go/src/github.com/135yshr/ctfsendai2018/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp .

EXPOSE 8080
ENTRYPOINT ["./goapp"]

