FROM golang:1.25.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./out/raks-coin-exchange ./main.go

FROM alpine:3.22.1

WORKDIR /app

COPY --from=build /app/out/raks-coin-exchange /app/raks-coin-exchange

CMD ["/app/raks-coin-exchange"]