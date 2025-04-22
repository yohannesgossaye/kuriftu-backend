FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.io,direct
ENV GOSUMDB=off
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

EXPOSE 8080

CMD ["/app/main"]