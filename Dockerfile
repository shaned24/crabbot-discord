FROM golang as builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o demo  examples/demo/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /build /opt/bin