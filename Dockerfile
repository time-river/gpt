FROM golang:1.17 as builder
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/web
COPY . /go/src/web
RUN go build -o webPing

FROM ubuntu as runner
WORKDIR /go/src/web
COPY --from=builder /go/src/web /go/src/web
EXPOSE 8080
ENTRYPOINT ["./webPing"]