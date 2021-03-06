# Build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY  --from=builder /app/public ./public
COPY  --from=builder /app/static ./static
RUN chmod +x ./main
EXPOSE 8080
CMD ./main