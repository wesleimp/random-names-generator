FROM golang:alpine AS base
WORKDIR /app

FROM node:10 as web_base
WORKDIR /app

# Prepare web_builder
FROM web_base as web_builder
ADD . .
RUN cd ./web; npm install && npm run build

# Go builder
FROM base as go_builder
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o app ./cmd/web/main.go

FROM scratch as final
WORKDIR /app
COPY --from=go_builder /app/app /app
COPY --from=web_builder /app/web/dist ./web/dist

EXPOSE 8080
ENTRYPOINT ["./app"]