#build stage
FROM golang:1.19 AS builder
WORKDIR /build
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#final stage
FROM alpine:latest
COPY --from=builder /build/local.env .
COPY --from=builder /build/main .
EXPOSE 6969
CMD [ "./main" ]