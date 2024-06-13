FROM golang:1.21 AS builder

WORKDIR /app

RUN apt update && \
    apt install -y libpcap-dev libsodium-dev libzmq3-dev libczmq-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o adam .

FROM golang:1.21

WORKDIR /app

RUN apt update && \
    apt install -y libpcap0.8 libsodium23 libzmq5 libczmq4 && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/adam /app/

CMD ["./adam"]