FROM golang:1.22-alpine as builder
WORKDIR /app
ADD . /app
ENV CGO_ENABLED=0
RUN go build -o compute

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/compute .
ENTRYPOINT ["./compute"]