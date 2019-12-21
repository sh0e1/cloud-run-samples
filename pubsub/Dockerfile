FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o pubsub

FROM alpine

COPY --from=builder /app/pubsub .

CMD ["/pubsub"]

