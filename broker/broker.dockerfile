FROM alpine:latest

COPY broker_bin .

CMD ["./broker_bin"]