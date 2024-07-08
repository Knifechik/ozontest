FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

RUN go mod download

COPY . .

RUN go build -o compost ./cmd/compost

FROM alpine

WORKDIR /app

COPY --from=builder /build/compost compost
COPY ./cmd/compost/migrate migrate/

CMD ["/app/compost"]