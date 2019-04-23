# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/inamura-nakamura-lab/timecard-api
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
ENV DB_VENDOR="mysql"
ENV USER="timecard"
ENV	PASS="timecard"
#ENV PROTOCOL="tcp(127.0.0.1:3306)"
ENV PROTOCOL="tcp(db)"
ENV	DBNAME="timecard"
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/inamura-nakamura-lab/timecard-api/app /app
EXPOSE 3001
ENTRYPOINT ["/app"]