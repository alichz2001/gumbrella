FROM golang:1.26 AS builder

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn task build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
