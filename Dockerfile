FROM golang:1.17 AS builder

WORKDIR /opt

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Build an application
COPY . .
RUN go build -o application .

FROM ubuntu AS production

WORKDIR /opt

COPY --from=builder /opt/application ./

CMD ["./application"]