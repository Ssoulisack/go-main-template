#build stage
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY *.mod .
COPY *.go .
COPY *.yaml .
COPY *.json .
RUN go mod download
COPY . .
COPY config.yaml .
RUN go get -d -v ./ .
RUN go build -o /go/bin/app -v ./ .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata nano
RUN mkdir logs
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config.yaml .

ENV TZ=Asia/Vientiane

ENTRYPOINT ["/app"]
LABEL Name=ithq/kklv2
EXPOSE 3000
