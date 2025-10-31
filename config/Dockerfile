FROM golang:1.23.3-alpine
WORKDIR /app
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY . .
RUN go mod download
CMD ["go", "run", "."]
