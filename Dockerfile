FROM golang:1.21

WORKDIR /app

RUN apt update && \
apt install -y libpcap-dev libsodium-dev libzmq3-dev libczmq-dev 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

CMD ["./adam"]