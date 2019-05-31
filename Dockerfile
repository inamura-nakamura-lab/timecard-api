# Build Container
FROM golang:latest as builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go mod download
RUN go build -o app main.go
# Runtime Container
FROM alpine
ENV DB_VENDOR="mysql"
ENV USER="timecard"
ENV	PASS="timecard"
ENV PROTOCOL="tcp(db)"
ENV	DBNAME="timecard"
ENV PORT=3000
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/app /app
EXPOSE 3001
ENTRYPOINT ["/app"]