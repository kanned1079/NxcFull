FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o user-services ./command/server

EXPOSE 50001

CMD ["./user-service"]