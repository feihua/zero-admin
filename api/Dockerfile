FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/admin api/admin.go


FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/admin /app/admin
COPY --from=builder /app/etc /app/etc

CMD ["./admin", "-f", "etc/admin-api.yaml"]
