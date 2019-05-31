FROM golang:1.12-stretch

LABEL maintainer="tozastation"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/pilu/fresh
RUN go get -u github.com/go-delve/delve/cmd/dlv
COPY . .

EXPOSE 3000

CMD fresh -c runner.conf main.go