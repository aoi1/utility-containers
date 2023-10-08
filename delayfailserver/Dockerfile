FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go mod download \
&& go build -o main .

FROM scratch
COPY --from=builder /app/main /main
CMD ["/main"]
