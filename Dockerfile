FROM golang:1.24.3

RUN apt-get update && \
    apt-get install -y curl vim

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ["sh", "-c", "go run src/main.go"]