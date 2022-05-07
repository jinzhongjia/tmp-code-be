FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache upx && apk add --no-cache git

WORKDIR /build

COPY . .
RUN go build -ldflags="-s -w" -o /app/tmp-code .


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/tmp-code /app/tmp-code

EXPOSE 8233

CMD ["./tmp-code"]