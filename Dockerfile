FROM golang:1.13-alpine as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build

FROM scratch
COPY --from=builder /app/uth0 /

EXPOSE 8080

CMD ["./uth0"]
