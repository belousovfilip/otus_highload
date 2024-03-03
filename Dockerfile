FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /social ./cmd/social

ENV CONFIG_PATH="config/config.yaml"

CMD ["/social"]