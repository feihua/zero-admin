FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY api/web/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/web api/web/web.go


FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
ENV CONFIG_FILE etc/web-api.yaml

WORKDIR /app
COPY --from=builder /app/web /app/web
COPY --from=builder /app/etc /app/etc

CMD ["sh", "-c", "./web -f ${CONFIG_FILE}"]
