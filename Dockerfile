FROM golang:1.26.4 AS builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /src

ARG SERVICE
ENV SERVICE=${SERVICE}

COPY go.work ./
COPY shared/go.mod shared/go.sum* ./shared/
COPY inventory/go.mod inventory/go.sum* ./inventory/
COPY payment/go.mod payment/go.sum* ./payment/
COPY order/go.mod order/go.sum* ./order/

RUN go work sync || true
RUN cd ${SERVICE} && go mod download

COPY shared ./shared
COPY inventory ./inventory
COPY payment ./payment
COPY order ./order

RUN cd ${SERVICE} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd

FROM alpine:3.21.3

RUN apk add --no-cache ca-certificates netcat-openbsd

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /out/server .

RUN chown -R appuser:appgroup /app

USER appuser

CMD ["./server"]