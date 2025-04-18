FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY api/front/etc /app/etc
RUN go build -ldflags="-s -w" -tags no_k8s  -o /app/front api/front/front.go


FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
ENV CONFIG_FILE etc/front-api.yaml

WORKDIR /app
COPY --from=builder /app/front /app/front
COPY --from=builder /app/etc /app/etc

CMD ["sh", "-c", "./front -f ${CONFIG_FILE}"]
