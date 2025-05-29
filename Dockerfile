FROM golang:1.24.3

WORKDIR /app
COPY main.go .
RUN go build -o minq-backend main.go
CMD ["./minq-backend"]

