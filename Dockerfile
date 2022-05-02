FROM golang:1.17 AS API
WORKDIR /sv-auth-gin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server ./main.go && mkdir run && mv ./server ./config.yaml run

FROM alpine:latest
ENV GIN_MODE=release
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >>/etc/apk/repositories && \
    apk add ca-certificates && update-ca-certificates

COPY --from=API /sv-auth-gin/run ./
CMD ./server
