FROM golang:1.14 as builder

WORKDIR /build
COPY main.go main.go

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .


FROM alpine:latest
COPY --from=builder /build/api .
ENTRYPOINT ["./api"]