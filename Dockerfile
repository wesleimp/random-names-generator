FROM golang:alpine AS base
WORKDIR /app
EXPOSE 8080


FROM base as builder
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o app ./cmd/web/main.go

FROM scratch as final
WORKDIR /app
COPY --from=builder /app/app /app
ENTRYPOINT ["./app"]