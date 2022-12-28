# Base go image for build
FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

# Run the app in a tiny Linux image
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/brokerApp /app
CMD [ "/app/brokerApp" ]