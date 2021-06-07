FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./assignment cmd/server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/assignment /app/
EXPOSE 3000
CMD ["/app/assignment"]