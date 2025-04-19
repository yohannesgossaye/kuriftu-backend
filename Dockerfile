FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod download

COPY . .

RUN find . -type f -name "*.go" -exec sed -i 's/\r$//' {} + && \
    find . -type f -name "*.yml" -exec sed -i 's/\r$//' {} + && \
    find . -type f -name "*.yaml" -exec sed -i 's/\r$//' {} + && \
    find . -type f -name "*.sql" -exec sed -i 's/\r$//' {} +

RUN test -f cmd/api/main.go || (echo "Error: cmd/api/main.go not found" && exit 1)
RUN test -f docs/swagger.yaml || (echo "Error: docs/swagger.yaml not found" && exit 1)

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/main ./cmd/api && \
    ls -l /app/main && \
    test -f /app/main || (echo "Error: Binary /app/main not found" && exit 1)

EXPOSE 8080

CMD ["/app/main"]